package hass_ws

import (
	"context"
	"errors"
	"fmt"
	"github.com/kjbreil/hass-ws/model"
	"io"
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
				c.logger.Error(fmt.Sprintf("read error: %v", err), "message read error")
				switch {
				case errors.Is(err, context.Canceled),
					websocket.CloseStatus(err) != -1,
					errors.Is(err, io.EOF),
					errors.Is(err, io.ErrUnexpectedEOF):
					return
				default:
					continue
				}
			}
			if c.OnMessage != nil {
				c.OnMessage(*message)
			}

			if message.ID != nil {
				if callback, ok := c.callbacks.Get(*message.ID); ok {
					// TODO: Investigate callback being closed
					callback <- message
					continue mainLoop
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
				id, err := c.sendWithCallback(&model.Message{
					Type: model.MessageTypePing,
				}, callback)
				if err != nil {
					c.logger.Error(fmt.Sprintf("ping send error: %v", err), "ping could not be sent")
					err = c.reconnect()
					if err != nil {
						c.logger.Error(fmt.Sprintf("reconnect failed: %v", err), "reconnect failed")
						err = c.Close()
						if err != nil {
							c.logger.Error(fmt.Sprintf("close error: %v", err), "context close error")
						}
						panic(errors.New("connection lost and cannot be reconnected"))
					}
					c.logger.Info("reconnect succeeded")
					return
				}
				go func(id int) {
					defer c.callbacks.Delete(id)
					// ping needs to come back within a second or connection needs restarting
					restartTicker := time.NewTicker(time.Second * 5)
					select {
					case <-callback:
					case <-restartTicker.C:
						c.logger.Error(fmt.Sprintf("pong not received attempting reconnect"))

						err := c.reconnect()
						if err != nil {
							c.logger.Error(fmt.Sprintf("reconnect failed: %v", err))
							err = c.Close()
							if err != nil {
								c.logger.Error(fmt.Sprintf("close error: %v", err))
							}
							panic(errors.New("connection lost and cannot be reconnected"))
						}
						c.logger.Info("reconnect succeeded")

					}
				}(id)
			}
		}
	}()

}
