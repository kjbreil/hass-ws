package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSirenToggle creates the object that can be sent to Home Assistant for domain siren, service toggle
// "Toggles a siren."
func NewSirenToggle(entities []string) *SirenToggle {
	serviceDomain := "siren"
	serviceType := "call_service"
	serviceService := "toggle"
	s := &SirenToggle{
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

type SirenToggle struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SirenToggle) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SirenToggle) SetID(id *int) {
	s.Id = id
}

// NewSirenTurnOff creates the object that can be sent to Home Assistant for domain siren, service turn_off
// "Turn siren off."
func NewSirenTurnOff(entities []string) *SirenTurnOff {
	serviceDomain := "siren"
	serviceType := "call_service"
	serviceService := "turn_off"
	s := &SirenTurnOff{
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

type SirenTurnOff struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SirenTurnOff) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SirenTurnOff) SetID(id *int) {
	s.Id = id
}

// NewSirenTurnOn creates the object that can be sent to Home Assistant for domain siren, service turn_on
// "Turn siren on."
func NewSirenTurnOn(entities []string, duration *string, tone *string, volumeLevel *int) *SirenTurnOn {
	serviceDomain := "siren"
	serviceType := "call_service"
	serviceService := "turn_on"
	s := &SirenTurnOn{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Duration    *string `json:"duration,omitempty"`
			Tone        *string `json:"tone,omitempty"`
			VolumeLevel *int    `json:"volume_level,omitempty"`
		}{
			Duration:    duration,
			Tone:        tone,
			VolumeLevel: volumeLevel,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SirenTurnOn struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Duration    *string `json:"duration,omitempty"`
		Tone        *string `json:"tone,omitempty"`
		VolumeLevel *int    `json:"volume_level,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SirenTurnOn) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SirenTurnOn) SetID(id *int) {
	s.Id = id
}
