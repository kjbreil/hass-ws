package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCounterDecrement creates the object that can be sent to Home Assistant for domain counter, service decrement
// "Decrements a counter by its step size."
func NewCounterDecrement(target Target) *CounterDecrement {
	serviceDomain := "counter"
	serviceType := "call_service"
	serviceService := "decrement"
	c := &CounterDecrement{
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
	return c
}

type CounterDecrement struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CounterDecrement) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CounterDecrement) Targets() []string {
	return c.Target.EntityId
}
func (c *CounterDecrement) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCounterIncrement creates the object that can be sent to Home Assistant for domain counter, service increment
// "Increments a counter by its step size."
func NewCounterIncrement(target Target) *CounterIncrement {
	serviceDomain := "counter"
	serviceType := "call_service"
	serviceService := "increment"
	c := &CounterIncrement{
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
	return c
}

type CounterIncrement struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CounterIncrement) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CounterIncrement) Targets() []string {
	return c.Target.EntityId
}
func (c *CounterIncrement) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCounterReset creates the object that can be sent to Home Assistant for domain counter, service reset
// "Resets a counter to its initial value."
func NewCounterReset(target Target) *CounterReset {
	serviceDomain := "counter"
	serviceType := "call_service"
	serviceService := "reset"
	c := &CounterReset{
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
	return c
}

type CounterReset struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CounterReset) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CounterReset) Targets() []string {
	return c.Target.EntityId
}
func (c *CounterReset) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCounterSetValue creates the object that can be sent to Home Assistant for domain counter, service set_value
// "Sets the counter to a specific value."
func NewCounterSetValue(target Target) *CounterSetValue {
	serviceDomain := "counter"
	serviceType := "call_service"
	serviceService := "set_value"
	c := &CounterSetValue{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: CounterSetValueParams{},
	}
	return c
}

type CounterSetValue struct {
	ServiceBase
	ServiceData CounterSetValueParams `json:"service_data,omitempty"`
}
type CounterSetValueParams struct {
	Value *float64 `json:"value,omitempty"`
}

func (c *CounterSetValue) Value(value float64) *CounterSetValue {
	c.ServiceData.Value = &value
	return c
}
func (c *CounterSetValue) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CounterSetValue) Targets() []string {
	return c.Target.EntityId
}
func (c *CounterSetValue) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}
