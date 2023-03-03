package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewConversationProcess creates the object that can be sent to Home Assistant for domain conversation, service process
// "Launch a conversation from a transcribed text."
func NewConversationProcess(target Target, conversationProcessParams *ConversationProcessParams) *ConversationProcess {
	serviceDomain := "conversation"
	serviceType := "call_service"
	serviceService := "process"
	c := &ConversationProcess{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *conversationProcessParams,
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

func (c *ConversationProcess) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ConversationProcess) SetID(id *int) {
	c.Id = id
}

// NewConversationReload creates the object that can be sent to Home Assistant for domain conversation, service reload
// ""
func NewConversationReload(target Target) *ConversationReload {
	serviceDomain := "conversation"
	serviceType := "call_service"
	serviceService := "reload"
	c := &ConversationReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
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
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ConversationReload) SetID(id *int) {
	c.Id = id
}
