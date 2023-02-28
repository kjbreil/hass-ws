package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Lock struct {
	Additional   map[string]interface{} `json:"additional,omitempty"`
	ChangedBy    *string                `json:"changed_by,omitempty"`
	CodeFormat   *string                `json:"code_format,omitempty"`
	FriendlyName *string                `json:"friendly_name,omitempty"`
	IsJammed     *bool                  `json:"is_jammed,omitempty"`
	IsLocked     *bool                  `json:"is_locked,omitempty"`
	IsLocking    *bool                  `json:"is_locking,omitempty"`
	IsUnlocking  *bool                  `json:"is_unlocking,omitempty"`
}

func GetLock(attributes map[string]interface{}) *Lock {
	var l Lock
	fillFields(&l, attributes)
	l.Additional = attributes
	return &l
}
