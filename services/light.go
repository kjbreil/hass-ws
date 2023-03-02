package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewLightToggle creates the object that can be sent to Home Assistant for domain light, service toggle
// "Toggles one or more lights, from on to off, or, off to on, based on their current state.\n"
func NewLightToggle(entities []string, brightness *int, brightnessPct *int, colorName *ColorName, effect *string, flash *Flash, kelvin *int, profile *string, transition *int) *LightToggle {
	serviceDomain := "light"
	serviceType := "call_service"
	serviceService := "toggle"
	l := &LightToggle{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Brightness    *int       `json:"brightness,omitempty"`
			BrightnessPct *int       `json:"brightness_pct,omitempty"`
			ColorName     *ColorName `json:"color_name,omitempty"`
			Effect        *string    `json:"effect,omitempty"`
			Flash         *Flash     `json:"flash,omitempty"`
			Kelvin        *int       `json:"kelvin,omitempty"`
			Profile       *string    `json:"profile,omitempty"`
			Transition    *int       `json:"transition,omitempty"`
		}{
			Brightness:    brightness,
			BrightnessPct: brightnessPct,
			ColorName:     colorName,
			Effect:        effect,
			Flash:         flash,
			Kelvin:        kelvin,
			Profile:       profile,
			Transition:    transition,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return l
}

type LightToggle struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Brightness    *int       `json:"brightness,omitempty"`
		BrightnessPct *int       `json:"brightness_pct,omitempty"`
		ColorName     *ColorName `json:"color_name,omitempty"`
		Effect        *string    `json:"effect,omitempty"`
		Flash         *Flash     `json:"flash,omitempty"`
		Kelvin        *int       `json:"kelvin,omitempty"`
		Profile       *string    `json:"profile,omitempty"`
		Transition    *int       `json:"transition,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (l *LightToggle) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LightToggle) SetID(id *int) {
	l.Id = id
}

// NewLightTurnOff creates the object that can be sent to Home Assistant for domain light, service turn_off
// "Turns off one or more lights."
func NewLightTurnOff(entities []string, flash *Flash, transition *int) *LightTurnOff {
	serviceDomain := "light"
	serviceType := "call_service"
	serviceService := "turn_off"
	l := &LightTurnOff{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Flash      *Flash `json:"flash,omitempty"`
			Transition *int   `json:"transition,omitempty"`
		}{
			Flash:      flash,
			Transition: transition,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return l
}

type LightTurnOff struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Flash      *Flash `json:"flash,omitempty"`
		Transition *int   `json:"transition,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (l *LightTurnOff) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LightTurnOff) SetID(id *int) {
	l.Id = id
}

// NewLightTurnOn creates the object that can be sent to Home Assistant for domain light, service turn_on
// "Turn on one or more lights and adjust properties of the light, even when they are turned on already.\n"
func NewLightTurnOn(entities []string, brightness *int, brightnessPct *int, brightnessStep *int, brightnessStepPct *int, colorName *ColorName, effect *string, flash *Flash, kelvin *int, profile *string, transition *int, white *int) *LightTurnOn {
	serviceDomain := "light"
	serviceType := "call_service"
	serviceService := "turn_on"
	l := &LightTurnOn{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Brightness        *int       `json:"brightness,omitempty"`
			BrightnessPct     *int       `json:"brightness_pct,omitempty"`
			BrightnessStep    *int       `json:"brightness_step,omitempty"`
			BrightnessStepPct *int       `json:"brightness_step_pct,omitempty"`
			ColorName         *ColorName `json:"color_name,omitempty"`
			Effect            *string    `json:"effect,omitempty"`
			Flash             *Flash     `json:"flash,omitempty"`
			Kelvin            *int       `json:"kelvin,omitempty"`
			Profile           *string    `json:"profile,omitempty"`
			Transition        *int       `json:"transition,omitempty"`
			White             *int       `json:"white,omitempty"`
		}{
			Brightness:        brightness,
			BrightnessPct:     brightnessPct,
			BrightnessStep:    brightnessStep,
			BrightnessStepPct: brightnessStepPct,
			ColorName:         colorName,
			Effect:            effect,
			Flash:             flash,
			Kelvin:            kelvin,
			Profile:           profile,
			Transition:        transition,
			White:             white,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return l
}

type LightTurnOn struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Brightness        *int       `json:"brightness,omitempty"`
		BrightnessPct     *int       `json:"brightness_pct,omitempty"`
		BrightnessStep    *int       `json:"brightness_step,omitempty"`
		BrightnessStepPct *int       `json:"brightness_step_pct,omitempty"`
		ColorName         *ColorName `json:"color_name,omitempty"`
		Effect            *string    `json:"effect,omitempty"`
		Flash             *Flash     `json:"flash,omitempty"`
		Kelvin            *int       `json:"kelvin,omitempty"`
		Profile           *string    `json:"profile,omitempty"`
		Transition        *int       `json:"transition,omitempty"`
		White             *int       `json:"white,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (l *LightTurnOn) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LightTurnOn) SetID(id *int) {
	l.Id = id
}
