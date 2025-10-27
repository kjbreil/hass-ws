package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewConversationProcess creates the object that can be sent to Home Assistant for domain conversation, service process
// "Launch a conversation from a transcribed text."
func NewConversationProcess(target Target) *ConversationProcess {
	serviceDomain := "conversation"
	serviceType := "call_service"
	serviceService := "process"
	c := &ConversationProcess{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: ConversationProcessParams{},
	}
	return c
}

type ConversationProcess struct {
	ServiceBase
	ServiceData ConversationProcessParams `json:"service_data,omitempty"`
}
type ConversationProcessParams struct {
	Text *string `json:"text,omitempty"`
}

func (c *ConversationProcess) Text(text string) *ConversationProcess {
	c.ServiceData.Text = &text
	return c
}
func (c *ConversationProcess) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *ConversationProcess) Targets() []string {
	return c.Target.EntityId
}
func (c *ConversationProcess) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewConversationReload creates the object that can be sent to Home Assistant for domain conversation, service reload
// ""
func NewConversationReload(target Target) *ConversationReload {
	serviceDomain := "conversation"
	serviceType := "call_service"
	serviceService := "reload"
	c := &ConversationReload{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type ConversationReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *ConversationReload) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *ConversationReload) Targets() []string {
	return c.Target.EntityId
}
func (c *ConversationReload) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}
