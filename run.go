package hass_ws

import (
	"context"
	"errors"
	"fmt"
	"github.com/kjbreil/hass-ws/model"
	"nhooyr.io/websocket"
	"time"
)

func (c *Client) run() {
	go func() {
	mainLoop:
		for {
			if c.ctx.Err() != nil {
				return
			}
			message, err := c.read()
			if err != nil {
				c.logger.Error(fmt.Errorf("read error: %v", err), "message read error")
				if errors.Is(err, context.Canceled) {
					return
				}

				if status := websocket.CloseStatus(err); status != -1 {
					return
				}
				continue
			}
			if c.OnMessage != nil {
				c.OnMessage(*message)
			}

			if message.ID != nil {
				for i, callback := range c.callbacks.GetMap() {
					if i == *message.ID {
						c.callbacks.Delete(i)
						callback <- message
						continue mainLoop
					}
				}
			}

			onTypeRan := c.OnType.Run(message)
			onEntityRan := c.OnEntity.Run(message)

			if !onTypeRan && !onEntityRan && c.OnUnhandled != nil {
				c.OnUnhandled(*message)
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(time.Second * 10)

		for {
			select {
			case <-c.ctx.Done():
				return
			case <-ticker.C:

				callback := make(chan *model.Message)
				_, err := c.sendWithCallback(&model.Message{
					Type: model.MessageTypePing,
				}, callback)
				if err != nil {
					c.logger.Error(fmt.Errorf("ping send error: %v", err), "ping could not be sent")
					err = c.reconnect()
					if err != nil {
						c.logger.Error(fmt.Errorf("reconnect failed: %v", err), "reconnect failed")
						err = c.Close()
						if err != nil {
							c.logger.Error(fmt.Errorf("close error: %v", err), "context close error")
						}
						panic(errors.New("connection lost and cannot be reconnected"))
					}
					return
				}
				go func() {
					// ping needs to come back within a second or connection needs restarting
					restartTicker := time.NewTicker(time.Second)
					select {
					case <-callback:
					case <-restartTicker.C:
						c.logger.Error(fmt.Errorf("pong not received attempting reconnect"), "ping could not be sent")

						err := c.reconnect()
						if err != nil {
							c.logger.Error(fmt.Errorf("reconnect failed: %v", err), "reconnect failed")
							err = c.Close()
							if err != nil {
								c.logger.Error(fmt.Errorf("close error: %v", err), "context close error")
							}
							panic(errors.New("connection lost and cannot be reconnected"))
						}
					}
				}()
			}
		}
	}()

}
