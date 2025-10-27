package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Lock struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetLock(attributes map[string]interface{}) *Lock {
	var l Lock
	fillFields(&l, attributes)
	l.Additional = attributes
	return &l
}
