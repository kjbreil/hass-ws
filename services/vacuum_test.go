package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestVacuumCleanSpot_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *VacuumCleanSpot
		want   string
	}{{
		fields: NewVacuumCleanSpot(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"vacuum\",\"service\":\"clean_spot\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestVacuumLocate_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *VacuumLocate
		want   string
	}{{
		fields: NewVacuumLocate(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"vacuum\",\"service\":\"locate\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestVacuumPause_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *VacuumPause
		want   string
	}{{
		fields: NewVacuumPause(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"vacuum\",\"service\":\"pause\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestVacuumReturnToBase_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *VacuumReturnToBase
		want   string
	}{{
		fields: NewVacuumReturnToBase(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"vacuum\",\"service\":\"return_to_base\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestVacuumSendCommand_JSON(t *testing.T) {
	command := "data"

	tests := []struct {
		name   string
		fields *VacuumSendCommand
		want   string
	}{{
		fields: NewVacuumSendCommand(Targets("climate.kitchen")).Command(command),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"vacuum\",\"service\":\"send_command\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"command\":\"data\"}}",
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
func TestVacuumSetFanSpeed_JSON(t *testing.T) {
	fanSpeed := "data"

	tests := []struct {
		name   string
		fields *VacuumSetFanSpeed
		want   string
	}{{
		fields: NewVacuumSetFanSpeed(Targets("climate.kitchen")).FanSpeed(fanSpeed),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"vacuum\",\"service\":\"set_fan_speed\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"fan_speed\":\"data\"}}",
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
func TestVacuumStart_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *VacuumStart
		want   string
	}{{
		fields: NewVacuumStart(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"vacuum\",\"service\":\"start\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestVacuumStartPause_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *VacuumStartPause
		want   string
	}{{
		fields: NewVacuumStartPause(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"vacuum\",\"service\":\"start_pause\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestVacuumStop_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *VacuumStop
		want   string
	}{{
		fields: NewVacuumStop(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"vacuum\",\"service\":\"stop\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestVacuumToggle_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *VacuumToggle
		want   string
	}{{
		fields: NewVacuumToggle(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"vacuum\",\"service\":\"toggle\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestVacuumTurnOff_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *VacuumTurnOff
		want   string
	}{{
		fields: NewVacuumTurnOff(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"vacuum\",\"service\":\"turn_off\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestVacuumTurnOn_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *VacuumTurnOn
		want   string
	}{{
		fields: NewVacuumTurnOn(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"vacuum\",\"service\":\"turn_on\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
