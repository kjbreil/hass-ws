package hass_ws

import (
	"github.com/kjbreil/hass-ws/model"
	"time"
)

type Response struct {
	msg     *model.Message
	done    chan struct{}
	timeout chan struct{}
}

func NewResponse() *Response {

	r := &Response{
		done: make(chan struct{}),
	}
	go func() {
		<-r.done
		defer func() {
			close(r.done)
			if _ = recover(); r != nil {
				return
			}
		}()
		if r.timeout != nil {
			r.timeout <- struct{}{}
		}
	}()

	return r
}

func (r *Response) Timeout(d time.Duration) *model.Message {
	r.timeout = make(chan struct{})
	defer close(r.timeout)
	select {
	case <-r.timeout:
		return r.msg
	case <-time.After(d):
		return nil
	}
	return nil
}
