package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type BinarySensor struct {
	Additional   map[string]interface{} `json:"additional,omitempty"`
	DeviceClass  *string                `json:"device_class,omitempty"`
	FriendlyName *string                `json:"friendly_name,omitempty"`
	IsOn         *bool                  `json:"is_on,omitempty"`
}

func GetBinarySensor(attributes map[string]interface{}) *BinarySensor {
	var b BinarySensor
	fillFields(&b, attributes)
	b.Additional = attributes
	return &b
}
