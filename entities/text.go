package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Text struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetText(attributes map[string]interface{}) *Text {
	var t Text
	fillFields(&t, attributes)
	t.Additional = attributes
	return &t
}
