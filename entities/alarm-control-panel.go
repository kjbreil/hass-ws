package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type AlarmControlPanel struct {
	Additional   map[string]interface{} `json:"additional,omitempty"`
	ChangedBy    *string                `json:"changed_by,omitempty"`
	CodeFormat   *string                `json:"code_format,omitempty"`
	FriendlyName *string                `json:"friendly_name,omitempty"`
	State        *string                `json:"state,omitempty"`
}

func GetAlarmControlPanel(attributes map[string]interface{}) *AlarmControlPanel {
	var a AlarmControlPanel
	fillFields(&a, attributes)
	a.Additional = attributes
	return &a
}
