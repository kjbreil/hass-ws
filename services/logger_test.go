package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestLoggerSetDefaultLevel_JSON(t *testing.T) {
	level := Levelcritical

	tests := []struct {
		name   string
		fields *LoggerSetDefaultLevel
		want   string
	}{{
		fields: NewLoggerSetDefaultLevel(Targets("climate.kitchen"), &LoggerSetDefaultLevelParams{Level: &level}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"logger\",\"service\":\"set_default_level\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"level\":\"critical\"}}",
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
func TestLoggerSetLevel_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *LoggerSetLevel
		want   string
	}{{
		fields: NewLoggerSetLevel(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"logger\",\"service\":\"set_level\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
