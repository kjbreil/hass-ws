package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewAutomationReload creates the object that can be sent to Home Assistant for domain automation, service reload
// "Reload the automation configuration."
func NewAutomationReload(entities []string) *AutomationReload {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "reload"
	a := &AutomationReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AutomationReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AutomationReload) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AutomationReload) SetID(id *int) {
	a.Id = id
}

// NewAutomationToggle creates the object that can be sent to Home Assistant for domain automation, service toggle
// "Toggle (enable / disable) an automation."
func NewAutomationToggle(entities []string) *AutomationToggle {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "toggle"
	a := &AutomationToggle{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AutomationToggle struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AutomationToggle) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AutomationToggle) SetID(id *int) {
	a.Id = id
}

// NewAutomationTrigger creates the object that can be sent to Home Assistant for domain automation, service trigger
// "Trigger the actions of an automation."
func NewAutomationTrigger(entities []string) *AutomationTrigger {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "trigger"
	a := &AutomationTrigger{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AutomationTrigger struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AutomationTrigger) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AutomationTrigger) SetID(id *int) {
	a.Id = id
}

// NewAutomationTurnOff creates the object that can be sent to Home Assistant for domain automation, service turn_off
// "Disable an automation."
func NewAutomationTurnOff(entities []string) *AutomationTurnOff {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "turn_off"
	a := &AutomationTurnOff{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AutomationTurnOff struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AutomationTurnOff) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AutomationTurnOff) SetID(id *int) {
	a.Id = id
}

// NewAutomationTurnOn creates the object that can be sent to Home Assistant for domain automation, service turn_on
// "Enable an automation."
func NewAutomationTurnOn(entities []string) *AutomationTurnOn {
	serviceDomain := "automation"
	serviceType := "call_service"
	serviceService := "turn_on"
	a := &AutomationTurnOn{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AutomationTurnOn struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AutomationTurnOn) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AutomationTurnOn) SetID(id *int) {
	a.Id = id
}
