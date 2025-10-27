package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestTimerCancel_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *TimerCancel
		want   string
	}{{
		fields: NewTimerCancel(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"timer\",\"service\":\"cancel\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestTimerFinish_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *TimerFinish
		want   string
	}{{
		fields: NewTimerFinish(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"timer\",\"service\":\"finish\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestTimerPause_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *TimerPause
		want   string
	}{{
		fields: NewTimerPause(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"timer\",\"service\":\"pause\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestTimerReload_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *TimerReload
		want   string
	}{{
		fields: NewTimerReload(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"timer\",\"service\":\"reload\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestTimerStart_JSON(t *testing.T) {
	duration := "data"

	tests := []struct {
		name   string
		fields *TimerStart
		want   string
	}{{
		fields: NewTimerStart(Targets("climate.kitchen")).Duration(duration),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"timer\",\"service\":\"start\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"duration\":\"data\"}}",
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
