package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Lock struct {
	ChangedBy   *string `json:"changed_by,omitempty"`
	CodeFormat  *string `json:"code_format,omitempty"`
	IsJammed    *bool   `json:"is_jammed,omitempty"`
	IsLocked    *bool   `json:"is_locked,omitempty"`
	IsLocking   *bool   `json:"is_locking,omitempty"`
	IsUnlocking *bool   `json:"is_unlocking,omitempty"`
}
