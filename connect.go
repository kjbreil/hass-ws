package hass_ws

import (
	"fmt"
	"github.com/kjbreil/hass-ws/model"
	"net/url"
	"nhooyr.io/websocket"
	"time"
)

type Config struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Token string `json:"token"`
}

func (c *Client) Connect() error {
	var err error
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%d", c.config.Host, c.config.Port), Path: "/api/websocket"}

	c.client, _, err = websocket.Dial(c.ctx, u.String(), nil)
	c.client.SetReadLimit(32768 * 100)
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

	c.run()

	// this is needed because Home Assistant might not respond with authorization fast enough
	// TODO: make authorization into a callback
	time.Sleep(400 * time.Millisecond)
	return nil
}
