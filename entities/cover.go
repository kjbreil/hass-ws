package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Cover struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetCover(attributes map[string]interface{}) *Cover {
	var c Cover
	fillFields(&c, attributes)
	c.Additional = attributes
	return &c
}
