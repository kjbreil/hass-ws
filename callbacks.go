package hass_ws

import (
	"github.com/kjbreil/hass-ws/model"
	"sync"
)

type callbacks struct {
	sync.Map
}

func newCallbacks() *callbacks {
	return &callbacks{}
}

func (c *callbacks) Get(key int) (chan *model.Message, bool) {
	ch, ok := c.Load(key)
	if ok {
		return ch.(chan *model.Message), true
	}
	return nil, false
}

func (c *callbacks) Set(key int, value chan *model.Message) {
	c.Store(key, value)
}

func (c *callbacks) GetMap() map[int]chan *model.Message {
	returnMap := make(map[int]chan *model.Message)

	c.Range(func(key, value any) bool {
		returnMap[key.(int)] = value.(chan *model.Message)
		return true
	})
	return returnMap
}

func (c *callbacks) Len() int {
	return len(c.GetMap())
}
