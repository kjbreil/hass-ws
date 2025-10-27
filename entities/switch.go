package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Switch struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetSwitch(attributes map[string]interface{}) *Switch {
	var s Switch
	fillFields(&s, attributes)
	s.Additional = attributes
	return &s
}
