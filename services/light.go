package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewLightToggle creates the object that can be sent to Home Assistant for domain light, service toggle
// "Toggles one or more lights, from on to off, or, off to on, based on their current state.\n"
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
	Brightness    *float64   `json:"brightness,omitempty"`
	BrightnessPct *float64   `json:"brightness_pct,omitempty"`
	ColorName     *ColorName `json:"color_name,omitempty"`
	ColorTemp     *float64   `json:"color_temp,omitempty"`
	Effect        *string    `json:"effect,omitempty"`
	Flash         *Flash     `json:"flash,omitempty"`
	Kelvin        *float64   `json:"kelvin,omitempty"`
	Profile       *string    `json:"profile,omitempty"`
	Transition    *float64   `json:"transition,omitempty"`
}

func (l *LightToggle) Brightness(brightness float64) *LightToggle {
	l.ServiceData.Brightness = &brightness
	return l
}
func (l *LightToggle) BrightnessPct(brightnessPct float64) *LightToggle {
	l.ServiceData.BrightnessPct = &brightnessPct
	return l
}
func (l *LightToggle) ColorName(colorName ColorName) *LightToggle {
	l.ServiceData.ColorName = &colorName
	return l
}
func (l *LightToggle) ColorTemp(colorTemp float64) *LightToggle {
	l.ServiceData.ColorTemp = &colorTemp
	return l
}
func (l *LightToggle) Effect(effect string) *LightToggle {
	l.ServiceData.Effect = &effect
	return l
}
func (l *LightToggle) Flash(flash Flash) *LightToggle {
	l.ServiceData.Flash = &flash
	return l
}
func (l *LightToggle) Kelvin(kelvin float64) *LightToggle {
	l.ServiceData.Kelvin = &kelvin
	return l
}
func (l *LightToggle) Profile(profile string) *LightToggle {
	l.ServiceData.Profile = &profile
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
	Flash      *Flash   `json:"flash,omitempty"`
	Transition *float64 `json:"transition,omitempty"`
}

func (l *LightTurnOff) Flash(flash Flash) *LightTurnOff {
	l.ServiceData.Flash = &flash
	return l
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
// "Turn on one or more lights and adjust properties of the light, even when they are turned on already.\n"
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
	Brightness        *float64   `json:"brightness,omitempty"`
	BrightnessPct     *float64   `json:"brightness_pct,omitempty"`
	BrightnessStep    *float64   `json:"brightness_step,omitempty"`
	BrightnessStepPct *float64   `json:"brightness_step_pct,omitempty"`
	ColorName         *ColorName `json:"color_name,omitempty"`
	ColorTemp         *float64   `json:"color_temp,omitempty"`
	Effect            *string    `json:"effect,omitempty"`
	Flash             *Flash     `json:"flash,omitempty"`
	Kelvin            *float64   `json:"kelvin,omitempty"`
	Profile           *string    `json:"profile,omitempty"`
	Transition        *float64   `json:"transition,omitempty"`
	White             *float64   `json:"white,omitempty"`
}

func (l *LightTurnOn) Brightness(brightness float64) *LightTurnOn {
	l.ServiceData.Brightness = &brightness
	return l
}
func (l *LightTurnOn) BrightnessPct(brightnessPct float64) *LightTurnOn {
	l.ServiceData.BrightnessPct = &brightnessPct
	return l
}
func (l *LightTurnOn) BrightnessStep(brightnessStep float64) *LightTurnOn {
	l.ServiceData.BrightnessStep = &brightnessStep
	return l
}
func (l *LightTurnOn) BrightnessStepPct(brightnessStepPct float64) *LightTurnOn {
	l.ServiceData.BrightnessStepPct = &brightnessStepPct
	return l
}
func (l *LightTurnOn) ColorName(colorName ColorName) *LightTurnOn {
	l.ServiceData.ColorName = &colorName
	return l
}
func (l *LightTurnOn) ColorTemp(colorTemp float64) *LightTurnOn {
	l.ServiceData.ColorTemp = &colorTemp
	return l
}
func (l *LightTurnOn) Effect(effect string) *LightTurnOn {
	l.ServiceData.Effect = &effect
	return l
}
func (l *LightTurnOn) Flash(flash Flash) *LightTurnOn {
	l.ServiceData.Flash = &flash
	return l
}
func (l *LightTurnOn) Kelvin(kelvin float64) *LightTurnOn {
	l.ServiceData.Kelvin = &kelvin
	return l
}
func (l *LightTurnOn) Profile(profile string) *LightTurnOn {
	l.ServiceData.Profile = &profile
	return l
}
func (l *LightTurnOn) Transition(transition float64) *LightTurnOn {
	l.ServiceData.Transition = &transition
	return l
}
func (l *LightTurnOn) White(white float64) *LightTurnOn {
	l.ServiceData.White = &white
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
