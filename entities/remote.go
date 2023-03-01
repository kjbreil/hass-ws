package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Remote struct {
	ActivityList    *[]string  `json:"activity_list,omitempty"`
	Additional      Additional `json:"additional,omitempty"`
	CurrentActivity *string    `json:"current_activity,omitempty"`
	FriendlyName    *string    `json:"friendly_name,omitempty"`
}

func GetRemote(attributes map[string]interface{}) *Remote {
	var r Remote
	fillFields(&r, attributes)
	r.Additional = attributes
	return &r
}
