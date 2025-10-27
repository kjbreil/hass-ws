package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Camera struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetCamera(attributes map[string]interface{}) *Camera {
	var c Camera
	fillFields(&c, attributes)
	c.Additional = attributes
	return &c
}
