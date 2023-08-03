package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewInputBooleanReload creates the object that can be sent to Home Assistant for domain input_boolean, service reload
// "Reload the input_boolean configuration"
func NewInputBooleanReload(target Target) *InputBooleanReload {
	serviceDomain := "input_boolean"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputBooleanReload{
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

type InputBooleanReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputBooleanReload) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputBooleanReload) Targets() []string {
	return i.Target.EntityId
}
func (i *InputBooleanReload) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputBooleanReload) SetID(id *int) {
	i.Id = id
}

// NewInputBooleanToggle creates the object that can be sent to Home Assistant for domain input_boolean, service toggle
// "Toggle an input boolean"
func NewInputBooleanToggle(target Target) *InputBooleanToggle {
	serviceDomain := "input_boolean"
	serviceType := "call_service"
	serviceService := "toggle"
	i := &InputBooleanToggle{
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

type InputBooleanToggle struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputBooleanToggle) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputBooleanToggle) Targets() []string {
	return i.Target.EntityId
}
func (i *InputBooleanToggle) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputBooleanToggle) SetID(id *int) {
	i.Id = id
}

// NewInputBooleanTurnOff creates the object that can be sent to Home Assistant for domain input_boolean, service turn_off
// "Turn off an input boolean"
func NewInputBooleanTurnOff(target Target) *InputBooleanTurnOff {
	serviceDomain := "input_boolean"
	serviceType := "call_service"
	serviceService := "turn_off"
	i := &InputBooleanTurnOff{
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

type InputBooleanTurnOff struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputBooleanTurnOff) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputBooleanTurnOff) Targets() []string {
	return i.Target.EntityId
}
func (i *InputBooleanTurnOff) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputBooleanTurnOff) SetID(id *int) {
	i.Id = id
}

// NewInputBooleanTurnOn creates the object that can be sent to Home Assistant for domain input_boolean, service turn_on
// "Turn on an input boolean"
func NewInputBooleanTurnOn(target Target) *InputBooleanTurnOn {
	serviceDomain := "input_boolean"
	serviceType := "call_service"
	serviceService := "turn_on"
	i := &InputBooleanTurnOn{
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

type InputBooleanTurnOn struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputBooleanTurnOn) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputBooleanTurnOn) Targets() []string {
	return i.Target.EntityId
}
func (i *InputBooleanTurnOn) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputBooleanTurnOn) SetID(id *int) {
	i.Id = id
}
