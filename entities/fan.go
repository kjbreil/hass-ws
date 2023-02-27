package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Fan struct {
	CurrentDirection  *string   `json:"current_direction,omitempty"`
	IsOn              *bool     `json:"is_on,omitempty"`
	Oscillating       *bool     `json:"oscillating,omitempty"`
	Percentage        *int      `json:"percentage,omitempty"`
	PresetMode        *string   `json:"preset_mode,omitempty"`
	PresetModes       *[]string `json:"preset_modes,omitempty"`
	SpeedCount        *int      `json:"speed_count,omitempty"`
	SupportedFeatures *int      `json:"supported_features,omitempty"`
}
