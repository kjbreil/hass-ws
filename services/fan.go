package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewFanDecreaseSpeed creates the object that can be sent to Home Assistant for domain fan, service decrease_speed
// "Decrease the speed of the fan by one speed or a percentage_step."
func NewFanDecreaseSpeed(entities []string, percentageStep *int) *FanDecreaseSpeed {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "decrease_speed"
	f := &FanDecreaseSpeed{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			PercentageStep *int `json:"percentage_step,omitempty"`
		}{PercentageStep: percentageStep},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FanDecreaseSpeed struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		PercentageStep *int `json:"percentage_step,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewFanIncreaseSpeed(entities []string, percentageStep *int) *FanIncreaseSpeed {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "increase_speed"
	f := &FanIncreaseSpeed{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			PercentageStep *int `json:"percentage_step,omitempty"`
		}{PercentageStep: percentageStep},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FanIncreaseSpeed struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		PercentageStep *int `json:"percentage_step,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewFanOscillate(entities []string) *FanOscillate {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "oscillate"
	f := &FanOscillate{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FanOscillate struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewFanSetDirection(entities []string, direction *Direction) *FanSetDirection {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "set_direction"
	f := &FanSetDirection{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Direction *Direction `json:"direction,omitempty"`
		}{Direction: direction},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FanSetDirection struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Direction *Direction `json:"direction,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewFanSetPercentage(entities []string, percentage *int) *FanSetPercentage {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "set_percentage"
	f := &FanSetPercentage{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Percentage *int `json:"percentage,omitempty"`
		}{Percentage: percentage},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FanSetPercentage struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Percentage *int `json:"percentage,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewFanSetPresetMode(entities []string, presetMode *string) *FanSetPresetMode {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "set_preset_mode"
	f := &FanSetPresetMode{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			PresetMode *string `json:"preset_mode,omitempty"`
		}{PresetMode: presetMode},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FanSetPresetMode struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		PresetMode *string `json:"preset_mode,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewFanToggle(entities []string) *FanToggle {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "toggle"
	f := &FanToggle{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FanToggle struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewFanTurnOff(entities []string) *FanTurnOff {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "turn_off"
	f := &FanTurnOff{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FanTurnOff struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewFanTurnOn(entities []string, percentage *int, presetMode *string, speed *string) *FanTurnOn {
	serviceDomain := "fan"
	serviceType := "call_service"
	serviceService := "turn_on"
	f := &FanTurnOn{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Percentage *int    `json:"percentage,omitempty"`
			PresetMode *string `json:"preset_mode,omitempty"`
			Speed      *string `json:"speed,omitempty"`
		}{
			Percentage: percentage,
			PresetMode: presetMode,
			Speed:      speed,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FanTurnOn struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Percentage *int    `json:"percentage,omitempty"`
		PresetMode *string `json:"preset_mode,omitempty"`
		Speed      *string `json:"speed,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (f *FanTurnOn) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FanTurnOn) SetID(id *int) {
	f.Id = id
}
