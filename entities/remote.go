package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Remote struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetRemote(attributes map[string]interface{}) *Remote {
	var r Remote
	fillFields(&r, attributes)
	r.Additional = attributes
	return &r
}
