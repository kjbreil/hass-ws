package hass_ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kjbreil/hass-ws/model"
	"log"
	"net/url"
)

type Config struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Token string `json:"token"`
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
