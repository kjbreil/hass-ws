package hass_ws

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/kjbreil/hass-ws/model"
	"nhooyr.io/websocket"
)

func (c *Client) sendWithCallback(msg *model.Message, callback chan *model.Message) (int, error) {

	if msg.ID == nil {
		id := c.NextID()
		msg.ID = &id
	}
	c.callbacks.Set(c.id, callback)

	d, _ := json.Marshal(msg)
	err := c.client.Write(c.ctx, websocket.MessageText, d)
	if err != nil {
		// if errors.Is(err, context.DeadlineExceeded) {
		//	c.client.Close(websocket.StatusNormalClosure, "")
		//	c.ctx, c.cancel = context.WithTimeout(context.Background(), 10*time.Second)
		//	c.Connect()
		// }
		return *msg.ID, err
	}
	if !msg.Type.Valid() {
		c.logger.Error(fmt.Sprintf("unknown message type: %s\n", msg.Type), "unknown message")
	}
	return *msg.ID, nil
}

func (c *Client) sendStringWithCallback(msg string, callback chan *model.Message) error {
	c.callbacks.Set(c.id, callback)

	err := c.client.Write(c.ctx, websocket.MessageText, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) send(msg *model.Message) error {
	if msg.Type != model.MessageTypeAuth {
		id := c.NextID()
		msg.ID = &id
	}

	d, _ := json.Marshal(msg)
	err := c.client.Write(c.ctx, websocket.MessageText, d)
	if err != nil {
		return err
	}
	if !msg.Type.Valid() {
		c.logger.Error(fmt.Sprintf("unknown message type: %s\n", msg.Type), "unknown message")
	}
	return nil
}
