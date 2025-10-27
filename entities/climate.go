package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Climate struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetClimate(attributes map[string]interface{}) *Climate {
	var c Climate
	fillFields(&c, attributes)
	c.Additional = attributes
	return &c
}
