package hass_ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kjbreil/hass-ws/model"
	"log"
	"net/url"
)

type Client struct {
	client *websocket.Conn
	config Config
	id     int

	subscriptions map[model.EventType]int
	onMessage     func(message model.Message)
}

type Config struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Token string `json:"token"`
}

func NewClient(c *Config) (*Client, error) {

	client := &Client{
		config: *c,
		onMessage: func(message model.Message) {
			if message.Type == model.MessageTypeEvent {
				if message.Event == nil && message.Event.EventType != nil && !message.Event.EventType.Valid() {
					fmt.Printf("invalid EventType\n%s\n", message.Raw)
				}
			}
			//log.Printf("Received message: %s", message.Raw)
		},
		subscriptions: make(map[model.EventType]int),
	}
	return client, nil
}

func (c *Client) AddSubscription(eventType model.EventType) {
	c.subscriptions[eventType] = -1
}

func (c *Client) Connect() error {
	var err error
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%d", c.config.Host, c.config.Port), Path: "/api/websocket"}

	c.client, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	msg := &model.Message{}
	msg, err = c.read()
	if msg.Type != model.MessageTypeAuthRequired {
		return fmt.Errorf("initial response not AuthRequired: %s", msg.Raw)
	}
	err = c.send(&model.Message{
		Type:        model.MessageTypeAuth,
		AccessToken: &c.config.Token,
	})
	if err != nil {
		return err
	}

	for eventType := range c.subscriptions {
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

	go func() {
		for {
			read, err := c.read()
			if err != nil {
				log.Panicln(err)
				return
			}
			c.onMessage(*read)
		}
	}()

	return nil
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
