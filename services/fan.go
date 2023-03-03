package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewFanDecreaseSpeed creates the object that can be sent to Home Assistant for domain fan, service decrease_speed
// "Decrease the speed of the fan by one speed or a percentage_step."
func NewFanDecreaseSpeed(target Target, fanDecreaseSpeedParams *FanDecreaseSpeedParams) *FanDecreaseSpeed {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "decrease_speed"
	f := &FanDecreaseSpeed{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *fanDecreaseSpeedParams,
	}
	return f
}

type FanDecreaseSpeed struct {
	ServiceBase
	ServiceData FanDecreaseSpeedParams `json:"service_data,omitempty"`
}
type FanDecreaseSpeedParams struct {
	PercentageStep *float64 `json:"percentage_step,omitempty"`
}

func (f *FanDecreaseSpeed) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FanDecreaseSpeed) SetID(id *int) {
	f.Id = id
}

// NewFanIncreaseSpeed creates the object that can be sent to Home Assistant for domain fan, service increase_speed
// "Increase the speed of the fan by one speed or a percentage_step."
func NewFanIncreaseSpeed(target Target, fanIncreaseSpeedParams *FanIncreaseSpeedParams) *FanIncreaseSpeed {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "increase_speed"
	f := &FanIncreaseSpeed{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *fanIncreaseSpeedParams,
	}
	return f
}

type FanIncreaseSpeed struct {
	ServiceBase
	ServiceData FanIncreaseSpeedParams `json:"service_data,omitempty"`
}
type FanIncreaseSpeedParams struct {
	PercentageStep *float64 `json:"percentage_step,omitempty"`
}

func (f *FanIncreaseSpeed) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FanIncreaseSpeed) SetID(id *int) {
	f.Id = id
}

// NewFanOscillate creates the object that can be sent to Home Assistant for domain fan, service oscillate
// "Oscillate the fan."
func NewFanOscillate(target Target) *FanOscillate {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "oscillate"
	f := &FanOscillate{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return f
}

type FanOscillate struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (f *FanOscillate) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FanOscillate) SetID(id *int) {
	f.Id = id
}

// NewFanSetDirection creates the object that can be sent to Home Assistant for domain fan, service set_direction
// "Set the fan rotation."
func NewFanSetDirection(target Target, fanSetDirectionParams *FanSetDirectionParams) *FanSetDirection {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "set_direction"
	f := &FanSetDirection{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *fanSetDirectionParams,
	}
	return f
}

type FanSetDirection struct {
	ServiceBase
	ServiceData FanSetDirectionParams `json:"service_data,omitempty"`
}
type FanSetDirectionParams struct {
	Direction *Direction `json:"direction,omitempty"`
}

func (f *FanSetDirection) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FanSetDirection) SetID(id *int) {
	f.Id = id
}

// NewFanSetPercentage creates the object that can be sent to Home Assistant for domain fan, service set_percentage
// "Set fan speed percentage."
func NewFanSetPercentage(target Target, fanSetPercentageParams *FanSetPercentageParams) *FanSetPercentage {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "set_percentage"
	f := &FanSetPercentage{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *fanSetPercentageParams,
	}
	return f
}

type FanSetPercentage struct {
	ServiceBase
	ServiceData FanSetPercentageParams `json:"service_data,omitempty"`
}
type FanSetPercentageParams struct {
	Percentage *float64 `json:"percentage,omitempty"`
}

func (f *FanSetPercentage) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FanSetPercentage) SetID(id *int) {
	f.Id = id
}

// NewFanSetPresetMode creates the object that can be sent to Home Assistant for domain fan, service set_preset_mode
// "Set preset mode for a fan device."
func NewFanSetPresetMode(target Target, fanSetPresetModeParams *FanSetPresetModeParams) *FanSetPresetMode {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "set_preset_mode"
	f := &FanSetPresetMode{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *fanSetPresetModeParams,
	}
	return f
}

type FanSetPresetMode struct {
	ServiceBase
	ServiceData FanSetPresetModeParams `json:"service_data,omitempty"`
}
type FanSetPresetModeParams struct {
	PresetMode *string `json:"preset_mode,omitempty"`
}

func (f *FanSetPresetMode) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FanSetPresetMode) SetID(id *int) {
	f.Id = id
}

// NewFanToggle creates the object that can be sent to Home Assistant for domain fan, service toggle
// "Toggle the fan on/off."
func NewFanToggle(target Target) *FanToggle {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "toggle"
	f := &FanToggle{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return f
}

type FanToggle struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (f *FanToggle) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FanToggle) SetID(id *int) {
	f.Id = id
}

// NewFanTurnOff creates the object that can be sent to Home Assistant for domain fan, service turn_off
// "Turn fan off."
func NewFanTurnOff(target Target) *FanTurnOff {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "turn_off"
	f := &FanTurnOff{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return f
}

type FanTurnOff struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (f *FanTurnOff) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FanTurnOff) SetID(id *int) {
	f.Id = id
}

// NewFanTurnOn creates the object that can be sent to Home Assistant for domain fan, service turn_on
// "Turn fan on."
func NewFanTurnOn(target Target, fanTurnOnParams *FanTurnOnParams) *FanTurnOn {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "turn_on"
	f := &FanTurnOn{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *fanTurnOnParams,
	}
	return f
}

type FanTurnOn struct {
	ServiceBase
	ServiceData FanTurnOnParams `json:"service_data,omitempty"`
}
type FanTurnOnParams struct {
	Percentage *float64 `json:"percentage,omitempty"`
	PresetMode *string  `json:"preset_mode,omitempty"`
	Speed      *string  `json:"speed,omitempty"`
}

func (f *FanTurnOn) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FanTurnOn) SetID(id *int) {
	f.Id = id
}
