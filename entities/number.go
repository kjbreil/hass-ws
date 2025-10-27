package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Number struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetNumber(attributes map[string]interface{}) *Number {
	var n Number
	fillFields(&n, attributes)
	n.Additional = attributes
	return &n
}
