package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestSystemLogClear_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *SystemLogClear
		want   string
	}{{
		fields: NewSystemLogClear(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"system_log\",\"service\":\"clear\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestSystemLogWrite_JSON(t *testing.T) {
	level := Levelcritical
	logger := "data"
	message := "data"

	tests := []struct {
		name   string
		fields *SystemLogWrite
		want   string
	}{{
		fields: NewSystemLogWrite(Targets("climate.kitchen"), &SystemLogWriteParams{
			Level:   &level,
			Logger:  &logger,
			Message: &message,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"system_log\",\"service\":\"write\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"level\":\"critical\",\"logger\":\"data\",\"message\":\"data\"}}",
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
