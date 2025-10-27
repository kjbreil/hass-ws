package websocket

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/goccy/go-json"
	"github.com/kjbreil/hass-ws/internal/callbacks"
	"github.com/kjbreil/hass-ws/model"
	ws "nhooyr.io/websocket"
)

// Sender handles sending WebSocket messages
type Sender struct {
	ctx       context.Context
	conn      *ws.Conn
	callbacks *callbacks.Callbacks
	nextID    func() int
	logger    *slog.Logger
}

// NewSender creates a new message sender
func NewSender(
	ctx context.Context,
	conn *ws.Conn,
	cb *callbacks.Callbacks,
	nextID func() int,
	logger *slog.Logger,
) *Sender {
	return &Sender{
		ctx:       ctx,
		conn:      conn,
		callbacks: cb,
		nextID:    nextID,
		logger:    logger,
	}
}

// SendWithCallback sends a message and registers a callback for the response
func (s *Sender) SendWithCallback(msg *model.Message, callback chan *model.Message) (int, error) {
	if msg.ID == nil {
		id := s.nextID()
		msg.ID = &id
	}
	s.callbacks.Set(*msg.ID, callback)

	d, _ := json.Marshal(msg)
	err := s.conn.Write(s.ctx, ws.MessageText, d)
	if err != nil {
		return *msg.ID, err
	}
	if !msg.Type.Valid() {
		s.logger.Error(fmt.Sprintf("unknown message type: %s\n", msg.Type), "unknown message")
	}
	return *msg.ID, nil
}

// SendStringWithCallback sends a raw JSON string and registers a callback
func (s *Sender) SendStringWithCallback(msg string, id int, callback chan *model.Message) error {
	s.callbacks.Set(id, callback)

	err := s.conn.Write(s.ctx, ws.MessageText, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

// Send sends a message without expecting a callback
func (s *Sender) Send(msg *model.Message) error {
	if msg.Type != model.MessageTypeAuth {
		id := s.nextID()
		msg.ID = &id
	}

	d, _ := json.Marshal(msg)
	err := s.conn.Write(s.ctx, ws.MessageText, d)
	if err != nil {
		return err
	}
	if !msg.Type.Valid() {
		s.logger.Error(fmt.Sprintf("unknown message type: %s\n", msg.Type), "unknown message")
	}
	return nil
}
