package handlers

import (
	"github.com/kjbreil/hass-ws/model"
	"time"
)

type Response struct {
	msg     *model.Message
	Done    chan struct{}
	timeout chan struct{}
}

func NewResponse() *Response {

	r := &Response{
		Done: make(chan struct{}),
	}
	go func() {
		<-r.Done
		defer func() {
			close(r.Done)
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

// SetMessage sets the response message (used internally by client)
func (r *Response) SetMessage(msg *model.Message) {
	r.msg = msg
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
