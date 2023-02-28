package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Sensor struct {
	Additional                 map[string]interface{} `json:"additional,omitempty"`
	DeviceClass                *string                `json:"device_class,omitempty"`
	FriendlyName               *string                `json:"friendly_name,omitempty"`
	NativeUnitOfMeasurement    *string                `json:"native_unit_of_measurement,omitempty"`
	Options                    *[]string              `json:"options,omitempty"`
	StateClass                 *string                `json:"state_class,omitempty"`
	SuggestedDisplayPrecision  *int                   `json:"suggested_display_precision,omitempty"`
	SuggestedUnitOfMeasurement *string                `json:"suggested_unit_of_measurement,omitempty"`
}

func GetSensor(attributes map[string]interface{}) *Sensor {
	var s Sensor
	FillFields(&s, attributes)
	s.Additional = attributes
	return &s
}
