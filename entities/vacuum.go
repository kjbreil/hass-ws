package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Vacuum struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetVacuum(attributes map[string]interface{}) *Vacuum {
	var v Vacuum
	fillFields(&v, attributes)
	v.Additional = attributes
	return &v
}
