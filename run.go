package hass_ws

import "log"

func (c *Client) run() {
	for {
		message, err := c.read()
		if err != nil {
			// TODO: Handle Auth Error properly
			log.Panicln(err)
			return
		}
		if message.ID != nil {

			for i, callback := range c.callbacks {
				if i == *message.ID {
					delete(c.callbacks, i)
					callback <- message
					return
				}
			}
		}

		c.OnType.Run(message)
		c.OnEntity.Run(message)
	}
}
