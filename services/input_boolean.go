package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewInputBooleanReload creates the object that can be sent to Home Assistant for domain input_boolean, service reload
// "Reloads helpers from the YAML-configuration."
func NewInputBooleanReload(target Target) *InputBooleanReload {
	serviceDomain := "input_boolean"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputBooleanReload{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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

// NewInputBooleanToggle creates the object that can be sent to Home Assistant for domain input_boolean, service toggle
// "Toggles the helper on/off."
func NewInputBooleanToggle(target Target) *InputBooleanToggle {
	serviceDomain := "input_boolean"
	serviceType := "call_service"
	serviceService := "toggle"
	i := &InputBooleanToggle{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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

// NewInputBooleanTurnOff creates the object that can be sent to Home Assistant for domain input_boolean, service turn_off
// "Turns off the helper."
func NewInputBooleanTurnOff(target Target) *InputBooleanTurnOff {
	serviceDomain := "input_boolean"
	serviceType := "call_service"
	serviceService := "turn_off"
	i := &InputBooleanTurnOff{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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

// NewInputBooleanTurnOn creates the object that can be sent to Home Assistant for domain input_boolean, service turn_on
// "Turns on the helper."
func NewInputBooleanTurnOn(target Target) *InputBooleanTurnOn {
	serviceDomain := "input_boolean"
	serviceType := "call_service"
	serviceService := "turn_on"
	i := &InputBooleanTurnOn{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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
