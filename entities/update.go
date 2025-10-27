package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Update struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetUpdate(attributes map[string]interface{}) *Update {
	var u Update
	fillFields(&u, attributes)
	u.Additional = attributes
	return &u
}
