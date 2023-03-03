package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestRecorderDisable_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *RecorderDisable
		want   string
	}{{
		fields: NewRecorderDisable(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"recorder\",\"service\":\"disable\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestRecorderEnable_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *RecorderEnable
		want   string
	}{{
		fields: NewRecorderEnable(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"recorder\",\"service\":\"enable\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestRecorderPurge_JSON(t *testing.T) {
	keepDays := 1.2

	tests := []struct {
		name   string
		fields *RecorderPurge
		want   string
	}{{
		fields: NewRecorderPurge(Targets("climate.kitchen"), &RecorderPurgeParams{KeepDays: &keepDays}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"recorder\",\"service\":\"purge\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"keep_days\":1.2}}",
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
func TestRecorderPurgeEntities_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *RecorderPurgeEntities
		want   string
	}{{
		fields: NewRecorderPurgeEntities(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"recorder\",\"service\":\"purge_entities\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
