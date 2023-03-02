package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Button struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetButton(attributes map[string]interface{}) *Button {
	var b Button
	fillFields(&b, attributes)
	b.Additional = attributes
	return &b
}
