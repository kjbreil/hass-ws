package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewClimateSetAuxHeat creates the object that can be sent to Home Assistant for domain climate, service set_aux_heat
// "Turn auxiliary heater on/off for climate device."
func NewClimateSetAuxHeat(entities []string) *ClimateSetAuxHeat {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_aux_heat"
	c := &ClimateSetAuxHeat{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type ClimateSetAuxHeat struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *ClimateSetAuxHeat) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ClimateSetAuxHeat) SetID(id *int) {
	c.Id = id
}

// NewClimateSetFanMode creates the object that can be sent to Home Assistant for domain climate, service set_fan_mode
// "Set fan operation for climate device."
func NewClimateSetFanMode(entities []string, fanMode *string) *ClimateSetFanMode {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_fan_mode"
	c := &ClimateSetFanMode{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			FanMode *string `json:"fan_mode,omitempty"`
		}{FanMode: fanMode},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type ClimateSetFanMode struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		FanMode *string `json:"fan_mode,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *ClimateSetFanMode) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ClimateSetFanMode) SetID(id *int) {
	c.Id = id
}

// NewClimateSetHumidity creates the object that can be sent to Home Assistant for domain climate, service set_humidity
// "Set target humidity of climate device."
func NewClimateSetHumidity(entities []string, humidity *int) *ClimateSetHumidity {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_humidity"
	c := &ClimateSetHumidity{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Humidity *int `json:"humidity,omitempty"`
		}{Humidity: humidity},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type ClimateSetHumidity struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Humidity *int `json:"humidity,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *ClimateSetHumidity) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ClimateSetHumidity) SetID(id *int) {
	c.Id = id
}

// NewClimateSetHvacMode creates the object that can be sent to Home Assistant for domain climate, service set_hvac_mode
// "Set HVAC operation mode for climate device."
func NewClimateSetHvacMode(entities []string, hvacMode *HvacMode) *ClimateSetHvacMode {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_hvac_mode"
	c := &ClimateSetHvacMode{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			HvacMode *HvacMode `json:"hvac_mode,omitempty"`
		}{HvacMode: hvacMode},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type ClimateSetHvacMode struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		HvacMode *HvacMode `json:"hvac_mode,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *ClimateSetHvacMode) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ClimateSetHvacMode) SetID(id *int) {
	c.Id = id
}

// NewClimateSetPresetMode creates the object that can be sent to Home Assistant for domain climate, service set_preset_mode
// "Set preset mode for climate device."
func NewClimateSetPresetMode(entities []string, presetMode *string) *ClimateSetPresetMode {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_preset_mode"
	c := &ClimateSetPresetMode{
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
	return c
}

type ClimateSetPresetMode struct {
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

func (c *ClimateSetPresetMode) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ClimateSetPresetMode) SetID(id *int) {
	c.Id = id
}

// NewClimateSetSwingMode creates the object that can be sent to Home Assistant for domain climate, service set_swing_mode
// "Set swing operation for climate device."
func NewClimateSetSwingMode(entities []string, swingMode *string) *ClimateSetSwingMode {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_swing_mode"
	c := &ClimateSetSwingMode{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			SwingMode *string `json:"swing_mode,omitempty"`
		}{SwingMode: swingMode},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type ClimateSetSwingMode struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		SwingMode *string `json:"swing_mode,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *ClimateSetSwingMode) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ClimateSetSwingMode) SetID(id *int) {
	c.Id = id
}

// NewClimateSetTemperature creates the object that can be sent to Home Assistant for domain climate, service set_temperature
// "Set target temperature of climate device."
func NewClimateSetTemperature(entities []string, hvacMode *HvacMode, targetTempHigh *int, targetTempLow *int, temperature *int) *ClimateSetTemperature {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_temperature"
	c := &ClimateSetTemperature{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			HvacMode       *HvacMode `json:"hvac_mode,omitempty"`
			TargetTempHigh *int      `json:"target_temp_high,omitempty"`
			TargetTempLow  *int      `json:"target_temp_low,omitempty"`
			Temperature    *int      `json:"temperature,omitempty"`
		}{
			HvacMode:       hvacMode,
			TargetTempHigh: targetTempHigh,
			TargetTempLow:  targetTempLow,
			Temperature:    temperature,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type ClimateSetTemperature struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		HvacMode       *HvacMode `json:"hvac_mode,omitempty"`
		TargetTempHigh *int      `json:"target_temp_high,omitempty"`
		TargetTempLow  *int      `json:"target_temp_low,omitempty"`
		Temperature    *int      `json:"temperature,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *ClimateSetTemperature) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ClimateSetTemperature) SetID(id *int) {
	c.Id = id
}

// NewClimateTurnOff creates the object that can be sent to Home Assistant for domain climate, service turn_off
// "Turn climate device off."
func NewClimateTurnOff(entities []string) *ClimateTurnOff {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "turn_off"
	c := &ClimateTurnOff{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type ClimateTurnOff struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *ClimateTurnOff) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ClimateTurnOff) SetID(id *int) {
	c.Id = id
}

// NewClimateTurnOn creates the object that can be sent to Home Assistant for domain climate, service turn_on
// "Turn climate device on."
func NewClimateTurnOn(entities []string) *ClimateTurnOn {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "turn_on"
	c := &ClimateTurnOn{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type ClimateTurnOn struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *ClimateTurnOn) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ClimateTurnOn) SetID(id *int) {
	c.Id = id
}
