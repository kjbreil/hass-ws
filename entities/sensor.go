package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Sensor struct {
	DeviceClass                *string   `json:"device_class,omitempty"`
	NativeUnitOfMeasurement    *string   `json:"native_unit_of_measurement,omitempty"`
	Options                    *[]string `json:"options,omitempty"`
	StateClass                 *string   `json:"state_class,omitempty"`
	SuggestedDisplayPrecision  *int      `json:"suggested_display_precision,omitempty"`
	SuggestedUnitOfMeasurement *string   `json:"suggested_unit_of_measurement,omitempty"`
}
