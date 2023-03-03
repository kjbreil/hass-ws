package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSirenToggle creates the object that can be sent to Home Assistant for domain siren, service toggle
// "Toggles a siren."
func NewSirenToggle(target Target, sirenToggleParams SirenToggleParams) *SirenToggle {
	serviceDomain := "siren"
	serviceType := "call_service"
	serviceService := "toggle"
	s := &SirenToggle{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: sirenToggleParams,
	}
	return s
}

type SirenToggle struct {
	ServiceBase
	ServiceData SirenToggleParams `json:"service_data,omitempty"`
}
type SirenToggleParams struct{}

func (s *SirenToggle) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SirenToggle) SetID(id *int) {
	s.Id = id
}

// NewSirenTurnOff creates the object that can be sent to Home Assistant for domain siren, service turn_off
// "Turn siren off."
func NewSirenTurnOff(target Target, sirenTurnOffParams SirenTurnOffParams) *SirenTurnOff {
	serviceDomain := "siren"
	serviceType := "call_service"
	serviceService := "turn_off"
	s := &SirenTurnOff{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: sirenTurnOffParams,
	}
	return s
}

type SirenTurnOff struct {
	ServiceBase
	ServiceData SirenTurnOffParams `json:"service_data,omitempty"`
}
type SirenTurnOffParams struct{}

func (s *SirenTurnOff) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SirenTurnOff) SetID(id *int) {
	s.Id = id
}

// NewSirenTurnOn creates the object that can be sent to Home Assistant for domain siren, service turn_on
// "Turn siren on."
func NewSirenTurnOn(target Target, sirenTurnOnParams SirenTurnOnParams) *SirenTurnOn {
	serviceDomain := "siren"
	serviceType := "call_service"
	serviceService := "turn_on"
	s := &SirenTurnOn{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: sirenTurnOnParams,
	}
	return s
}

type SirenTurnOn struct {
	ServiceBase
	ServiceData SirenTurnOnParams `json:"service_data,omitempty"`
}
type SirenTurnOnParams struct {
	Duration    *string  `json:"duration,omitempty"`
	Tone        *string  `json:"tone,omitempty"`
	VolumeLevel *float64 `json:"volume_level,omitempty"`
}

func (s *SirenTurnOn) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SirenTurnOn) SetID(id *int) {
	s.Id = id
}
