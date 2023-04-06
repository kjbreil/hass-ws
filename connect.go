package hass_ws

import (
	"context"
	"errors"
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/kjbreil/hass-ws/model"
	"net/url"
	"nhooyr.io/websocket"
	"os"
	"time"
)

type Config struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Token string `json:"token"`
}

func ParseConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var c struct {
		Websocket Config
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	return &c.Websocket, nil
}

func (c *Client) Connect() error {

	if c.cancel != nil {
		c.cancel()
	}

	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.callbacks = newCallbacks()
	c.id = 0

	var err error
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%d", c.config.Host, c.config.Port), Path: "/api/websocket"}

	c.client, _, err = websocket.Dial(c.ctx, u.String(), nil)

	if err != nil {
		c.cancel()
		return fmt.Errorf("dial: %w", err)
	}

	c.client.SetReadLimit(32768 * 100)

	// Mark as running
	c.running = true

	msg := &model.Message{}
	msg, err = c.read()
	if err != nil {
		c.cancel()
		return fmt.Errorf("read: %w", err)
	}
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

	if c.InitStates {
		c.GetStates()
	}

	return nil
}

func (c *Client) reconnect() error {
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
			return nil
		}
		c.logger.Error(fmt.Errorf("reconnect failed: %v", err), "reconnect failed")
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return fmt.Errorf("websocket reconnect failed: %w", err)
	}
	return nil
}
