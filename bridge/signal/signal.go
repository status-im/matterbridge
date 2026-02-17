package signal

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/42wim/matterbridge/bridge"
	"github.com/42wim/matterbridge/bridge/config"
	"github.com/gorilla/websocket"
)

const maxAttachmentSize = 20 * 1024 * 1024 // 20MB

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

// signal-cli REST API receive types

type signalMessage struct {
	Envelope signalEnvelope `json:"envelope"`
}

type signalEnvelope struct {
	Source      string             `json:"source"`
	SourceName  string             `json:"sourceName"`
	Timestamp   int64              `json:"timestamp"`
	DataMessage *signalDataMessage `json:"dataMessage"`
	EditMessage *signalEditMessage `json:"editMessage"`
}

type signalDataMessage struct {
	Message      string              `json:"message"`
	Timestamp    int64               `json:"timestamp"`
	GroupInfo    *signalGroupInfo    `json:"groupInfo"`
	Attachments  []signalAttachment  `json:"attachments"`
	Quote        *signalQuote        `json:"quote"`
	RemoteDelete *signalRemoteDelete `json:"remoteDelete"`
}

type signalEditMessage struct {
	TargetSentTimestamp int64             `json:"targetSentTimestamp"`
	DataMessage         signalDataMessage `json:"dataMessage"`
}

type signalGroupInfo struct {
	GroupID string `json:"groupId"`
	Type    string `json:"type"`
}

type signalAttachment struct {
	ContentType string `json:"contentType"`
	Filename    string `json:"filename"`
	ID          string `json:"id"`
	Size        int64  `json:"size"`
}

