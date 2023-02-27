package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Climate struct {
	CurrentHumidity       *int      `json:"current_humidity,omitempty"`
	CurrentTemperature    *float64  `json:"current_temperature,omitempty"`
	FanMode               *string   `json:"fan_mode,omitempty"`
	FanModes              *[]string `json:"fan_modes,omitempty"`
	HvacAction            *string   `json:"hvac_action,omitempty"`
	HvacMode              *[]string `json:"hvac_mode,omitempty"`
	HvacModes             *[]string `json:"hvac_modes,omitempty"`
	IsAuxHeat             *bool     `json:"is_aux_heat,omitempty"`
	MaxHumidity           *int      `json:"max_humidity,omitempty"`
	MaxTemp               *float64  `json:"max_temp,omitempty"`
	MinHumidity           *int      `json:"min_humidity,omitempty"`
	MinTemp               *float64  `json:"min_temp,omitempty"`
	Precision             *float64  `json:"precision,omitempty"`
	PresetMode            *string   `json:"preset_mode,omitempty"`
	PresetModes           *[]string `json:"preset_modes,omitempty"`
	SupportedFeatures     *int      `json:"supported_features,omitempty"`
	SwingMode             *string   `json:"swing_mode,omitempty"`
	SwingModes            *[]string `json:"swing_modes,omitempty"`
	TargetHumidity        *float64  `json:"target_humidity,omitempty"`
	TargetTemperature     *float64  `json:"target_temperature,omitempty"`
	TargetTemperatureHigh *float64  `json:"target_temperature_high,omitempty"`
	TargetTemperatureLow  *float64  `json:"target_temperature_low,omitempty"`
	TargetTemperatureStep *float64  `json:"target_temperature_step,omitempty"`
	TemperatureUnit       *string   `json:"temperature_unit,omitempty"`
}
