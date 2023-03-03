package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewClimateSetAuxHeat creates the object that can be sent to Home Assistant for domain climate, service set_aux_heat
// "Turn auxiliary heater on/off for climate device."
func NewClimateSetAuxHeat(target Target, climateSetAuxHeatParams ClimateSetAuxHeatParams) *ClimateSetAuxHeat {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_aux_heat"
	c := &ClimateSetAuxHeat{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: climateSetAuxHeatParams,
	}
	return c
}

type ClimateSetAuxHeat struct {
	ServiceBase
	ServiceData ClimateSetAuxHeatParams `json:"service_data,omitempty"`
}
type ClimateSetAuxHeatParams struct{}

func (c *ClimateSetAuxHeat) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ClimateSetAuxHeat) SetID(id *int) {
	c.Id = id
}

// NewClimateSetFanMode creates the object that can be sent to Home Assistant for domain climate, service set_fan_mode
// "Set fan operation for climate device."
func NewClimateSetFanMode(target Target, climateSetFanModeParams ClimateSetFanModeParams) *ClimateSetFanMode {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_fan_mode"
	c := &ClimateSetFanMode{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: climateSetFanModeParams,
	}
	return c
}

type ClimateSetFanMode struct {
	ServiceBase
	ServiceData ClimateSetFanModeParams `json:"service_data,omitempty"`
}
type ClimateSetFanModeParams struct {
	FanMode *string `json:"fan_mode,omitempty"`
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
func NewClimateSetHumidity(target Target, climateSetHumidityParams ClimateSetHumidityParams) *ClimateSetHumidity {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_humidity"
	c := &ClimateSetHumidity{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: climateSetHumidityParams,
	}
	return c
}

type ClimateSetHumidity struct {
	ServiceBase
	ServiceData ClimateSetHumidityParams `json:"service_data,omitempty"`
}
type ClimateSetHumidityParams struct {
	Humidity *float64 `json:"humidity,omitempty"`
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
func NewClimateSetHvacMode(target Target, climateSetHvacModeParams ClimateSetHvacModeParams) *ClimateSetHvacMode {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_hvac_mode"
	c := &ClimateSetHvacMode{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: climateSetHvacModeParams,
	}
	return c
}

type ClimateSetHvacMode struct {
	ServiceBase
	ServiceData ClimateSetHvacModeParams `json:"service_data,omitempty"`
}
type ClimateSetHvacModeParams struct {
	HvacMode *HvacMode `json:"hvac_mode,omitempty"`
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
func NewClimateSetPresetMode(target Target, climateSetPresetModeParams ClimateSetPresetModeParams) *ClimateSetPresetMode {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_preset_mode"
	c := &ClimateSetPresetMode{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: climateSetPresetModeParams,
	}
	return c
}

type ClimateSetPresetMode struct {
	ServiceBase
	ServiceData ClimateSetPresetModeParams `json:"service_data,omitempty"`
}
type ClimateSetPresetModeParams struct {
	PresetMode *string `json:"preset_mode,omitempty"`
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
func NewClimateSetSwingMode(target Target, climateSetSwingModeParams ClimateSetSwingModeParams) *ClimateSetSwingMode {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_swing_mode"
	c := &ClimateSetSwingMode{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: climateSetSwingModeParams,
	}
	return c
}

type ClimateSetSwingMode struct {
	ServiceBase
	ServiceData ClimateSetSwingModeParams `json:"service_data,omitempty"`
}
type ClimateSetSwingModeParams struct {
	SwingMode *string `json:"swing_mode,omitempty"`
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
func NewClimateSetTemperature(target Target, climateSetTemperatureParams ClimateSetTemperatureParams) *ClimateSetTemperature {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_temperature"
	c := &ClimateSetTemperature{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: climateSetTemperatureParams,
	}
	return c
}

type ClimateSetTemperature struct {
	ServiceBase
	ServiceData ClimateSetTemperatureParams `json:"service_data,omitempty"`
}
type ClimateSetTemperatureParams struct {
	HvacMode       *HvacMode `json:"hvac_mode,omitempty"`
	TargetTempHigh *float64  `json:"target_temp_high,omitempty"`
	TargetTempLow  *float64  `json:"target_temp_low,omitempty"`
	Temperature    *float64  `json:"temperature,omitempty"`
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
func NewClimateTurnOff(target Target, climateTurnOffParams ClimateTurnOffParams) *ClimateTurnOff {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "turn_off"
	c := &ClimateTurnOff{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: climateTurnOffParams,
	}
	return c
}

type ClimateTurnOff struct {
	ServiceBase
	ServiceData ClimateTurnOffParams `json:"service_data,omitempty"`
}
type ClimateTurnOffParams struct{}

func (c *ClimateTurnOff) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ClimateTurnOff) SetID(id *int) {
	c.Id = id
}

// NewClimateTurnOn creates the object that can be sent to Home Assistant for domain climate, service turn_on
// "Turn climate device on."
func NewClimateTurnOn(target Target, climateTurnOnParams ClimateTurnOnParams) *ClimateTurnOn {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "turn_on"
	c := &ClimateTurnOn{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: climateTurnOnParams,
	}
	return c
}

type ClimateTurnOn struct {
	ServiceBase
	ServiceData ClimateTurnOnParams `json:"service_data,omitempty"`
}
type ClimateTurnOnParams struct{}

func (c *ClimateTurnOn) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *ClimateTurnOn) SetID(id *int) {
	c.Id = id
}
