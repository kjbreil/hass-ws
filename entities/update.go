package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Update struct {
	AutoUpdate       *bool   `json:"auto_update,omitempty"`
	InProgress       *int    `json:"in_progress,omitempty"`
	InstalledVersion *string `json:"installed_version,omitempty"`
	LatestVersion    *string `json:"latest_version,omitempty"`
	ReleaseSummary   *string `json:"release_summary,omitempty"`
	ReleaseUrl       *string `json:"release_url,omitempty"`
	Title            *string `json:"title,omitempty"`
}
