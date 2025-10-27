package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Humidifier struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetHumidifier(attributes map[string]interface{}) *Humidifier {
	var h Humidifier
	fillFields(&h, attributes)
	h.Additional = attributes
	return &h
}
