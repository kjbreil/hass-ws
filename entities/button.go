package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Button struct {
	Additional   map[string]interface{} `json:"additional,omitempty"`
	FriendlyName *string                `json:"friendly_name,omitempty"`
}

func GetButton(attributes map[string]interface{}) *Button {
	var b Button
	FillFields(&b, attributes)
	b.Additional = attributes
	return &b
}
