package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Introduction struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetIntroduction(attributes map[string]interface{}) *Introduction {
	var i Introduction
	fillFields(&i, attributes)
	i.Additional = attributes
	return &i
}
