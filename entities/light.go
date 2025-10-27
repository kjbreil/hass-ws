package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Light struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetLight(attributes map[string]interface{}) *Light {
	var l Light
	fillFields(&l, attributes)
	l.Additional = attributes
	return &l
}
