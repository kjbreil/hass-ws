package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewClimateSetAuxHeat creates the object that can be sent to Home Assistant for domain climate, service set_aux_heat
// "Turn auxiliary heater on/off for climate device."
func NewClimateSetAuxHeat(target Target) *ClimateSetAuxHeat {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_aux_heat"
	c := &ClimateSetAuxHeat{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type ClimateSetAuxHeat struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *ClimateSetAuxHeat) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *ClimateSetAuxHeat) Targets() []string {
	return c.Target.EntityId
}
func (c *ClimateSetAuxHeat) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewClimateSetFanMode creates the object that can be sent to Home Assistant for domain climate, service set_fan_mode
// "Set fan operation for climate device."
func NewClimateSetFanMode(target Target) *ClimateSetFanMode {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_fan_mode"
	c := &ClimateSetFanMode{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: ClimateSetFanModeParams{},
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

func (c *ClimateSetFanMode) FanMode(fanMode string) *ClimateSetFanMode {
	c.ServiceData.FanMode = &fanMode
	return c
}
func (c *ClimateSetFanMode) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *ClimateSetFanMode) Targets() []string {
	return c.Target.EntityId
}
func (c *ClimateSetFanMode) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewClimateSetHumidity creates the object that can be sent to Home Assistant for domain climate, service set_humidity
// "Set target humidity of climate device."
func NewClimateSetHumidity(target Target) *ClimateSetHumidity {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_humidity"
	c := &ClimateSetHumidity{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: ClimateSetHumidityParams{},
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

func (c *ClimateSetHumidity) Humidity(humidity float64) *ClimateSetHumidity {
	c.ServiceData.Humidity = &humidity
	return c
}
func (c *ClimateSetHumidity) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *ClimateSetHumidity) Targets() []string {
	return c.Target.EntityId
}
func (c *ClimateSetHumidity) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewClimateSetHvacMode creates the object that can be sent to Home Assistant for domain climate, service set_hvac_mode
// "Set HVAC operation mode for climate device."
func NewClimateSetHvacMode(target Target) *ClimateSetHvacMode {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_hvac_mode"
	c := &ClimateSetHvacMode{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: ClimateSetHvacModeParams{},
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

func (c *ClimateSetHvacMode) HvacMode(hvacMode HvacMode) *ClimateSetHvacMode {
	c.ServiceData.HvacMode = &hvacMode
	return c
}
func (c *ClimateSetHvacMode) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *ClimateSetHvacMode) Targets() []string {
	return c.Target.EntityId
}
func (c *ClimateSetHvacMode) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewClimateSetPresetMode creates the object that can be sent to Home Assistant for domain climate, service set_preset_mode
// "Set preset mode for climate device."
func NewClimateSetPresetMode(target Target) *ClimateSetPresetMode {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_preset_mode"
	c := &ClimateSetPresetMode{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: ClimateSetPresetModeParams{},
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

func (c *ClimateSetPresetMode) PresetMode(presetMode string) *ClimateSetPresetMode {
	c.ServiceData.PresetMode = &presetMode
	return c
}
func (c *ClimateSetPresetMode) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *ClimateSetPresetMode) Targets() []string {
	return c.Target.EntityId
}
func (c *ClimateSetPresetMode) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewClimateSetSwingMode creates the object that can be sent to Home Assistant for domain climate, service set_swing_mode
// "Set swing operation for climate device."
func NewClimateSetSwingMode(target Target) *ClimateSetSwingMode {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_swing_mode"
	c := &ClimateSetSwingMode{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: ClimateSetSwingModeParams{},
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

func (c *ClimateSetSwingMode) SwingMode(swingMode string) *ClimateSetSwingMode {
	c.ServiceData.SwingMode = &swingMode
	return c
}
func (c *ClimateSetSwingMode) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *ClimateSetSwingMode) Targets() []string {
	return c.Target.EntityId
}
func (c *ClimateSetSwingMode) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewClimateSetTemperature creates the object that can be sent to Home Assistant for domain climate, service set_temperature
// "Set target temperature of climate device."
func NewClimateSetTemperature(target Target) *ClimateSetTemperature {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "set_temperature"
	c := &ClimateSetTemperature{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: ClimateSetTemperatureParams{},
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

func (c *ClimateSetTemperature) HvacMode(hvacMode HvacMode) *ClimateSetTemperature {
	c.ServiceData.HvacMode = &hvacMode
	return c
}
func (c *ClimateSetTemperature) TargetTempHigh(targetTempHigh float64) *ClimateSetTemperature {
	c.ServiceData.TargetTempHigh = &targetTempHigh
	return c
}
func (c *ClimateSetTemperature) TargetTempLow(targetTempLow float64) *ClimateSetTemperature {
	c.ServiceData.TargetTempLow = &targetTempLow
	return c
}
func (c *ClimateSetTemperature) Temperature(temperature float64) *ClimateSetTemperature {
	c.ServiceData.Temperature = &temperature
	return c
}
func (c *ClimateSetTemperature) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *ClimateSetTemperature) Targets() []string {
	return c.Target.EntityId
}
func (c *ClimateSetTemperature) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewClimateTurnOff creates the object that can be sent to Home Assistant for domain climate, service turn_off
// "Turn climate device off."
func NewClimateTurnOff(target Target) *ClimateTurnOff {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "turn_off"
	c := &ClimateTurnOff{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type ClimateTurnOff struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *ClimateTurnOff) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *ClimateTurnOff) Targets() []string {
	return c.Target.EntityId
}
func (c *ClimateTurnOff) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewClimateTurnOn creates the object that can be sent to Home Assistant for domain climate, service turn_on
// "Turn climate device on."
func NewClimateTurnOn(target Target) *ClimateTurnOn {
	serviceDomain := "climate"
	serviceType := "call_service"
	serviceService := "turn_on"
	c := &ClimateTurnOn{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type ClimateTurnOn struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *ClimateTurnOn) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *ClimateTurnOn) Targets() []string {
	return c.Target.EntityId
}
func (c *ClimateTurnOn) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}
