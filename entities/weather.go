package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Weather struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetWeather(attributes map[string]interface{}) *Weather {
	var w Weather
	fillFields(&w, attributes)
	w.Additional = attributes
	return &w
}
