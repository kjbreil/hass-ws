package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSirenToggle creates the object that can be sent to Home Assistant for domain siren, service toggle
// "Toggles a siren."
func NewSirenToggle(target Target) *SirenToggle {
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
		ServiceData: nil,
	}
	return s
}

type SirenToggle struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *SirenToggle) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SirenToggle) Targets() []string {
	return s.Target.EntityId
}
func (s *SirenToggle) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *SirenToggle) SetID(id *int) {
	s.Id = id
}

// NewSirenTurnOff creates the object that can be sent to Home Assistant for domain siren, service turn_off
// "Turn siren off."
func NewSirenTurnOff(target Target) *SirenTurnOff {
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
		ServiceData: nil,
	}
	return s
}

type SirenTurnOff struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *SirenTurnOff) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SirenTurnOff) Targets() []string {
	return s.Target.EntityId
}
func (s *SirenTurnOff) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *SirenTurnOff) SetID(id *int) {
	s.Id = id
}

// NewSirenTurnOn creates the object that can be sent to Home Assistant for domain siren, service turn_on
// "Turn siren on."
func NewSirenTurnOn(target Target) *SirenTurnOn {
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
		ServiceData: SirenTurnOnParams{},
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

func (s *SirenTurnOn) Duration(duration string) *SirenTurnOn {
	s.ServiceData.Duration = &duration
	return s
}
func (s *SirenTurnOn) Tone(tone string) *SirenTurnOn {
	s.ServiceData.Tone = &tone
	return s
}
func (s *SirenTurnOn) VolumeLevel(volumeLevel float64) *SirenTurnOn {
	s.ServiceData.VolumeLevel = &volumeLevel
	return s
}
func (s *SirenTurnOn) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SirenTurnOn) Targets() []string {
	return s.Target.EntityId
}
func (s *SirenTurnOn) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *SirenTurnOn) SetID(id *int) {
	s.Id = id
}
