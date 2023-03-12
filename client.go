//go:generate go run ./helpers/

package hass_ws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kjbreil/hass-ws/model"
	"github.com/kjbreil/hass-ws/services"
	"nhooyr.io/websocket"
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

	ctx    context.Context
	cancel context.CancelFunc

	OnType     model.OnTypeHandlers
	OnEntity   model.OnEntityHandlers
	OnGetState func(states []model.Result)
}

func NewClient(c *Config) (*Client, error) {
	client := &Client{
		config:      *c,
		OnMessage:   nil,
		OnUnhandled: nil,

		subscriptions: make(map[model.EventType]int),
		OnEntity:      make(model.OnEntityHandlers),
	}
	return client, nil
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
	_, _ = c.sendWithCallback(&model.Message{
		Type: model.MessageTypeGetStates,
	}, callback)
	go func() {
		message := <-callback
		close(callback)
		if c.OnGetState != nil {
			c.OnGetState(message.Result)
		}

		c.OnType.RunStates(message)
		c.OnEntity.RunStates(message)
	}()
	return
}

func (c *Client) GetServices() *model.Message {
	callback := make(chan *model.Message)
	id, _ := c.sendWithCallback(&model.Message{
		Type: model.MessageTypeGetServices,
	}, callback)
	ticker := time.NewTicker(time.Second * 10)
	select {
	case <-ticker.C:
		c.callbacks.Delete(id)
		close(callback)

		return nil
	case message := <-callback:
		close(callback)
		return message
	}

}

func (c *Client) CallService(service services.Service) {
	service.SetID(c.NextID())
	callback := make(chan *model.Message)
	err := c.sendStringWithCallback(service.JSON(), callback)
	if err != nil {
		ERROR.Printf("Error sending: %v", err)
		return
	}
	go func() {
		message := <-callback
		if message.Error != nil {
			ERROR.Printf("service %s error code: %s message: %s", service.Name(), message.Error.Code, message.Error.Message)
		}
		//INFO.Printf("callback message received: %v", message)
		close(callback)
	}()
	return
}

func (c *Client) NextID() *int {
	c.id++
	return &c.id
}

func (c *Client) sendStringWithCallback(msg string, callback chan *model.Message) error {
	c.callbacks.Set(c.id, callback)

	err := c.client.Write(c.ctx, websocket.MessageText, []byte(msg))
	if err != nil {
		return err
	}
	return nil
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

	messageString := string(message)

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
				fmt.Println(messageString)
				return nil, err
			}

			return &msg, nil
		}

		fmt.Println(messageString)
		return nil, err
	}
	msg.Raw = message
	return &msg, nil
}

func (c *Client) Close() error {
	c.running = false
	// close the context
	c.cancel()
	// TODO: Chain errors
	var err error
	err = c.client.Close(websocket.StatusNormalClosure, "")
	if err != nil {
		return err
	}
	return nil
}
