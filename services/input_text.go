package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewInputTextReload creates the object that can be sent to Home Assistant for domain input_text, service reload
// "Reload the input_text configuration."
func NewInputTextReload(entities []string) *InputTextReload {
	serviceDomain := "input_text"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputTextReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return i
}

type InputTextReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputTextReload) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputTextReload) SetID(id *int) {
	i.Id = id
}

// NewInputTextSetValue creates the object that can be sent to Home Assistant for domain input_text, service set_value
// "Set the value of an input text entity."
func NewInputTextSetValue(entities []string, value *string) *InputTextSetValue {
	serviceDomain := "input_text"
	serviceType := "call_service"
	serviceService := "set_value"
	i := &InputTextSetValue{
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
	return i
}

type InputTextSetValue struct {
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

func (i *InputTextSetValue) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputTextSetValue) SetID(id *int) {
	i.Id = id
}
