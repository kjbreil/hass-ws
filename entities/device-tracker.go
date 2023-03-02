package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type DeviceTracker struct {
	Additional       Additional `json:"additional,omitempty"`
	BatteryLevel     *int       `json:"battery_level,omitempty"`
	FriendlyName     *string    `json:"friendly_name,omitempty"`
	Hostname         *string    `json:"hostname,omitempty"`
	IpAddress        *string    `json:"ip_address,omitempty"`
	IsConnected      *bool      `json:"is_connected,omitempty"`
	Latitude         *float64   `json:"latitude,omitempty"`
	LocationAccuracy *int       `json:"location_accuracy,omitempty"`
	LocationName     *string    `json:"location_name,omitempty"`
	Longitude        *float64   `json:"longitude,omitempty"`
	MacAddress       *string    `json:"mac_address,omitempty"`
	SourceType       *string    `json:"source_type,omitempty"`
}

func GetDeviceTracker(attributes map[string]interface{}) *DeviceTracker {
	var d DeviceTracker
	fillFields(&d, attributes)
	d.Additional = attributes
	return &d
}
