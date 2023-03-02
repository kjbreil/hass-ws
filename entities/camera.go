package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Camera struct {
	Additional             Additional `json:"additional,omitempty"`
	Brand                  *string    `json:"brand,omitempty"`
	FrameInterval          *float64   `json:"frame_interval,omitempty"`
	FriendlyName           *string    `json:"friendly_name,omitempty"`
	FrontendStreamType     *string    `json:"frontend_stream_type,omitempty"`
	IsOn                   *bool      `json:"is_on,omitempty"`
	IsRecording            *bool      `json:"is_recording,omitempty"`
	IsStreaming            *bool      `json:"is_streaming,omitempty"`
	Model                  *string    `json:"model,omitempty"`
	MotionDetectionEnabled *bool      `json:"motion_detection_enabled,omitempty"`
}

func GetCamera(attributes map[string]interface{}) *Camera {
	var c Camera
	fillFields(&c, attributes)
	c.Additional = attributes
	return &c
}
