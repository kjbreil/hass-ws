package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewInputTextReload creates the object that can be sent to Home Assistant for domain input_text, service reload
// "Reload the input_text configuration."
func NewInputTextReload(target Target) *InputTextReload {
	serviceDomain := "input_text"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputTextReload{
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

type InputTextReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputTextReload) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputTextReload) Targets() []string {
	return i.Target.EntityId
}
func (i *InputTextReload) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputTextReload) SetID(id *int) {
	i.Id = id
}

// NewInputTextSetValue creates the object that can be sent to Home Assistant for domain input_text, service set_value
// "Set the value of an input text entity."
func NewInputTextSetValue(target Target) *InputTextSetValue {
	serviceDomain := "input_text"
	serviceType := "call_service"
	serviceService := "set_value"
	i := &InputTextSetValue{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: InputTextSetValueParams{},
	}
	return i
}

type InputTextSetValue struct {
	ServiceBase
	ServiceData InputTextSetValueParams `json:"service_data,omitempty"`
}
type InputTextSetValueParams struct {
	Value *string `json:"value,omitempty"`
}

func (i *InputTextSetValue) Value(value string) *InputTextSetValue {
	i.ServiceData.Value = &value
	return i
}
func (i *InputTextSetValue) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputTextSetValue) Targets() []string {
	return i.Target.EntityId
}
func (i *InputTextSetValue) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputTextSetValue) SetID(id *int) {
	i.Id = id
}
