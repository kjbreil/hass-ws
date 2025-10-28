//go:generate sh -c "cd ../../ && go run ./internal/generator/"

package hass

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/kjbreil/hass-ws/internal/callbacks"
	"github.com/kjbreil/hass-ws/internal/handlers"
	"github.com/kjbreil/hass-ws/internal/websocket"
	"github.com/kjbreil/hass-ws/logger"
	"github.com/kjbreil/hass-ws/model"
	"github.com/kjbreil/hass-ws/pkg/rest"
	"github.com/kjbreil/hass-ws/services"
	ws "nhooyr.io/websocket"
)

// Response is a type alias for handlers.Response to expose it publicly
type Response = handlers.Response

// Client is the main Home Assistant WebSocket client
type Client struct {
	conn       *websocket.Connection
	sender     *websocket.Sender
	receiver   *websocket.Receiver
	health     *websocket.HealthMonitor
	router     *handlers.Router
	callbacks  *callbacks.Callbacks
	logger     *slog.Logger
	id         int
	config     *Config
	InitStates bool

	// Public event handlers
	OnMessage   func(message model.Message)
	OnUnhandled func(message model.Message)
	OnType      model.OnTypeHandlers
	OnEntity    model.OnEntityHandlers
	OnGetState  func(states []model.Result)

	subscriptions map[model.EventType]struct{}
}

// NewClient creates a new Home Assistant WebSocket client
func NewClient(c *Config) (*Client, error) {
	return NewClientWithLogger(c, defaultLogger())
}

// NewClientWithLogger creates a new client with a custom logger
func NewClientWithLogger(c *Config, lgr *slog.Logger) (*Client, error) {
	// TODO: Validate Config
	client := &Client{
		config:        c,
		logger:        slog.New(logger.NewWrapHandler(lgr.Handler())),
		callbacks:     callbacks.New(),
		subscriptions: make(map[model.EventType]struct{}),
		OnEntity:      make(model.OnEntityHandlers),
	}

	// Initialize router
	client.router = &handlers.Router{
		OnType:   &client.OnType,
		OnEntity: &client.OnEntity,
	}

	return client, nil
}

// defaultLogger creates a default JSON logger
func defaultLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

// SetLogger sets a custom logger
func (c *Client) SetLogger(logger *slog.Logger) {
	c.logger = logger
}

// Logger returns the current logger
func (c *Client) Logger() *slog.Logger {
	return c.logger
}

// AddSubscription adds an event type subscription
func (c *Client) AddSubscription(eventType model.EventType) {
	c.subscriptions[eventType] = struct{}{}
}

// Connect establishes a WebSocket connection to Home Assistant
func (c *Client) Connect() error {
	// Create websocket config
	wsConfig := &websocket.Config{
		Host:  c.config.Host,
		SSL:   c.config.SSL,
		Port:  c.config.Port,
		Token: c.config.Token,
	}

	// Create connection
	c.conn = websocket.NewConnection(websocket.ConnectionOptions{
		Config:      wsConfig,
		Logger:      c.logger,
		NextID:      c.NextID,
		ResetID:     c.resetID,
		Callbacks:   c.callbacks,
		InitStates:  c.InitStates,
		GetStates:   c.GetStates,
		RunReceiver: c.startReceiver,
	})

	// Establish connection and authorize
	err := c.conn.Connect()
	if err != nil {
		return err
	}

	// Subscribe to events
	err = c.conn.Subscribe(c.subscriptions)
	if err != nil {
		return err
	}

	// Update router handlers
	c.router.OnMessage = c.OnMessage
	c.router.OnUnhandled = c.OnUnhandled

	// Create sender
	c.sender = websocket.NewSender(
		c.conn.Context(),
		c.conn.Conn(),
		c.callbacks,
		c.NextID,
		c.logger,
	)

	// Start receiver and health monitor
	c.startReceiver()

	// Initialize states if requested
	if c.InitStates {
		c.GetStates()
	}

	return nil
}

// startReceiver starts the message receiver and health monitor
func (c *Client) startReceiver() {
	// Create receiver
	c.receiver = websocket.NewReceiver(
		c.conn.Context(),
		c.read,
		c.callbacks,
		c.router,
		c.logger,
	)
	c.receiver.Start()

	// Create health monitor
	c.health = websocket.NewHealthMonitor(
		c.conn.Context(),
		c.sender.SendWithCallback,
		c.callbacks,
		c.conn.Reconnect,
		c.Close,
		c.logger,
	)
	c.health.Start()
}

// NextID returns the next message ID
func (c *Client) NextID() int {
	c.id++
	return c.id
}

// resetID resets the message ID counter
func (c *Client) resetID() {
	c.id = 0
}

