package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Select struct {
	Additional    map[string]interface{} `json:"additional,omitempty"`
	CurrentOption *string                `json:"current_option,omitempty"`
	FriendlyName  *string                `json:"friendly_name,omitempty"`
	Options       *[]string              `json:"options,omitempty"`
}

func GetSelect(attributes map[string]interface{}) *Select {
	var s Select
	FillFields(&s, attributes)
	s.Additional = attributes
	return &s
}
