package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewInputButtonPress creates the object that can be sent to Home Assistant for domain input_button, service press
// "Press the input button entity."
func NewInputButtonPress(target Target) *InputButtonPress {
	serviceDomain := "input_button"
	serviceType := "call_service"
	serviceService := "press"
	i := &InputButtonPress{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return i
}

type InputButtonPress struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputButtonPress) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputButtonPress) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputButtonPress) SetID(id *int) {
	i.Id = id
}

// NewInputButtonReload creates the object that can be sent to Home Assistant for domain input_button, service reload
// ""
func NewInputButtonReload(target Target) *InputButtonReload {
	serviceDomain := "input_button"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputButtonReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return i
}

type InputButtonReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputButtonReload) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputButtonReload) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputButtonReload) SetID(id *int) {
	i.Id = id
}
