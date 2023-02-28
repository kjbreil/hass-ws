package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Number struct {
	Additional              map[string]interface{} `json:"additional,omitempty"`
	DeviceClass             *string                `json:"device_class,omitempty"`
	FriendlyName            *string                `json:"friendly_name,omitempty"`
	Mode                    *string                `json:"mode,omitempty"`
	NativeMaxValue          *float64               `json:"native_max_value,omitempty"`
	NativeMinValue          *float64               `json:"native_min_value,omitempty"`
	NativeStep              *float64               `json:"native_step,omitempty"`
	NativeUnitOfMeasurement *string                `json:"native_unit_of_measurement,omitempty"`
	NativeValue             *float64               `json:"native_value,omitempty"`
}

func GetNumber(attributes map[string]interface{}) *Number {
	var n Number
	fillFields(&n, attributes)
	n.Additional = attributes
	return &n
}
