package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewButtonPress creates the object that can be sent to Home Assistant for domain button, service press
// "Press the button entity."
func NewButtonPress(entities []string) *ButtonPress {
	serviceDomain := "button"
	serviceType := "call_service"
	serviceService := "press"
	b := &ButtonPress{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return b
}

type ButtonPress struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (b *ButtonPress) JSON() string {
	data, _ := json.Marshal(b)
	return string(data)
}
func (b *ButtonPress) SetID(id *int) {
	b.Id = id
}
