package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewLightToggle creates the object that can be sent to Home Assistant for domain light, service toggle
// "Toggles one or more lights, from on to off, or off to on, based on their current state."
func NewLightToggle(target Target) *LightToggle {
	serviceDomain := "light"
	serviceType := "call_service"
	serviceService := "toggle"
	l := &LightToggle{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: LightToggleParams{},
	}
	return l
}

type LightToggle struct {
	ServiceBase
	ServiceData LightToggleParams `json:"service_data,omitempty"`
}
type LightToggleParams struct {
	BrightnessPct   *float64 `json:"brightness_pct,omitempty"`
	ColorTempKelvin *float64 `json:"color_temp_kelvin,omitempty"`
	Effect          *string  `json:"effect,omitempty"`
	Transition      *float64 `json:"transition,omitempty"`
}

func (l *LightToggle) BrightnessPct(brightnessPct float64) *LightToggle {
	l.ServiceData.BrightnessPct = &brightnessPct
	return l
}
func (l *LightToggle) ColorTempKelvin(colorTempKelvin float64) *LightToggle {
	l.ServiceData.ColorTempKelvin = &colorTempKelvin
	return l
}
func (l *LightToggle) Effect(effect string) *LightToggle {
	l.ServiceData.Effect = &effect
	return l
}
func (l *LightToggle) Transition(transition float64) *LightToggle {
	l.ServiceData.Transition = &transition
	return l
}
func (l *LightToggle) JSON() string {
	data, _ := gojson.Marshal(l)
	return string(data)
}
func (l *LightToggle) Targets() []string {
	return l.Target.EntityId
}
func (l *LightToggle) Name() string {
	return fmt.Sprintf("%s.%s", *l.Domain, *l.Service)
}

// NewLightTurnOff creates the object that can be sent to Home Assistant for domain light, service turn_off
// "Turns off one or more lights."
func NewLightTurnOff(target Target) *LightTurnOff {
	serviceDomain := "light"
	serviceType := "call_service"
	serviceService := "turn_off"
	l := &LightTurnOff{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: LightTurnOffParams{},
	}
	return l
}

type LightTurnOff struct {
	ServiceBase
	ServiceData LightTurnOffParams `json:"service_data,omitempty"`
}
type LightTurnOffParams struct {
	Transition *float64 `json:"transition,omitempty"`
}

func (l *LightTurnOff) Transition(transition float64) *LightTurnOff {
	l.ServiceData.Transition = &transition
	return l
}
func (l *LightTurnOff) JSON() string {
	data, _ := gojson.Marshal(l)
	return string(data)
}
func (l *LightTurnOff) Targets() []string {
	return l.Target.EntityId
}
func (l *LightTurnOff) Name() string {
	return fmt.Sprintf("%s.%s", *l.Domain, *l.Service)
}

// NewLightTurnOn creates the object that can be sent to Home Assistant for domain light, service turn_on
// "Turns on one or more lights and adjusts their properties, even when they are turned on already."
func NewLightTurnOn(target Target) *LightTurnOn {
	serviceDomain := "light"
	serviceType := "call_service"
	serviceService := "turn_on"
	l := &LightTurnOn{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: LightTurnOnParams{},
	}
	return l
}

type LightTurnOn struct {
	ServiceBase
	ServiceData LightTurnOnParams `json:"service_data,omitempty"`
}
type LightTurnOnParams struct {
	BrightnessPct     *float64 `json:"brightness_pct,omitempty"`
	BrightnessStepPct *float64 `json:"brightness_step_pct,omitempty"`
	ColorTempKelvin   *float64 `json:"color_temp_kelvin,omitempty"`
	Effect            *string  `json:"effect,omitempty"`
	Transition        *float64 `json:"transition,omitempty"`
}

func (l *LightTurnOn) BrightnessPct(brightnessPct float64) *LightTurnOn {
	l.ServiceData.BrightnessPct = &brightnessPct
	return l
}
func (l *LightTurnOn) BrightnessStepPct(brightnessStepPct float64) *LightTurnOn {
	l.ServiceData.BrightnessStepPct = &brightnessStepPct
	return l
}
func (l *LightTurnOn) ColorTempKelvin(colorTempKelvin float64) *LightTurnOn {
	l.ServiceData.ColorTempKelvin = &colorTempKelvin
	return l
}
func (l *LightTurnOn) Effect(effect string) *LightTurnOn {
	l.ServiceData.Effect = &effect
	return l
}
func (l *LightTurnOn) Transition(transition float64) *LightTurnOn {
	l.ServiceData.Transition = &transition
	return l
}
func (l *LightTurnOn) JSON() string {
	data, _ := gojson.Marshal(l)
	return string(data)
}
func (l *LightTurnOn) Targets() []string {
	return l.Target.EntityId
}
func (l *LightTurnOn) Name() string {
	return fmt.Sprintf("%s.%s", *l.Domain, *l.Service)
}
