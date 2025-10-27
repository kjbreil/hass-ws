package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type AlarmControlPanel struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetAlarmControlPanel(attributes map[string]interface{}) *AlarmControlPanel {
	var a AlarmControlPanel
	fillFields(&a, attributes)
	a.Additional = attributes
	return &a
}
