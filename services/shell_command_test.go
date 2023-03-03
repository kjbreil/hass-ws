package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestShellCommandShutdownTyr_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *ShellCommandShutdownTyr
		want   string
	}{{
		fields: NewShellCommandShutdownTyr(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"shell_command\",\"service\":\"shutdown_tyr\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestShellCommandStartNcbp_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *ShellCommandStartNcbp
		want   string
	}{{
		fields: NewShellCommandStartNcbp(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"shell_command\",\"service\":\"start_ncbp\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestShellCommandStopNcbp_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *ShellCommandStopNcbp
		want   string
	}{{
		fields: NewShellCommandStopNcbp(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"shell_command\",\"service\":\"stop_ncbp\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
