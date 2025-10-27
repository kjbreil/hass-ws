package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type DeviceTracker struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetDeviceTracker(attributes map[string]interface{}) *DeviceTracker {
	var d DeviceTracker
	fillFields(&d, attributes)
	d.Additional = attributes
	return &d
}
