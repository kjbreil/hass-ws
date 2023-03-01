package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewSwitchToggle creates the object that can be sent to Home Assistant for domain switch, service toggle
// "Toggles a switch state"
func NewSwitchToggle(entities []string) *SwitchToggle {
	serviceDomain := "switch"
	serviceType := "call_service"
	serviceService := "toggle"
	s := &SwitchToggle{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SwitchToggle struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewSwitchTurnOff(entities []string) *SwitchTurnOff {
	serviceDomain := "switch"
	serviceType := "call_service"
	serviceService := "turn_off"
	s := &SwitchTurnOff{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SwitchTurnOff struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewSwitchTurnOn(entities []string) *SwitchTurnOn {
	serviceDomain := "switch"
	serviceType := "call_service"
	serviceService := "turn_on"
	s := &SwitchTurnOn{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SwitchTurnOn struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SwitchTurnOn) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SwitchTurnOn) SetID(id *int) {
	s.Id = id
}
