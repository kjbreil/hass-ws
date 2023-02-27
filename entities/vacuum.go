package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Vacuum struct {
	BatteryIcon  *string   `json:"battery_icon,omitempty"`
	BatteryLevel *int      `json:"battery_level,omitempty"`
	Error        *string   `json:"error,omitempty"`
	FanSpeed     *string   `json:"fan_speed,omitempty"`
	FanSpeedList *[]string `json:"fan_speed_list,omitempty"`
	Name         *string   `json:"name,omitempty"`
	State        *string   `json:"state,omitempty"`
}
