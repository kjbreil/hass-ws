package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Light struct {
	Brightness          *int       `json:"brightness,omitempty"`
	ColorMode           *string    `json:"color_mode,omitempty"`
	ColorTemp           *int       `json:"color_temp,omitempty"`
	Effect              *string    `json:"effect,omitempty"`
	EffectList          *[]string  `json:"effect_list,omitempty"`
	HsColor             *[]float64 `json:"hs_color,omitempty"`
	IsOn                *bool      `json:"is_on,omitempty"`
	MaxMireds           *int       `json:"max_mireds,omitempty"`
	MinMireds           *int       `json:"min_mireds,omitempty"`
	RgbColor            *[]float64 `json:"rgb_color,omitempty"`
	RgbwColor           *[]float64 `json:"rgbw_color,omitempty"`
	RgbwwColor          *[]float64 `json:"rgbww_color,omitempty"`
	SupportedColorModes *[]string  `json:"supported_color_modes,omitempty"`
	SupportedFeatures   *int       `json:"supported_features,omitempty"`
	WhiteValue          *int       `json:"white_value,omitempty"`
	XyColor             *[]float64 `json:"xy_color,omitempty"`
}
