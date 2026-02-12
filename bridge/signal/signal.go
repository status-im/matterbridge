package signal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/42wim/matterbridge/bridge"
	"github.com/42wim/matterbridge/bridge/config"
	"github.com/gorilla/websocket"
)

type Bsignal struct {
	*bridge.Config

	apiURL    string
	wsURL     string
	number    string
	fetchDone chan bool
	client    *http.Client
	// maps signal-cli internal_id (from receive) -> id (used in gateway config and send)
	groupMap map[string]string
}

type signalMessage struct {
	Envelope signalEnvelope `json:"envelope"`
}

type signalEnvelope struct {
	Source      string             `json:"source"`
	SourceName  string             `json:"sourceName"`
	Timestamp   int64              `json:"timestamp"`
	DataMessage *signalDataMessage `json:"dataMessage"`
}

type signalDataMessage struct {
	Message   string           `json:"message"`
	Timestamp int64            `json:"timestamp"`
	GroupInfo *signalGroupInfo `json:"groupInfo"`
}

type signalGroupInfo struct {
	GroupID string `json:"groupId"`
	Type    string `json:"type"`
}

type signalSendRequest struct {
	Message    string   `json:"message"`
	Number     string   `json:"number"`
	Recipients []string `json:"recipients"`
}

type signalGroup struct {
	ID         string `json:"id"`
	InternalID string `json:"internal_id"`
}

func New(cfg *bridge.Config) bridge.Bridger {
	return &Bsignal{
		Config:    cfg,
		fetchDone: make(chan bool),
		client:    &http.Client{Timeout: 30 * time.Second},
		groupMap:  make(map[string]string),
	}
}

func (b *Bsignal) Connect() error {
	b.apiURL = b.GetString("SignalAPIURL")
	if b.apiURL == "" {
		return fmt.Errorf("SignalAPIURL not configured")
	}

	b.number = b.GetString("Number")
	if b.number == "" {
		return fmt.Errorf("Number not configured")
	}

	// Derive WebSocket URL from HTTP URL
	b.wsURL = strings.Replace(b.apiURL, "http://", "ws://", 1)
	b.wsURL = strings.Replace(b.wsURL, "https://", "wss://", 1)

	b.Log.Infof("Connecting to signal-cli REST API at %s for number %s", b.apiURL, b.number)

	resp, err := b.client.Get(fmt.Sprintf("%s/v1/groups/%s", b.apiURL, b.number))
	if err != nil {
		return fmt.Errorf("failed to connect to signal-cli API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("signal-cli API returned status %d", resp.StatusCode)
	}

	// Build group ID mapping: receive endpoint returns internal_id,
	// but gateway config and send endpoint use id
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read groups response: %w", err)
	}

	var groups []signalGroup
	if err := json.Unmarshal(body, &groups); err != nil {
		return fmt.Errorf("failed to parse groups: %w", err)
	}

	for _, g := range groups {
		id := strings.TrimPrefix(g.ID, "group.")
		b.groupMap[g.InternalID] = id
		b.Log.Debugf("Group mapping: %s -> %s", g.InternalID, id)
	}

	b.Log.Info("Connected to signal-cli REST API")
	go b.receiveLoop()

	return nil
}

func (b *Bsignal) Disconnect() error {
	close(b.fetchDone)
	return nil
}

func (b *Bsignal) JoinChannel(channel config.ChannelInfo) error {
	b.Log.Infof("Joining Signal group %s", channel.Name)
	return nil
}

func (b *Bsignal) Send(msg config.Message) (string, error) {
	if msg.Event == config.EventMsgDelete {
		return "", nil
	}

	b.Log.Debugf("=> Sending message to Signal group %s: %s", msg.Channel, msg.Text)

	req := signalSendRequest{
		Message:    msg.Text,
		Number:     b.number,
		Recipients: []string{"group." + msg.Channel},
	}

	data, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal message: %w", err)
	}

	resp, err := b.client.Post(
		fmt.Sprintf("%s/v2/send", b.apiURL),
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", fmt.Errorf("failed to send message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("signal-cli API returned %d: %s", resp.StatusCode, string(body))
	}

	return "", nil
}

func (b *Bsignal) receiveLoop() {
	for {
		select {
		case <-b.fetchDone:
			return
		default:
			b.receiveMessages()
			// Backoff before reconnecting
			time.Sleep(3 * time.Second)
		}
	}
}

func (b *Bsignal) receiveMessages() {
	wsURL := fmt.Sprintf("%s/v1/receive/%s", b.wsURL, b.number)
	b.Log.Infof("Opening WebSocket connection to %s", wsURL)

	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		b.Log.Errorf("WebSocket connection failed: %s", err)
		return
	}
	defer conn.Close()

	b.Log.Info("WebSocket connected, listening for messages")

	for {
		select {
		case <-b.fetchDone:
			return
		default:
		}

		_, rawMsg, err := conn.ReadMessage()
		if err != nil {
			b.Log.Errorf("WebSocket read error: %s", err)
			return // Will reconnect via receiveLoop
		}

		var msg signalMessage
		if err := json.Unmarshal(rawMsg, &msg); err != nil {
			b.Log.Debugf("Failed to parse WebSocket message: %s", err)
			continue
		}

		b.handleMessage(msg)
	}
}

func (b *Bsignal) handleMessage(msg signalMessage) {
	dm := msg.Envelope.DataMessage
	if dm == nil || dm.Message == "" {
		return
	}

	if dm.GroupInfo == nil {
		return
	}

	if msg.Envelope.Source == b.number {
		return
	}

	// Map internal_id from receive to id used in gateway config
	channel, ok := b.groupMap[dm.GroupInfo.GroupID]
	if !ok {
		b.Log.Warnf("Unknown group internal_id: %s", dm.GroupInfo.GroupID)
		return
	}

	username := msg.Envelope.SourceName
	if username == "" {
		username = msg.Envelope.Source
	}

	b.Remote <- config.Message{
		Username:  username,
		Text:      dm.Message,
		Channel:   channel,
		Account:   b.Account,
		Protocol:  "signal",
		Timestamp: time.UnixMilli(dm.Timestamp),
	}
}