package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewAutomationReload creates the object that can be sent to Home Assistant for domain automation, service reload
// "Reload the automation configuration."
func NewAutomationReload(target Target) *AutomationReload {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "reload"
	a := &AutomationReload{
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
	return a
}

type AutomationReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (a *AutomationReload) JSON() string {
	data, _ := gojson.Marshal(a)
	return string(data)
}
func (a *AutomationReload) Targets() []string {
	return a.Target.EntityId
}
func (a *AutomationReload) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}

// NewAutomationToggle creates the object that can be sent to Home Assistant for domain automation, service toggle
// "Toggle (enable / disable) an automation."
func NewAutomationToggle(target Target) *AutomationToggle {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "toggle"
	a := &AutomationToggle{
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
	return a
}

type AutomationToggle struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (a *AutomationToggle) JSON() string {
	data, _ := gojson.Marshal(a)
	return string(data)
}
func (a *AutomationToggle) Targets() []string {
	return a.Target.EntityId
}
func (a *AutomationToggle) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}

// NewAutomationTrigger creates the object that can be sent to Home Assistant for domain automation, service trigger
// "Trigger the actions of an automation."
func NewAutomationTrigger(target Target) *AutomationTrigger {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "trigger"
	a := &AutomationTrigger{
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
	return a
}

type AutomationTrigger struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (a *AutomationTrigger) JSON() string {
	data, _ := gojson.Marshal(a)
	return string(data)
}
func (a *AutomationTrigger) Targets() []string {
	return a.Target.EntityId
}
func (a *AutomationTrigger) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}

// NewAutomationTurnOff creates the object that can be sent to Home Assistant for domain automation, service turn_off
// "Disable an automation."
func NewAutomationTurnOff(target Target) *AutomationTurnOff {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "turn_off"
	a := &AutomationTurnOff{
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
	return a
}

type AutomationTurnOff struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (a *AutomationTurnOff) JSON() string {
	data, _ := gojson.Marshal(a)
	return string(data)
}
func (a *AutomationTurnOff) Targets() []string {
	return a.Target.EntityId
}
func (a *AutomationTurnOff) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}

// NewAutomationTurnOn creates the object that can be sent to Home Assistant for domain automation, service turn_on
// "Enable an automation."
func NewAutomationTurnOn(target Target) *AutomationTurnOn {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "turn_on"
	a := &AutomationTurnOn{
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
	return a
}

type AutomationTurnOn struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (a *AutomationTurnOn) JSON() string {
	data, _ := gojson.Marshal(a)
	return string(data)
}
func (a *AutomationTurnOn) Targets() []string {
	return a.Target.EntityId
}
func (a *AutomationTurnOn) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}
