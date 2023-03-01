package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Cover struct {
	Additional               Additional `json:"additional,omitempty"`
	CurrentCoverPosition     *int       `json:"current_cover_position,omitempty"`
	CurrentCoverTiltPosition *int       `json:"current_cover_tilt_position,omitempty"`
	DeviceClass              *string    `json:"device_class,omitempty"`
	FriendlyName             *string    `json:"friendly_name,omitempty"`
	IsClosed                 *bool      `json:"is_closed,omitempty"`
	IsClosing                *bool      `json:"is_closing,omitempty"`
	IsOpening                *bool      `json:"is_opening,omitempty"`
	SupportedFeatures        *int       `json:"supported_features,omitempty"`
}

func GetCover(attributes map[string]interface{}) *Cover {
	var c Cover
	fillFields(&c, attributes)
	c.Additional = attributes
	return &c
}
