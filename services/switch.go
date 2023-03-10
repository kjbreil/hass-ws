package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSwitchToggle creates the object that can be sent to Home Assistant for domain switch, service toggle
// "Toggles a switch state"
func NewSwitchToggle(target Target) *SwitchToggle {
	serviceDomain := "switch"
	serviceType := "call_service"
	serviceService := "toggle"
	s := &SwitchToggle{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return s
}

type SwitchToggle struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *SwitchToggle) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SwitchToggle) SetID(id *int) {
	s.Id = id
}

// NewSwitchTurnOff creates the object that can be sent to Home Assistant for domain switch, service turn_off
// "Turn a switch off"
func NewSwitchTurnOff(target Target) *SwitchTurnOff {
	serviceDomain := "switch"
	serviceType := "call_service"
	serviceService := "turn_off"
	s := &SwitchTurnOff{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return s
}

type SwitchTurnOff struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *SwitchTurnOff) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SwitchTurnOff) SetID(id *int) {
	s.Id = id
}

// NewSwitchTurnOn creates the object that can be sent to Home Assistant for domain switch, service turn_on
// "Turn a switch on"
func NewSwitchTurnOn(target Target) *SwitchTurnOn {
	serviceDomain := "switch"
	serviceType := "call_service"
	serviceService := "turn_on"
	s := &SwitchTurnOn{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return s
}

type SwitchTurnOn struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *SwitchTurnOn) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SwitchTurnOn) SetID(id *int) {
	s.Id = id
}
