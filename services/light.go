package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewLightToggle creates the object that can be sent to Home Assistant for domain light, service toggle
// "Toggles one or more lights, from on to off, or, off to on, based on their current state.\n"
func NewLightToggle(target Target, lightToggleParams LightToggleParams) *LightToggle {
	serviceDomain := "light"
	serviceType := "call_service"
	serviceService := "toggle"
	l := &LightToggle{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: lightToggleParams,
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

func (l *LightToggle) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LightToggle) SetID(id *int) {
	l.Id = id
}

// NewLightTurnOff creates the object that can be sent to Home Assistant for domain light, service turn_off
// "Turns off one or more lights."
func NewLightTurnOff(target Target, lightTurnOffParams LightTurnOffParams) *LightTurnOff {
	serviceDomain := "light"
	serviceType := "call_service"
	serviceService := "turn_off"
	l := &LightTurnOff{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: lightTurnOffParams,
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

func (l *LightTurnOff) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LightTurnOff) SetID(id *int) {
	l.Id = id
}

// NewLightTurnOn creates the object that can be sent to Home Assistant for domain light, service turn_on
// "Turn on one or more lights and adjust properties of the light, even when they are turned on already.\n"
func NewLightTurnOn(target Target, lightTurnOnParams LightTurnOnParams) *LightTurnOn {
	serviceDomain := "light"
	serviceType := "call_service"
	serviceService := "turn_on"
	l := &LightTurnOn{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: lightTurnOnParams,
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

func (l *LightTurnOn) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LightTurnOn) SetID(id *int) {
	l.Id = id
}
