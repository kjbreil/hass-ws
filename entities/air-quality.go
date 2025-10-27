package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type AirQuality struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetAirQuality(attributes map[string]interface{}) *AirQuality {
	var a AirQuality
	fillFields(&a, attributes)
	a.Additional = attributes
	return &a
}
