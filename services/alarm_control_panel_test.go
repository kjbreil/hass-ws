package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestAlarmControlPanelAlarmArmAway_JSON(t *testing.T) {
	code := "data"

	tests := []struct {
		name   string
		fields *AlarmControlPanelAlarmArmAway
		want   string
	}{{
		fields: NewAlarmControlPanelAlarmArmAway(Targets("climate.kitchen")).Code(code),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"alarm_control_panel\",\"service\":\"alarm_arm_away\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"code\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAlarmControlPanelAlarmArmCustomBypass_JSON(t *testing.T) {
	code := "data"

	tests := []struct {
		name   string
		fields *AlarmControlPanelAlarmArmCustomBypass
		want   string
	}{{
		fields: NewAlarmControlPanelAlarmArmCustomBypass(Targets("climate.kitchen")).Code(code),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"alarm_control_panel\",\"service\":\"alarm_arm_custom_bypass\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"code\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAlarmControlPanelAlarmArmHome_JSON(t *testing.T) {
	code := "data"

	tests := []struct {
		name   string
		fields *AlarmControlPanelAlarmArmHome
		want   string
	}{{
		fields: NewAlarmControlPanelAlarmArmHome(Targets("climate.kitchen")).Code(code),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"alarm_control_panel\",\"service\":\"alarm_arm_home\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"code\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAlarmControlPanelAlarmArmNight_JSON(t *testing.T) {
	code := "data"

	tests := []struct {
		name   string
		fields *AlarmControlPanelAlarmArmNight
		want   string
	}{{
		fields: NewAlarmControlPanelAlarmArmNight(Targets("climate.kitchen")).Code(code),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"alarm_control_panel\",\"service\":\"alarm_arm_night\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"code\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAlarmControlPanelAlarmArmVacation_JSON(t *testing.T) {
	code := "data"

	tests := []struct {
		name   string
		fields *AlarmControlPanelAlarmArmVacation
		want   string
	}{{
		fields: NewAlarmControlPanelAlarmArmVacation(Targets("climate.kitchen")).Code(code),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"alarm_control_panel\",\"service\":\"alarm_arm_vacation\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"code\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAlarmControlPanelAlarmDisarm_JSON(t *testing.T) {
	code := "data"

	tests := []struct {
		name   string
		fields *AlarmControlPanelAlarmDisarm
		want   string
	}{{
		fields: NewAlarmControlPanelAlarmDisarm(Targets("climate.kitchen")).Code(code),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"alarm_control_panel\",\"service\":\"alarm_disarm\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"code\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAlarmControlPanelAlarmTrigger_JSON(t *testing.T) {
	code := "data"

	tests := []struct {
		name   string
		fields *AlarmControlPanelAlarmTrigger
		want   string
	}{{
		fields: NewAlarmControlPanelAlarmTrigger(Targets("climate.kitchen")).Code(code),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"alarm_control_panel\",\"service\":\"alarm_trigger\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"code\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
