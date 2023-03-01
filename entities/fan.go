package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Fan struct {
	Additional        Additional `json:"additional,omitempty"`
	CurrentDirection  *string    `json:"current_direction,omitempty"`
	FriendlyName      *string    `json:"friendly_name,omitempty"`
	IsOn              *bool      `json:"is_on,omitempty"`
	Oscillating       *bool      `json:"oscillating,omitempty"`
	Percentage        *int       `json:"percentage,omitempty"`
	PresetMode        *string    `json:"preset_mode,omitempty"`
	PresetModes       *[]string  `json:"preset_modes,omitempty"`
	SpeedCount        *int       `json:"speed_count,omitempty"`
	SupportedFeatures *int       `json:"supported_features,omitempty"`
}

func GetFan(attributes map[string]interface{}) *Fan {
	var f Fan
	fillFields(&f, attributes)
	f.Additional = attributes
	return &f
}
