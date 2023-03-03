package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestOpnsenseCloseNotice_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *OpnsenseCloseNotice
		want   string
	}{{
		fields: NewOpnsenseCloseNotice(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"opnsense\",\"service\":\"close_notice\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestOpnsenseFileNotice_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *OpnsenseFileNotice
		want   string
	}{{
		fields: NewOpnsenseFileNotice(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"opnsense\",\"service\":\"file_notice\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestOpnsenseRestartService_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *OpnsenseRestartService
		want   string
	}{{
		fields: NewOpnsenseRestartService(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"opnsense\",\"service\":\"restart_service\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestOpnsenseSendWol_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *OpnsenseSendWol
		want   string
	}{{
		fields: NewOpnsenseSendWol(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"opnsense\",\"service\":\"send_wol\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestOpnsenseStartService_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *OpnsenseStartService
		want   string
	}{{
		fields: NewOpnsenseStartService(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"opnsense\",\"service\":\"start_service\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestOpnsenseStopService_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *OpnsenseStopService
		want   string
	}{{
		fields: NewOpnsenseStopService(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"opnsense\",\"service\":\"stop_service\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestOpnsenseSystemHalt_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *OpnsenseSystemHalt
		want   string
	}{{
		fields: NewOpnsenseSystemHalt(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"opnsense\",\"service\":\"system_halt\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestOpnsenseSystemReboot_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *OpnsenseSystemReboot
		want   string
	}{{
		fields: NewOpnsenseSystemReboot(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"opnsense\",\"service\":\"system_reboot\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
