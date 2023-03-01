package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Climate struct {
	Additional         Additional `json:"additional,omitempty"`
	CurrentHumidity    *int       `json:"current_humidity,omitempty"`
	CurrentTemperature *float64   `json:"current_temperature,omitempty"`
	FanMode            *string    `json:"fan_mode,omitempty"`
	FanModes           *[]string  `json:"fan_modes,omitempty"`
	FriendlyName       *string    `json:"friendly_name,omitempty"`
	HvacAction         *string    `json:"hvac_action,omitempty"`
	HvacMode           *[]string  `json:"hvac_mode,omitempty"`
	HvacModes          *[]string  `json:"hvac_modes,omitempty"`
	IsAuxHeat          *bool      `json:"is_aux_heat,omitempty"`
	MaxHumidity        *int       `json:"max_humidity,omitempty"`
	MaxTemp            *float64   `json:"max_temp,omitempty"`
	MinHumidity        *int       `json:"min_humidity,omitempty"`
	MinTemp            *float64   `json:"min_temp,omitempty"`
	Precision          *float64   `json:"precision,omitempty"`
	PresetMode         *string    `json:"preset_mode,omitempty"`
	PresetModes        *[]string  `json:"preset_modes,omitempty"`
	SupportedFeatures  *int       `json:"supported_features,omitempty"`
	SwingMode          *string    `json:"swing_mode,omitempty"`
	SwingModes         *[]string  `json:"swing_modes,omitempty"`
	TargetHumidity     *float64   `json:"target_humidity,omitempty"`
	TargetTemp         *float64   `json:"target_temp,omitempty"`
	TargetTempHigh     *float64   `json:"target_temp_high,omitempty"`
	TargetTempLow      *float64   `json:"target_temp_low,omitempty"`
	TargetTempStep     *float64   `json:"target_temp_step,omitempty"`
	TemperatureUnit    *string    `json:"temperature_unit,omitempty"`
}

func GetClimate(attributes map[string]interface{}) *Climate {
	var c Climate
	fillFields(&c, attributes)
	c.Additional = attributes
	return &c
}
