package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Siren struct {
	Additional     Additional `json:"additional,omitempty"`
	AvailableTones *[]string  `json:"available_tones,omitempty"`
	FriendlyName   *string    `json:"friendly_name,omitempty"`
	IsOn           *bool      `json:"is_on,omitempty"`
}

func GetSiren(attributes map[string]interface{}) *Siren {
	var s Siren
	fillFields(&s, attributes)
	s.Additional = attributes
	return &s
}
