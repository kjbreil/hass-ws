package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewButtonPress creates the object that can be sent to Home Assistant for domain button, service press
// "Press the button entity."
func NewButtonPress(target Target) *ButtonPress {
	serviceDomain := "button"
	serviceType := "call_service"
	serviceService := "press"
	b := &ButtonPress{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return b
}

type ButtonPress struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (b *ButtonPress) JSON() string {
	data, _ := gojson.Marshal(b)
	return string(data)
}
func (b *ButtonPress) Targets() []string {
	return b.Target.EntityId
}
func (b *ButtonPress) Name() string {
	return fmt.Sprintf("%s.%s", *b.Domain, *b.Service)
}
func (b *ButtonPress) SetID(id *int) {
	b.Id = id
}
