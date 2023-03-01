package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewTextSetValue creates the object that can be sent to Home Assistant for domain text, service set_value
// "Set value of a text entity."
func NewTextSetValue(entities []string, value *string) *TextSetValue {
	serviceDomain := "text"
	serviceType := "call_service"
	serviceService := "set_value"
	t := &TextSetValue{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Value *string `json:"value,omitempty"`
		}{Value: value},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return t
}

type TextSetValue struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Value *string `json:"value,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (t *TextSetValue) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TextSetValue) SetID(id *int) {
	t.Id = id
}
