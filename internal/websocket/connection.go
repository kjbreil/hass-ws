package websocket

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/kjbreil/hass-ws/internal/callbacks"
	"github.com/kjbreil/hass-ws/model"
	"github.com/kjbreil/hass-ws/pkg/rest"
	ws "nhooyr.io/websocket"
)

// Config represents the WebSocket connection configuration
type Config struct {
	Host  string
	SSL   bool
	Port  int
	Token string
}

// Connection manages the WebSocket connection to Home Assistant
type Connection struct {
	ctx         context.Context
	cancel      context.CancelFunc
	conn        *ws.Conn
	config      *Config
	callbacks   *callbacks.Callbacks
	logger      *slog.Logger
	running     bool
	nextID      func() int
	resetID     func()
	restClient  *rest.Client
	initStates  bool
	getStates   func()
	runReceiver func()
}

// ConnectionOptions holds options for creating a connection
type ConnectionOptions struct {
	Config      *Config
	Logger      *slog.Logger
	NextID      func() int
	ResetID     func()
	Callbacks   *callbacks.Callbacks
	InitStates  bool
	GetStates   func()
	RunReceiver func()
}

// NewConnection creates a new WebSocket connection manager
func NewConnection(opts ConnectionOptions) *Connection {
	return &Connection{
		config:      opts.Config,
		logger:      opts.Logger,
		callbacks:   opts.Callbacks,
		nextID:      opts.NextID,
		resetID:     opts.ResetID,
		initStates:  opts.InitStates,
		getStates:   opts.GetStates,
		runReceiver: opts.RunReceiver,
	}
}

// Connect establishes a WebSocket connection and performs authorization
func (c *Connection) Connect() error {
	if c.cancel != nil {
		c.cancel()
	}

	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.resetID()

	c.restClient = rest.New(c.config.Host, c.config.Port, c.config.Token, c.config.SSL)

	var err error
	scheme := "ws"
	if c.config.SSL {
		scheme = "wss"
	}

	u := url.URL{Scheme: scheme, Host: fmt.Sprintf("%s:%d", c.config.Host, c.config.Port), Path: "/api/websocket"}

	c.conn, _, err = ws.Dial(c.ctx, u.String(), nil)

	if err != nil {
		c.cancel()
		return fmt.Errorf("dial: %w", err)
	}

	c.conn.SetReadLimit(32768 * 100)

	// Mark as running
	c.running = true

	// Authorize the websocket connection
	err = c.authorize()
	if err != nil {
		return err
	}

	return nil
}

// authorize performs WebSocket authentication
func (c *Connection) authorize() error {
	// Wait for auth required message
	msg := &model.Message{}
	var err error
	msg, err = c.read()
	if err != nil {
		c.cancel()
		return fmt.Errorf("read: %w", err)
	}
	if msg.Type != model.MessageTypeAuthRequired {
		return fmt.Errorf("initial response not AuthRequired: %s", msg.Raw)
	}

	// Send auth message
	err = c.send(&model.Message{
		Type:        model.MessageTypeAuth,
		AccessToken: &c.config.Token,
	})
	if err != nil {
		return err
	}

	return nil
}

// Subscribe subscribes to event types
func (c *Connection) Subscribe(eventTypes map[model.EventType]struct{}) error {
	for eventType := range eventTypes {
		var err error
		if eventType == model.EventTypeAll {
			err = c.send(&model.Message{
				Type: model.MessageTypeSubscribeEvents,
			})
		} else {
			err = c.send(&model.Message{
				Type:      model.MessageTypeSubscribeEvents,
				EventType: &eventType,
			})
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// send sends a message without expecting a callback
func (c *Connection) send(msg *model.Message) error {
	sender := NewSender(c.ctx, c.conn, c.callbacks, c.nextID, c.logger)
	return sender.Send(msg)
}

// read reads a message from the WebSocket
func (c *Connection) read() (*model.Message, error) {
	_, d, err := c.conn.Read(c.ctx)
	if err != nil {
		return nil, err
	}
	msg, err := model.ParseMessage(d)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

// Reconnect attempts to reconnect the WebSocket connection
func (c *Connection) Reconnect() error {
	if c.cancel != nil {
		c.cancel()
	}
	if !c.running {
		return errors.New("cannot reconnect, client is not running")
	}

	var err error
	retries := 120
	for i := 0; i < retries; i++ {
		c.logger.Info(fmt.Sprintf("attempting to reconnect: %d", i))
		err = c.Connect()
		if err == nil {
			// Re-run the receiver after reconnection
			if c.runReceiver != nil {
				c.runReceiver()
			}
			// Re-fetch states if needed
			if c.initStates && c.getStates != nil {
				c.getStates()
			}
			return nil
		}
		c.logger.Error(fmt.Sprintf("reconnect failed: %v", err), "reconnect failed")
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return fmt.Errorf("websocket reconnect failed: %w", err)
	}
	return nil
}

// Context returns the connection context
func (c *Connection) Context() context.Context {
	return c.ctx
}

// Conn returns the underlying WebSocket connection
func (c *Connection) Conn() *ws.Conn {
	return c.conn
}

// RestClient returns the REST client
func (c *Connection) RestClient() *rest.Client {
	return c.restClient
}

// Close cancels the connection context
func (c *Connection) Close() {
	if c.cancel != nil {
		c.cancel()
	}
}
