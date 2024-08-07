package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCounterConfigure creates the object that can be sent to Home Assistant for domain counter, service configure
// "Change counter parameters."
func NewCounterConfigure(target Target) *CounterConfigure {
	serviceDomain := "counter"
	serviceType := "call_service"
	serviceService := "configure"
	c := &CounterConfigure{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: CounterConfigureParams{},
	}
	return c
}

type CounterConfigure struct {
	ServiceBase
	ServiceData CounterConfigureParams `json:"service_data,omitempty"`
}
type CounterConfigureParams struct {
	Initial *float64 `json:"initial,omitempty"`
	Maximum *float64 `json:"maximum,omitempty"`
	Minimum *float64 `json:"minimum,omitempty"`
	Step    *float64 `json:"step,omitempty"`
	Value   *float64 `json:"value,omitempty"`
}

func (c *CounterConfigure) Initial(initial float64) *CounterConfigure {
	c.ServiceData.Initial = &initial
	return c
}
func (c *CounterConfigure) Maximum(maximum float64) *CounterConfigure {
	c.ServiceData.Maximum = &maximum
	return c
}
func (c *CounterConfigure) Minimum(minimum float64) *CounterConfigure {
	c.ServiceData.Minimum = &minimum
	return c
}
func (c *CounterConfigure) Step(step float64) *CounterConfigure {
	c.ServiceData.Step = &step
	return c
}
func (c *CounterConfigure) Value(value float64) *CounterConfigure {
	c.ServiceData.Value = &value
	return c
}
func (c *CounterConfigure) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CounterConfigure) Targets() []string {
	return c.Target.EntityId
}
func (c *CounterConfigure) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCounterDecrement creates the object that can be sent to Home Assistant for domain counter, service decrement
// "Decrement a counter."
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
// "Increment a counter."
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
// "Reset a counter."
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
