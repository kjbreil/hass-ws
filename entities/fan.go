package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Fan struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetFan(attributes map[string]interface{}) *Fan {
	var f Fan
	fillFields(&f, attributes)
	f.Additional = attributes
	return &f
}
