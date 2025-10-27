package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type BinarySensor struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetBinarySensor(attributes map[string]interface{}) *BinarySensor {
	var b BinarySensor
	fillFields(&b, attributes)
	b.Additional = attributes
	return &b
}
