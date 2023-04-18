package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewInputNumberDecrement creates the object that can be sent to Home Assistant for domain input_number, service decrement
// "Decrement the value of an input number entity by its stepping."
func NewInputNumberDecrement(target Target) *InputNumberDecrement {
	serviceDomain := "input_number"
	serviceType := "call_service"
	serviceService := "decrement"
	i := &InputNumberDecrement{
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

type InputNumberDecrement struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputNumberDecrement) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputNumberDecrement) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputNumberDecrement) SetID(id *int) {
	i.Id = id
}

// NewInputNumberIncrement creates the object that can be sent to Home Assistant for domain input_number, service increment
// "Increment the value of an input number entity by its stepping."
func NewInputNumberIncrement(target Target) *InputNumberIncrement {
	serviceDomain := "input_number"
	serviceType := "call_service"
	serviceService := "increment"
	i := &InputNumberIncrement{
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

type InputNumberIncrement struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputNumberIncrement) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputNumberIncrement) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputNumberIncrement) SetID(id *int) {
	i.Id = id
}

// NewInputNumberReload creates the object that can be sent to Home Assistant for domain input_number, service reload
// "Reload the input_number configuration."
func NewInputNumberReload(target Target) *InputNumberReload {
	serviceDomain := "input_number"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputNumberReload{
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

type InputNumberReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputNumberReload) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputNumberReload) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputNumberReload) SetID(id *int) {
	i.Id = id
}

// NewInputNumberSetValue creates the object that can be sent to Home Assistant for domain input_number, service set_value
// "Set the value of an input number entity."
func NewInputNumberSetValue(target Target) *InputNumberSetValue {
	serviceDomain := "input_number"
	serviceType := "call_service"
	serviceService := "set_value"
	i := &InputNumberSetValue{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: InputNumberSetValueParams{},
	}
	return i
}

type InputNumberSetValue struct {
	ServiceBase
	ServiceData InputNumberSetValueParams `json:"service_data,omitempty"`
}
type InputNumberSetValueParams struct {
	Value *float64 `json:"value,omitempty"`
}

func (i *InputNumberSetValue) Value(value float64) *InputNumberSetValue {
	i.ServiceData.Value = &value
	return i
}
func (i *InputNumberSetValue) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputNumberSetValue) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputNumberSetValue) SetID(id *int) {
	i.Id = id
}
