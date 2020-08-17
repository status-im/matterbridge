package status

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/42wim/matterbridge/bridge"
	"github.com/42wim/matterbridge/bridge/config"

	crypto "github.com/ethereum/go-ethereum/crypto"
	gethbridge "github.com/status-im/status-go/eth-node/bridge/geth"
	gonode "github.com/status-im/status-go/node"
	params "github.com/status-im/status-go/params"
	status "github.com/status-im/status-go/protocol"
	alias "github.com/status-im/status-go/protocol/identity/alias"
	"github.com/status-im/status-go/protocol/protobuf"
)

type Bstatus struct {
	*bridge.Config

	// Maximum number of characters a username can have
	maxUsernameLen int

	// message fetching loop controls
	fetchInterval time.Duration
	fetchTimeout  time.Duration
	fetchDone     chan bool

	// Waku node settings
	wakuListenPort int
	wakuListenAddr string
	wakuDataDir    string

	// ENS settings
	ensName      string        // Make sure the bridge account owns the ENS name
	ensVerifyURL string        // URL of Infura endpoint to call
	ensContract  string        // Address of ENS resolving contract
	ensDone      chan bool     // Control channel for stopping ENS checking loop
	ensInterval  time.Duration // Frequency of ENS verification checks

	privateKey *ecdsa.PrivateKey  // secret for Status chat identity
	nodeConfig *params.NodeConfig // configuration for Waku node
	statusNode *gonode.StatusNode // Ethereum Waku node to run in background
	messenger  *status.Messenger  // Status messaging layer instance
}

func New(cfg *bridge.Config) bridge.Bridger {
	return &Bstatus{
		Config:         cfg,
		fetchDone:      make(chan bool),
		wakuListenPort: 30303,
		wakuListenAddr: "0.0.0.0",
		// TODO parametrize those
		maxUsernameLen: 40,
		wakuDataDir:    "/tmp/matterbridge-status-data",
		fetchTimeout:   500 * time.Millisecond,
		fetchInterval:  500 * time.Millisecond,
		// ENS checks are slow, also db has a lock
		ensInterval:  30 * time.Second,
		ensVerifyURL: "https://mainnet.infura.io/v3/f315575765b14720b32382a61a89341a",
		ensContract:  "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e",
	}
}

func (b *Bstatus) Connect() error {
	// Will be displayed to other users if registered
	b.ensName = b.GetString("Nick")

	keyHex := strings.TrimPrefix(b.GetString("PrivateKey"), "0x")
	if privKey, err := crypto.HexToECDSA(keyHex); err != nil {
		return errors.Wrap(err, "Failed to parse PrivateKey")
	} else {
		b.privateKey = privKey
	}

	b.nodeConfig = b.generateConfig()
	b.statusNode = gonode.New()
	accsMgr, _ := b.statusNode.AccountManager()

	if err := b.statusNode.Start(b.nodeConfig, accsMgr); err != nil {
		return errors.Wrap(err, "Failed to start Status node")
	}

	// Create a custom logger to suppress DEBUG messages
	logger, _ := zap.NewProduction()

	// Using an in-memory SQLite DB since we have nothing worth preserving
	db, err := sql.Open("sqlite3", "file:mem?mode=memory&cache=shared")
	if err != nil {
		return errors.Wrap(err, "Failed to open sqlite database")
	}
	options := []status.Option{
		status.WithDatabase(db),
		status.WithCustomLogger(logger),
	}

	var instID string = uuid.New().String()

	messenger, err := status.NewMessenger(
		b.privateKey,
		gethbridge.NewNodeBridge(b.statusNode.GethNode()),
		instID,
		options...,
	)
	if err != nil {
		return errors.Wrap(err, "Failed to create Messenger")
	}
	if err := messenger.Start(); err != nil {
		return errors.Wrap(err, "Failed to start Messenger")
	}
	if err := messenger.Init(); err != nil {
		return errors.Wrap(err, "Failed to init Messenger")
	}
	b.messenger = messenger

	// Start a routine for periodically fetching messages
	go b.fetchMessagesLoop()
	// Start a routine for periodically checking ENS names
	go b.checkEnsNamesLoop()

	return nil
}

func (b *Bstatus) Disconnect() error {
	b.stopMessagesLoop()
	b.stopEnsNamesLoop()
	if err := b.messenger.Shutdown(); err != nil {
		return errors.Wrap(err, "Failed to stop Status messenger")
	}
	if err := b.statusNode.Stop(); err != nil {
		return errors.Wrap(err, "Failed to stop Status node")
	}
	return nil
}

func (b *Bstatus) JoinChannel(channel config.ChannelInfo) error {
	chat := status.CreatePublicChat(channel.Name, b.messenger.Timesource())
	b.messenger.Join(chat)
	b.messenger.SaveChat(&chat)
	return nil
}

func (b *Bstatus) Send(msg config.Message) (string, error) {
	if !b.Connected() {
		return "", fmt.Errorf("bridge %s not connected, dropping message %#v to bridge", b.Account, msg)
	}

	if skipBridgeMessage(msg) {
		return "", nil
	}
	b.Log.Infof("=> Sending message %#v", msg)

	// Use a timeout for sending messages
	ctx, cancel := context.WithTimeout(context.Background(), b.fetchTimeout)
	defer cancel()

	msgHash, err := b.messenger.SendChatMessage(ctx, b.genStatusMsg(msg))
	if err != nil {
		return "", errors.Wrap(err, "failed to send message")
	}
	// TODO handle the delivery event?
	return fmt.Sprintf("%#x", msgHash), nil
}

