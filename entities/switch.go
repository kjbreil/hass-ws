package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Switch struct {
	Additional     Additional `json:"additional,omitempty"`
	CurrentPowerW  *float64   `json:"current_power_w,omitempty"`
	FriendlyName   *string    `json:"friendly_name,omitempty"`
	IsOn           *bool      `json:"is_on,omitempty"`
	TodayEnergyKwh *float64   `json:"today_energy_kwh,omitempty"`
}

func GetSwitch(attributes map[string]interface{}) *Switch {
	var s Switch
	fillFields(&s, attributes)
	s.Additional = attributes
	return &s
}
