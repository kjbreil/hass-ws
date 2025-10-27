package callbacks

import (
	"github.com/kjbreil/hass-ws/model"
	"sync"
)

type Callbacks struct {
	callbacks map[int]chan *model.Message
	m         *sync.Mutex
}

func New() *Callbacks {
	return &Callbacks{
		callbacks: make(map[int]chan *model.Message),
		m:         &sync.Mutex{},
	}
}

func (c *Callbacks) Get(key int) (chan *model.Message, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	if cb, ok := c.callbacks[key]; ok {
		return cb, true
	}
	return nil, false
}

func (c *Callbacks) Set(key int, value chan *model.Message) {
	c.m.Lock()
	defer c.m.Unlock()
	c.callbacks[key] = value
}
func (c *Callbacks) Delete(key int) {
	c.m.Lock()
	defer c.m.Unlock()
	if cb, ok := c.callbacks[key]; ok {
		close(cb)
		delete(c.callbacks, key)
	}
}

func (c *Callbacks) Len() int {
	c.m.Lock()
	defer c.m.Unlock()
	return len(c.callbacks)
}
