package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Cover struct {
	CurrentCoverPosition     *int    `json:"current_cover_position,omitempty"`
	CurrentCoverTiltPosition *int    `json:"current_cover_tilt_position,omitempty"`
	DeviceClass              *string `json:"device_class,omitempty"`
	IsClosed                 *bool   `json:"is_closed,omitempty"`
	IsClosing                *bool   `json:"is_closing,omitempty"`
	IsOpening                *bool   `json:"is_opening,omitempty"`
	SupportedFeatures        *int    `json:"supported_features,omitempty"`
}
