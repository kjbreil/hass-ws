package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestUpdateClearSkipped_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *UpdateClearSkipped
		want   string
	}{{
		fields: NewUpdateClearSkipped(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"update\",\"service\":\"clear_skipped\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestUpdateInstall_JSON(t *testing.T) {
	version := "data"

	tests := []struct {
		name   string
		fields *UpdateInstall
		want   string
	}{{
		fields: NewUpdateInstall(Targets("climate.kitchen"), &UpdateInstallParams{Version: &version}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"update\",\"service\":\"install\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"version\":\"data\"}}",
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
func TestUpdateSkip_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *UpdateSkip
		want   string
	}{{
		fields: NewUpdateSkip(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"update\",\"service\":\"skip\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