func (b *Bstatus) Connected() bool {
	return b.statusNode.IsRunning()
}

// Converts a bridge message into a Status message
func (b *Bstatus) genStatusMsg(msg config.Message) (sMsg *status.Message) {
	sMsg = &status.Message{}
	sMsg.EnsName = b.ensName
	sMsg.ChatId = msg.Channel
	sMsg.ContentType = protobuf.ChatMessage_TEXT_PLAIN
	// We need to prefix messages with usernames
	sMsg.Text = fmt.Sprintf("%s%s", msg.Username, msg.Text)
	return
}

// Generate a sane configuration for a Status Node
func (b *Bstatus) generateConfig() *params.NodeConfig {
	options := []params.Option{
		params.WithFleet("eth.prod"),
		b.withNodeName(),
		b.withListenAddr(),
		b.withWakuEnabled(),
	}

	var configFiles []string
	config, err := params.NewNodeConfigWithDefaultsAndFiles(
		b.wakuDataDir,
		params.MainNetworkID,
		options,
		configFiles,
	)

	if err != nil {
		b.Log.WithError(err).Error("Failed to generate config")
	}
	return config
}

func (b *Bstatus) stopMessagesLoop() {
	close(b.fetchDone)
}

func (b *Bstatus) stopEnsNamesLoop() {
	close(b.ensDone)
}

// Main loop for fetching Status messages and relaying them to the bridge
func (b *Bstatus) fetchMessagesLoop() {
	ticker := time.NewTicker(b.fetchInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			mResp, err := b.messenger.RetrieveAll()
			if err != nil {
				b.Log.WithError(err).Error("Failed to retrieve messages")
				continue
			}
			for _, msg := range mResp.Messages {
				if b.skipStatusMessage(msg) {
					continue
				}
				b.propagateMessage(msg)
			}
		case <-b.fetchDone:
			return
		}
	}
}

func (b *Bstatus) checkEnsNamesLoop() {
	ticker := time.NewTicker(b.ensInterval)
	defer ticker.Stop()
	ctx, cancelVerifyENS := context.WithCancel(context.Background())
	for {
		select {
		case <-ticker.C:
			_, err := b.messenger.VerifyENSNames(ctx, b.ensVerifyURL, b.ensContract)
			if err != nil {
				b.Log.WithError(err).Error("Failed to validate ENS name")
				continue
			}
		case <-b.ensDone:
			cancelVerifyENS()
			return
		}
	}
}

func (b *Bstatus) propagateMessage(msg *status.Message) {
	pubKey := publicKeyToHex(msg.SigPubKey)
	var username string
	// Contact can have an ENS Name, but needs to be verified
	contact, err := b.messenger.GetContactByID(pubKey)
	if err != nil {
		b.Log.WithError(err).Error("Not yet verified contact:", pubKey)
		username, err = alias.GenerateFromPublicKeyString(pubKey)
		if err != nil { // fallback to full public key
			b.Log.WithError(err).Error("Failed to generate Chat name")
			username = pubKey
		}
	} else if contact.ENSVerified { // trim our domain for brevity
		username = strings.TrimSuffix(contact.Name, ".stateofus.eth")
	} else { // fallback to 3-word chat name
		username = contact.Alias
	}
	// Trim username in case of LONG ENS names
	if len(username) > b.maxUsernameLen {
		username = username[:b.maxUsernameLen]
	}
	// Send message for processing
	b.Remote <- config.Message{
		Timestamp: time.Unix(int64(msg.WhisperTimestamp), 0),
		Username:  username,
		UserID:    pubKey,
		Text:      msg.Text,
		Channel:   msg.ChatId,
		ID:        fmt.Sprintf("%#x", msg.ID),
		Account:   b.Account,
	}
}

// skipStatusMessage defines which Status messages can be ignored
func (b *Bstatus) skipStatusMessage(msg *status.Message) bool {
	// skip messages from ourselves
	if isPubKeyEqual(msg.SigPubKey, &b.privateKey.PublicKey) {
		return true
	}

	// skip empty messages
	if msg.Text == "" {
		return true
	}

	return false
}

// skipBridgeMessage defines which messages from the bridge should be ignored
func skipBridgeMessage(msg config.Message) bool {
	// skip delete messages
	if msg.Event == config.EventMsgDelete {
		return true
	}
	return false
}

func (b *Bstatus) withListenAddr() params.Option {
	if addr := b.GetString("ListenAddr"); addr != "" {
		b.wakuListenAddr = addr
	}
	if port := b.GetInt("ListenPort"); port != 0 {
		b.wakuListenPort = port
	}
	return func(c *params.NodeConfig) error {
		c.ListenAddr = fmt.Sprintf("%s:%d", b.wakuListenAddr, b.wakuListenPort)
		return nil
	}
}

func (b *Bstatus) withNodeName() params.Option {
	return func(c *params.NodeConfig) error {
		// To make it different from Statusd and StatusIM nodes
		c.Name = "StatusBridge"
		return nil
	}
}

func (b *Bstatus) withWakuEnabled() params.Option {
	return func(c *params.NodeConfig) error {
		// Disable Whisper
		c.WhisperConfig.Enabled = false
		// Enable Waku
		c.WakuConfig.Enabled = true
		return nil
	}
}
