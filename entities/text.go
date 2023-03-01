package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Text struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Mode         *string    `json:"mode,omitempty"`
	NativeMax    *int       `json:"native_max,omitempty"`
	NativeMin    *int       `json:"native_min,omitempty"`
	NativeValue  *string    `json:"native_value,omitempty"`
	Pattern      *string    `json:"pattern,omitempty"`
}

func GetText(attributes map[string]interface{}) *Text {
	var t Text
	fillFields(&t, attributes)
	t.Additional = attributes
	return &t
}
