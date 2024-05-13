//go:generate go run ./helpers/

package hass_ws

import (
	"context"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/kjbreil/hass-ws/logger"
	"github.com/kjbreil/hass-ws/model"
	"github.com/kjbreil/hass-ws/services"
	"log/slog"
	"nhooyr.io/websocket"
	"os"
	"time"
)

type Client struct {
	client  *websocket.Conn
	config  Config
	id      int
	running bool

	subscriptions map[model.EventType]int
	OnMessage     func(message model.Message)
	OnUnhandled   func(message model.Message)
	InitStates    bool

	eventWatch map[string]func(event model.Event)

	callbacks *callbacks

	logger *slog.Logger

	ctx    context.Context
	cancel context.CancelFunc

	OnType     model.OnTypeHandlers
	OnEntity   model.OnEntityHandlers
	OnGetState func(states []model.Result)
}

func NewClient(c *Config) (*Client, error) {
	return NewClientWithLogger(c, defaultLogger())
}

func NewClientWithLogger(c *Config, lgr *slog.Logger) (*Client, error) {
	// TODO: Validate Config
	client := &Client{
		config:        *c,
		OnMessage:     nil,
		OnUnhandled:   nil,
		logger:        slog.New(logger.NewWrapHandler(lgr.Handler())),
		subscriptions: make(map[model.EventType]int),
		OnEntity:      make(model.OnEntityHandlers),
	}
	return client, nil
}

func defaultLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func (c *Client) SetLogger(logger *slog.Logger) {
	c.logger = logger
}
func (c *Client) Logger() *slog.Logger {
	return c.logger
}

func (c *Client) AddSubscription(eventType model.EventType) {
	c.subscriptions[eventType] = -1
}

func (c *Client) WriteMessage(messageType websocket.MessageType, data []byte) error {
	return c.client.Write(c.ctx, messageType, data)
}

// GetStates requests all the current states and runs them through the handlers
func (c *Client) GetStates() {
	callback := make(chan *model.Message)
	id := c.NextID()

	_, _ = c.sendWithCallback(&model.Message{
		ID:   &id,
		Type: model.MessageTypeGetStates,
	}, callback)
	go func() {

	}()

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

	return
}

func (c *Client) GetEntityRegistry() *model.Message {
	callback := make(chan *model.Message)
	id := c.NextID()

	_, _ = c.sendWithCallback(&model.Message{
		ID:   &id,
		Type: model.MessageTypeGetEntityRegistry,
	}, callback)
	return c.handleCallback(id, callback)
}

func (c *Client) handleCallback(id int, callback chan *model.Message) *model.Message {
	ticker := time.NewTicker(time.Second * 10)

	defer c.callbacks.Delete(id)
	select {
	case <-c.ctx.Done():
		return nil
	case <-ticker.C:
		// TODO: Make error message
		return nil
	case message := <-callback:
		return message
	}

	return nil
}
func (c *Client) GetDeviceRegistry() *model.Message {
	callback := make(chan *model.Message)
	id, _ := c.sendWithCallback(&model.Message{
		Type: model.MessageTypeGetDeviceRegistry,
	}, callback)
	return c.handleCallback(id, callback)
}

func (c *Client) GetServices() *model.Message {
	callback := make(chan *model.Message)
	id, _ := c.sendWithCallback(&model.Message{
		Type: model.MessageTypeGetServices,
	}, callback)
	return c.handleCallback(id, callback)
}

func (c *Client) CallService(service services.Service) *Response {
	id := c.NextID()
	service.SetID(&id)
	callback := make(chan *model.Message)

	rsp := NewResponse()

	err := c.sendStringWithCallback(service.JSON(), callback)
	if err != nil {
		c.logger.Error(err.Error())
		rsp.msg = &model.Message{
			Type: "",
			Error: &model.Error{
				Code:    "",
				Message: err.Error(),
			},
		}
		return rsp
	}
	go func(id int) {
		msg := c.handleCallback(id, callback)
		defer func() {
			rsp.done <- struct{}{}
		}()
		if msg != nil {
			if msg.Error != nil {
				c.logger.Error(fmt.Sprintf("service %s error code: %s message: %s", service.Name(), msg.Error.Code, msg.Error.Message), "Error sending")
			}
		}
		rsp.msg = msg

	}(id)
	return rsp
}

func (c *Client) NextID() int {
	c.id++
	return c.id
}

func (c *Client) read() (*model.Message, error) {
	if c.client == nil {
		return nil, errors.New("client is nil")
	}
	_, message, err := c.client.Read(c.ctx)
	if err != nil {
		return nil, err
	}

	var msg model.Message

	// messageString := string(message)

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
				// fmt.Println(messageString)
				return nil, err
			}

			return &msg, nil
		}

		return nil, fmt.Errorf("%s gave error: %w", string(message), err)
	}
	msg.Raw = message
	return &msg, nil
}

func (c *Client) Close() error {
	c.running = false
	// TODO: Chain errors
	var err error
	err = c.client.Close(websocket.StatusNormalClosure, "")
	// close the context
	c.cancel()
	if err != nil {
		return err
	}
	return nil
}
