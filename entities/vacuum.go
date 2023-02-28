package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Vacuum struct {
	Additional   map[string]interface{} `json:"additional,omitempty"`
	BatteryIcon  *string                `json:"battery_icon,omitempty"`
	BatteryLevel *int                   `json:"battery_level,omitempty"`
	Error        *string                `json:"error,omitempty"`
	FanSpeed     *string                `json:"fan_speed,omitempty"`
	FanSpeedList *[]string              `json:"fan_speed_list,omitempty"`
	FriendlyName *string                `json:"friendly_name,omitempty"`
	Name         *string                `json:"name,omitempty"`
	State        *string                `json:"state,omitempty"`
}

func GetVacuum(attributes map[string]interface{}) *Vacuum {
	var v Vacuum
	FillFields(&v, attributes)
	v.Additional = attributes
	return &v
}
