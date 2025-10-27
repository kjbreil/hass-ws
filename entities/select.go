package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Select struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetSelect(attributes map[string]interface{}) *Select {
	var s Select
	fillFields(&s, attributes)
	s.Additional = attributes
	return &s
}
