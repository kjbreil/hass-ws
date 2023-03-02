package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewInputButtonPress creates the object that can be sent to Home Assistant for domain input_button, service press
// "Press the input button entity."
func NewInputButtonPress(entities []string) *InputButtonPress {
	serviceDomain := "input_button"
	serviceType := "call_service"
	serviceService := "press"
	i := &InputButtonPress{
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

type InputButtonPress struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputButtonPress) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputButtonPress) SetID(id *int) {
	i.Id = id
}

// NewInputButtonReload creates the object that can be sent to Home Assistant for domain input_button, service reload
// ""
func NewInputButtonReload(entities []string) *InputButtonReload {
	serviceDomain := "input_button"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputButtonReload{
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

type InputButtonReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputButtonReload) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputButtonReload) SetID(id *int) {
	i.Id = id
}
