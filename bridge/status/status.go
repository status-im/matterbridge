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
	//"go.uber.org/zap"

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

	// message fetching loop controls
	fetchInterval time.Duration
	fetchTimeout  time.Duration
	fetchDone     chan bool

	// Whisper node settings
	whisperListenPort int
	whisperListenAddr string
	whisperDataDir    string

	privateKey *ecdsa.PrivateKey  // secret for Status chat identity
	nodeConfig *params.NodeConfig // configuration for Whisper node
	statusNode *gonode.StatusNode // Ethereum Whisper node to run in background
	messenger  *status.Messenger  // Status messaging layer instance
}

func New(cfg *bridge.Config) bridge.Bridger {
	return &Bstatus{
		Config:            cfg,
		fetchDone:         make(chan bool),
		whisperListenPort: 30303,
		whisperListenAddr: "0.0.0.0",
		// TODO parametrize those
		whisperDataDir: "/tmp/matterbridge-status-data",
		fetchTimeout:   500 * time.Millisecond,
		fetchInterval:  500 * time.Millisecond,
	}
}

func (b *Bstatus) Connect() error {
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
	//logger, _ := zap.NewProduction()
	//logger := zap.NewNop()

	// Using an in-memory SQLite DB since we have nothing worth preserving
	db, _ := sql.Open("sqlite3", "file:mem?mode=memory&cache=shared")
	options := []status.Option{
		status.WithDatabase(db),
		//status.WithCustomLogger(logger),
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
	b.messenger = messenger

	// Start a routine for periodically fetching messages
	go b.fetchMessagesLoop()

	return nil
}

func (b *Bstatus) Disconnect() error {
	b.stopMessagesLoops()
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

	msgHash, err := b.messenger.SendChatMessage(ctx, genStatusMsg(msg))
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
func genStatusMsg(msg config.Message) (sMsg *status.Message) {
	sMsg.ChatId = msg.Channel
	sMsg.Text = msg.Text
	sMsg.ContentType = protobuf.ChatMessage_TEXT_PLAIN
	return
}

// Generate a sane configuration for a Status Node
func (b *Bstatus) generateConfig() *params.NodeConfig {
	options := []params.Option{
		params.WithFleet(params.FleetBeta),
		b.withListenAddr(),
	}

	var configFiles []string
	config, err := params.NewNodeConfigWithDefaultsAndFiles(
		b.whisperDataDir,
		params.MainNetworkID,
		options,
		configFiles,
	)
	if err != nil {
		b.Log.WithError(err).Error("Failed to generate config")
	}
	return config
}

func (b *Bstatus) stopMessagesLoops() {
	close(b.fetchDone)
}

// Main loop for fetching Status messages and relaying them to the bridge
func (b *Bstatus) fetchMessagesLoop() {
	t := time.NewTicker(b.fetchInterval)
	defer t.Stop()
	for {
		select {
		case <-t.C:
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

func (b *Bstatus) propagateMessage(msg *status.Message) {
	pubKey := publicKeyToHex(msg.SigPubKey)
	alias, err := alias.GenerateFromPublicKeyString(pubKey)
	if err != nil {
		b.Log.WithError(err).Error("Failed to generate Chat name")
	}
	// Send message for processing
	b.Remote <- config.Message{
		Timestamp: time.Unix(int64(msg.WhisperTimestamp), 0),
		Username:  alias,
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
		b.whisperListenAddr = addr
	}
	if port := b.GetInt("ListenPort"); port != 0 {
		b.whisperListenPort = port
	}
	return func(c *params.NodeConfig) error {
		c.ListenAddr = fmt.Sprintf("%s:%d", b.whisperListenAddr, b.whisperListenPort)
		return nil
	}
}