// GetStates requests all current states and runs them through handlers
func (c *Client) GetStates() {
	callback := make(chan *model.Message)
	id := c.NextID()

	_, _ = c.sender.SendWithCallback(&model.Message{
		ID:   &id,
		Type: model.MessageTypeGetStates,
	}, callback)

	go func(id int) {
		msg := c.handleCallback(id, callback)
		if msg != nil {
			if c.OnGetState != nil {
				c.OnGetState(msg.Result)
			}

			c.OnType.RunStates(msg)
			c.OnEntity.RunStates(msg)
		}
	}(id)
}

// GetEntityRegistry requests the entity registry
func (c *Client) GetEntityRegistry() *model.Message {
	callback := make(chan *model.Message)
	id := c.NextID()

	_, _ = c.sender.SendWithCallback(&model.Message{
		ID:   &id,
		Type: model.MessageTypeGetEntityRegistry,
	}, callback)
	return c.handleCallback(id, callback)
}

// GetDeviceRegistry requests the device registry
func (c *Client) GetDeviceRegistry() *model.Message {
	callback := make(chan *model.Message)
	id, _ := c.sender.SendWithCallback(&model.Message{
		Type: model.MessageTypeGetDeviceRegistry,
	}, callback)
	return c.handleCallback(id, callback)
}

// GetServices requests available services
func (c *Client) GetServices() *model.Message {
	callback := make(chan *model.Message)
	id, _ := c.sender.SendWithCallback(&model.Message{
		Type: model.MessageTypeGetServices,
	}, callback)
	return c.handleCallback(id, callback)
}

// CallService calls a Home Assistant service
func (c *Client) CallService(service services.Service) *Response {
	id := c.NextID()
	service.SetID(&id)
	callback := make(chan *model.Message)

	rsp := handlers.NewResponse()

	err := c.sender.SendStringWithCallback(service.JSON(), id, callback)
	if err != nil {
		c.logger.Error(err.Error())
		rsp.SetMessage(&model.Message{
			Type: "",
			Error: &model.Error{
				Code:    "",
				Message: err.Error(),
			},
		})
		return rsp
	}
	go func(id int) {
		msg := c.handleCallback(id, callback)
		defer func() {
			rsp.Done <- struct{}{}
		}()
		if msg != nil {
			if msg.Error != nil {
				c.logger.Error(fmt.Sprintf("service %s error code: %s message: %s", service.Name(), msg.Error.Code, msg.Error.Message), "Error sending")
			}
		}
		rsp.SetMessage(msg)

	}(id)
	return rsp
}

// handleCallback waits for a callback response with a timeout
func (c *Client) handleCallback(id int, callback chan *model.Message) *model.Message {
	ticker := time.NewTicker(time.Second * 10)

	defer c.callbacks.Delete(id)
	select {
	case <-c.conn.Context().Done():
		return nil
	case <-ticker.C:
		// TODO: Make error message
		return nil
	case message := <-callback:
		return message
	}
}

// read reads a message from the WebSocket connection
func (c *Client) read() (*model.Message, error) {
	if c.conn.Conn() == nil {
		return nil, errors.New("client is nil")
	}
	_, message, err := c.conn.Conn().Read(c.conn.Context())
	if err != nil {
		return nil, err
	}

	var msg model.Message

	err = json.Unmarshal(message, &msg)
	if err != nil {
		// try to flip to Service
		var d map[string]interface{}

		innerErr := json.Unmarshal(message, &d)
		if innerErr != nil {
			return nil, err
		}
		if r, ok := d["result"]; ok {
			d["service_result"] = r
			delete(d, "result")
			message, _ = json.Marshal(&d)

			err = json.Unmarshal(message, &msg)
			if err != nil {
				return nil, err
			}

			return &msg, nil
		}

		return nil, fmt.Errorf("%s gave error: %w", string(message), err)
	}
	msg.Raw = message
	return &msg, nil
}

// WriteMessage writes a raw message to the WebSocket
func (c *Client) WriteMessage(messageType ws.MessageType, data []byte) error {
	return c.conn.Conn().Write(c.conn.Context(), messageType, data)
}

// Close closes the WebSocket connection
func (c *Client) Close() error {
	// TODO: Chain errors
	var err error
	if c.conn != nil && c.conn.Conn() != nil {
		err = c.conn.Conn().Close(ws.StatusNormalClosure, "")
	}
	// close the connection context
	if c.conn != nil {
		c.conn.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

// RestClient returns the REST client for direct HTTP API access
func (c *Client) RestClient() *rest.Client {
	if c.conn != nil {
		return c.conn.RestClient()
	}
	return nil
}
