package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewConversationProcess creates the object that can be sent to Home Assistant for domain conversation, service process
// "Launch a conversation from a transcribed text."
func NewConversationProcess(entities []string, text *string) *ConversationProcess {
	serviceDomain := "conversation"
	serviceType := "call_service"
	serviceService := "process"
	c := &ConversationProcess{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Text *string `json:"text,omitempty"`
		}{Text: text},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type ConversationProcess struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Text *string `json:"text,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewConversationReload(entities []string) *ConversationReload {
	serviceDomain := "conversation"
	serviceType := "call_service"
	serviceService := "reload"
	c := &ConversationReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type ConversationReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *ConversationReload) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ConversationReload) SetID(id *int) {
	c.Id = id
}
