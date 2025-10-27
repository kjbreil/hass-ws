package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Sensor struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetSensor(attributes map[string]interface{}) *Sensor {
	var s Sensor
	fillFields(&s, attributes)
	s.Additional = attributes
	return &s
}
