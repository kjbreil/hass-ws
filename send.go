package hass_ws

import (
	"encoding/json"
	"github.com/kjbreil/hass-ws/model"
	"log"
	"nhooyr.io/websocket"
)

func (c *Client) sendWithCallback(msg *model.Message, callback chan *model.Message) (int, error) {
	if msg.Type != model.MessageTypeAuth {
		msg.ID = c.NextID()
		c.callbacks.Set(c.id, callback)

	}

	d, _ := json.Marshal(msg)
	err := c.client.Write(c.ctx, websocket.MessageText, d)
	if err != nil {
		//if errors.Is(err, context.DeadlineExceeded) {
		//	c.client.Close(websocket.StatusNormalClosure, "")
		//	c.ctx, c.cancel = context.WithTimeout(context.Background(), 10*time.Second)
		//	c.Connect()
		//}
		return *msg.ID, err
	}
	if !msg.Type.Valid() {
		log.Panicf("unknown message type: %s\n", msg.Type)
	}
	return *msg.ID, nil
}

func (c *Client) send(msg *model.Message) error {
	if msg.Type != model.MessageTypeAuth {
		msg.ID = c.NextID()
	}

	d, _ := json.Marshal(msg)
	err := c.client.Write(c.ctx, websocket.MessageText, d)
	if err != nil {
		return err
	}
	if !msg.Type.Valid() {
		log.Panicf("unknown message type: %s\n", msg.Type)
	}
	return nil
}
