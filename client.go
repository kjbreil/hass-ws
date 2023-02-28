package hass_ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/kjbreil/hass-ws/model"
	"log"
)

type Client struct {
	client *websocket.Conn
	config Config
	id     int

	subscriptions map[model.EventType]int
	onMessage     func(message model.Message)

	eventWatch map[string]func(event model.Event)

	OnType   model.OnTypeHandlers
	OnEntity model.OnEntityHandlers
}

func NewClient(c *Config) (*Client, error) {

	client := &Client{
		config: *c,
		onMessage: func(message model.Message) {
			log.Panicln("onMessage handler not setup")
		},
		subscriptions: make(map[model.EventType]int),
		OnEntity:      make(model.OnEntityHandlers),
	}
	return client, nil
}

func (c *Client) SetOnMessage(handlerFunc func(message model.Message)) {
	c.onMessage = handlerFunc
}

func (c *Client) AddSubscription(eventType model.EventType) {
	c.subscriptions[eventType] = -1
}

func (c *Client) WriteMessage(messageType int, data []byte) error {
	return c.client.WriteMessage(messageType, data)
}

func (c *Client) send(msg *model.Message) error {
	if msg.Type != model.MessageTypeAuth {
		c.id++
		msg.ID = &c.id
	}

	d, _ := json.Marshal(msg)
	err := c.client.WriteMessage(websocket.TextMessage, d)
	if err != nil {
		return err
	}
	if !msg.Type.Valid() {
		log.Panicf("unknown message type: %s\n", msg.Type)
	}
	return nil
}

func (c *Client) read() (*model.Message, error) {
	_, message, err := c.client.ReadMessage()
	if err != nil {
		return nil, err
	}
	var msg model.Message

	err = json.Unmarshal(message, &msg)
	if err != nil {
		return nil, err
	}
	msg.Raw = message
	return &msg, nil
}

func (c *Client) Close() error {
	var err error
	err = c.client.Close()
	if err != nil {
		return err
	}
	return nil
}
