package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Siren struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetSiren(attributes map[string]interface{}) *Siren {
	var s Siren
	fillFields(&s, attributes)
	s.Additional = attributes
	return &s
}
