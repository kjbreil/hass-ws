package hass_ws

import (
	"github.com/kjbreil/hass-ws/model"
	"log"
	"time"
)

func (c *Client) run() {
	go func() {
	mainLoop:
		for {
			message, err := c.read()
			if err != nil {
				// TODO: Handle Auth Error properly
				log.Println(err)
				continue
			}
			if c.OnMessage != nil {
				c.OnMessage(*message)
			}

			if message.ID != nil {
				for i, callback := range c.callbacks {
					if i == *message.ID {
						delete(c.callbacks, i)
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
			case <-ticker.C:

				callback := make(chan *model.Message)
				_, err := c.sendWithCallback(&model.Message{
					Type: model.MessageTypePing,
				}, callback)
				if err != nil {
					log.Panicln(err)
				}
				go func() {
					<-callback
				}()
			}
		}
	}()

}
