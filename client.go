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

	callbacks map[int]chan *model.Message

	OnType   model.OnTypeHandlers
	OnEntity model.OnEntityHandlers
}

func NewClient(c *Config) (*Client, error) {

	client := &Client{
		config: *c,
		onMessage: func(message model.Message) {
		},
		subscriptions: make(map[model.EventType]int),
		OnEntity:      make(model.OnEntityHandlers),
		callbacks:     make(map[int]chan *model.Message),
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

// GetStates requests all the current states and runs them through the handlers
func (c *Client) GetStates() {
	callback := make(chan *model.Message)
	_ = c.sendWithCallback(&model.Message{
		Type: model.MessageTypeGetStates,
	}, callback)
	message := <-callback
	close(callback)

	c.OnType.RunStates(message)
	c.OnEntity.RunStates(message)
	return
}

//func (c *Client) GetServices() *model.Message {
//	callback := make(chan *model.Message)
//	_ = c.sendWithCallback(&model.Message{
//		Type: model.MessageTypeGetServices,
//	}, callback)
//	message := <-callback
//	close(callback)
//	os.WriteFile("services.json", message.Raw, 0666)
//	return message
//}

func (c *Client) sendWithCallback(msg *model.Message, callback chan *model.Message) error {
	if msg.Type != model.MessageTypeAuth {
		c.id++
		msg.ID = &c.id
		c.callbacks[c.id] = callback
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
