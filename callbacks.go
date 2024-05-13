package hass_ws

import (
	"github.com/kjbreil/hass-ws/model"
	"sync"
)

type callbacks struct {
	callbacks map[int]chan *model.Message
	m         *sync.Mutex
}

func newCallbacks() *callbacks {
	return &callbacks{
		callbacks: make(map[int]chan *model.Message),
		m:         &sync.Mutex{},
	}
}

func (c *callbacks) Get(key int) (chan *model.Message, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	if cb, ok := c.callbacks[key]; ok {
		return cb, true
	}
	return nil, false
}

func (c *callbacks) Set(key int, value chan *model.Message) {
	c.m.Lock()
	defer c.m.Unlock()
	c.callbacks[key] = value
}
func (c *callbacks) Delete(key int) {
	c.m.Lock()
	defer c.m.Unlock()
	if cb, ok := c.callbacks[key]; ok {
		close(cb)
		delete(c.callbacks, key)
	}
}

func (c *callbacks) Len() int {
	c.m.Lock()
	defer c.m.Unlock()
	return len(c.callbacks)
}