type signalQuote struct {
	ID     int64  `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

type signalRemoteDelete struct {
	Timestamp int64 `json:"timestamp"`
}

// signal-cli REST API send types

type signalSendRequest struct {
	Message           string   `json:"message"`
	Number            string   `json:"number"`
	Recipients        []string `json:"recipients"`
	Base64Attachments []string `json:"base64_attachments,omitempty"`
	QuoteTimestamp    *int64   `json:"quote_timestamp,omitempty"`
	QuoteAuthor       string   `json:"quote_author,omitempty"`
	QuoteMessage      string   `json:"quote_message,omitempty"`
	EditTimestamp     *int64   `json:"edit_timestamp,omitempty"`
}

type signalSendResponse struct {
	Timestamp string `json:"timestamp"`
}

type signalRemoteDeleteRequest struct {
	Recipient string `json:"recipient"`
	Timestamp int64  `json:"timestamp"`
}

type signalGroup struct {
	ID         string `json:"id"`
	InternalID string `json:"internal_id"`
}

func New(cfg *bridge.Config) bridge.Bridger {
	return &Bsignal{
		Config:    cfg,
		fetchDone: make(chan bool),
		client:    &http.Client{Timeout: 60 * time.Second},
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
	// Handle delete from other bridges
	if msg.Event == config.EventMsgDelete {
		return b.handleSendDelete(msg)
	}

	// Prepend username to message text since Signal has no webhook/impersonation
	text := msg.Text
	if msg.Username != "" {
		text = msg.Username + msg.Text
	}

	b.Log.Debugf("=> Sending message to Signal group %s: %s", msg.Channel, text)

	req := signalSendRequest{
		Message:    text,
		Number:     b.number,
		Recipients: []string{"group." + msg.Channel},
	}

	// Handle edit from other bridges — msg.ID is set when gateway finds existing message in cache
	if msg.ID != "" {
		if ts, err := strconv.ParseInt(msg.ID, 10, 64); err == nil {
			req.EditTimestamp = &ts
			b.Log.Debugf("Editing Signal message with timestamp %d", ts)
		}
	}

	if msg.ParentID != "" {
		if ts, err := strconv.ParseInt(msg.ParentID, 10, 64); err == nil {
			req.QuoteTimestamp = &ts
		}
	}

	if files, ok := msg.Extra["file"]; ok {
		for _, f := range files {
			fi, ok := f.(config.FileInfo)
			if !ok {
				continue
			}
			var data []byte
			if fi.Data != nil {
				data = *fi.Data
			} else if fi.URL != "" {
				resp, err := b.client.Get(fi.URL)
				if err != nil {
					b.Log.Errorf("Failed to download attachment from %s: %s", fi.URL, err)
					continue
				}
				data, _ = io.ReadAll(io.LimitReader(resp.Body, maxAttachmentSize))
				resp.Body.Close()
			}
			if len(data) > 0 {
				mt := mimeFromFilename(fi.Name)
				b64 := fmt.Sprintf("data:%s;base64,%s", mt, base64.StdEncoding.EncodeToString(data))
				req.Base64Attachments = append(req.Base64Attachments, b64)
			}
		}
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

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("signal-cli API returned %d: %s", resp.StatusCode, string(respBody))
	}

	var sendResp signalSendResponse
	if err := json.Unmarshal(respBody, &sendResp); err == nil && sendResp.Timestamp != "" {
		return sendResp.Timestamp, nil
	}

	return "", nil
}

func (b *Bsignal) handleSendDelete(msg config.Message) (string, error) {
	if msg.ID == "" {
		b.Log.Debug("Ignoring delete event with empty message ID")
		return "", nil
	}

	ts, err := strconv.ParseInt(msg.ID, 10, 64)
	if err != nil {
		b.Log.Debugf("Ignoring delete event with non-numeric ID: %s", msg.ID)
		return "", nil
	}

	b.Log.Debugf("=> Deleting Signal message with timestamp %d in group %s", ts, msg.Channel)

	delReq := signalRemoteDeleteRequest{
		Recipient: "group." + msg.Channel,
		Timestamp: ts,
	}

	data, err := json.Marshal(delReq)
	if err != nil {
		return "", fmt.Errorf("failed to marshal delete request: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/v1/remote-delete/%s", b.apiURL, b.number),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create delete request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := b.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send delete request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("signal-cli delete returned %d: %s", resp.StatusCode, string(respBody))
	}

	return "", nil
}

// WebSocket receive loop with automatic reconnection

func (b *Bsignal) receiveLoop() {
	for {
		select {
		case <-b.fetchDone:
			return
		default:
			b.receiveMessages()
			time.Sleep(3 * time.Second)
		}
	}
}

func (b *Bsignal) receiveMessages() {
	wsURL := fmt.Sprintf("%s/v1/receive/%s", b.wsURL, b.number)
	b.Log.Debugf("Opening WebSocket connection to %s", wsURL)

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
			return
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
	if msg.Envelope.Source == b.number {
		return
	}

	if msg.Envelope.EditMessage != nil {
		b.handleEditMessage(msg)
		return
	}

	dm := msg.Envelope.DataMessage
	if dm == nil {
		return
	}

	if dm.GroupInfo == nil {
		return
	}

	channel, ok := b.groupMap[dm.GroupInfo.GroupID]
	if !ok {
		b.Log.Warnf("Unknown group internal_id: %s", dm.GroupInfo.GroupID)
		return
	}

	if dm.RemoteDelete != nil {
		b.Remote <- config.Message{
			Event:    config.EventMsgDelete,
			ID:       strconv.FormatInt(dm.RemoteDelete.Timestamp, 10),
			Channel:  channel,
			Account:  b.Account,
			Protocol: "signal",
		}
		return
	}

	if dm.Message == "" && len(dm.Attachments) == 0 {
		return
	}

	username := msg.Envelope.SourceName
	if username == "" {
		username = msg.Envelope.Source
	}

	rmsg := config.Message{
		Username:  username,
		Text:      dm.Message,
		Channel:   channel,
		Account:   b.Account,
		Protocol:  "signal",
		ID:        strconv.FormatInt(dm.Timestamp, 10),
		Timestamp: time.UnixMilli(dm.Timestamp),
	}

	if dm.Quote != nil {
		rmsg.ParentID = strconv.FormatInt(dm.Quote.ID, 10)
	}

	for _, att := range dm.Attachments {
		data, err := b.downloadAttachment(att.ID)
		if err != nil {
			b.Log.Errorf("Failed to download attachment %s: %s", att.ID, err)
			continue
		}
		fi := config.FileInfo{
			Name:    att.Filename,
			Data:    data,
			Size:    att.Size,
			Comment: dm.Message,
		}
		if rmsg.Extra == nil {
			rmsg.Extra = make(map[string][]interface{})
		}
		rmsg.Extra["file"] = append(rmsg.Extra["file"], fi)
	}

	if len(dm.Attachments) > 0 && dm.Message != "" {
		rmsg.Text = ""
	}

	b.Remote <- rmsg
}

func (b *Bsignal) handleEditMessage(msg signalMessage) {
	em := msg.Envelope.EditMessage
	dm := &em.DataMessage

	if dm.GroupInfo == nil {
		return
	}

	channel, ok := b.groupMap[dm.GroupInfo.GroupID]
	if !ok {
		b.Log.Warnf("Unknown group internal_id: %s", dm.GroupInfo.GroupID)
		return
	}

	if dm.Message == "" {
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
		ID:        strconv.FormatInt(em.TargetSentTimestamp, 10),
		Timestamp: time.UnixMilli(dm.Timestamp),
	}
}

func (b *Bsignal) downloadAttachment(id string) (*[]byte, error) {
	resp, err := b.client.Get(fmt.Sprintf("%s/v1/attachments/%s", b.apiURL, id))
	if err != nil {
		return nil, fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("attachment endpoint returned %d", resp.StatusCode)
	}

	data, err := io.ReadAll(io.LimitReader(resp.Body, maxAttachmentSize))
	if err != nil {
		return nil, fmt.Errorf("read failed: %w", err)
	}
	return &data, nil
}

func mimeFromFilename(name string) string {
	t := mime.TypeByExtension(filepath.Ext(name))
	if t == "" {
		return "application/octet-stream"
	}
	return t
}