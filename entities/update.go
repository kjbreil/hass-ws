package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Update struct {
	Additional       Additional `json:"additional,omitempty"`
	AutoUpdate       *bool      `json:"auto_update,omitempty"`
	FriendlyName     *string    `json:"friendly_name,omitempty"`
	InProgress       *int       `json:"in_progress,omitempty"`
	InstalledVersion *string    `json:"installed_version,omitempty"`
	LatestVersion    *string    `json:"latest_version,omitempty"`
	ReleaseSummary   *string    `json:"release_summary,omitempty"`
	ReleaseUrl       *string    `json:"release_url,omitempty"`
	Title            *string    `json:"title,omitempty"`
}

func GetUpdate(attributes map[string]interface{}) *Update {
	var u Update
	fillFields(&u, attributes)
	u.Additional = attributes
	return &u
}
