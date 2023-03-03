package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewAutomationReload creates the object that can be sent to Home Assistant for domain automation, service reload
// "Reload the automation configuration."
func NewAutomationReload(target Target, automationReloadParams AutomationReloadParams) *AutomationReload {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "reload"
	a := &AutomationReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: automationReloadParams,
	}
	return a
}

type AutomationReload struct {
	ServiceBase
	ServiceData AutomationReloadParams `json:"service_data,omitempty"`
}
type AutomationReloadParams struct{}

func (a *AutomationReload) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AutomationReload) SetID(id *int) {
	a.Id = id
}

// NewAutomationToggle creates the object that can be sent to Home Assistant for domain automation, service toggle
// "Toggle (enable / disable) an automation."
func NewAutomationToggle(target Target, automationToggleParams AutomationToggleParams) *AutomationToggle {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "toggle"
	a := &AutomationToggle{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: automationToggleParams,
	}
	return a
}

type AutomationToggle struct {
	ServiceBase
	ServiceData AutomationToggleParams `json:"service_data,omitempty"`
}
type AutomationToggleParams struct{}

func (a *AutomationToggle) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AutomationToggle) SetID(id *int) {
	a.Id = id
}

// NewAutomationTrigger creates the object that can be sent to Home Assistant for domain automation, service trigger
// "Trigger the actions of an automation."
func NewAutomationTrigger(target Target, automationTriggerParams AutomationTriggerParams) *AutomationTrigger {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "trigger"
	a := &AutomationTrigger{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: automationTriggerParams,
	}
	return a
}

type AutomationTrigger struct {
	ServiceBase
	ServiceData AutomationTriggerParams `json:"service_data,omitempty"`
}
type AutomationTriggerParams struct{}

func (a *AutomationTrigger) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AutomationTrigger) SetID(id *int) {
	a.Id = id
}

// NewAutomationTurnOff creates the object that can be sent to Home Assistant for domain automation, service turn_off
// "Disable an automation."
func NewAutomationTurnOff(target Target, automationTurnOffParams AutomationTurnOffParams) *AutomationTurnOff {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "turn_off"
	a := &AutomationTurnOff{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: automationTurnOffParams,
	}
	return a
}

type AutomationTurnOff struct {
	ServiceBase
	ServiceData AutomationTurnOffParams `json:"service_data,omitempty"`
}
type AutomationTurnOffParams struct{}

func (a *AutomationTurnOff) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AutomationTurnOff) SetID(id *int) {
	a.Id = id
}

// NewAutomationTurnOn creates the object that can be sent to Home Assistant for domain automation, service turn_on
// "Enable an automation."
func NewAutomationTurnOn(target Target, automationTurnOnParams AutomationTurnOnParams) *AutomationTurnOn {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "turn_on"
	a := &AutomationTurnOn{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: automationTurnOnParams,
	}
	return a
}

type AutomationTurnOn struct {
	ServiceBase
	ServiceData AutomationTurnOnParams `json:"service_data,omitempty"`
}
type AutomationTurnOnParams struct{}

func (a *AutomationTurnOn) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AutomationTurnOn) SetID(id *int) {
	a.Id = id
}
