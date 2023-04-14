package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestSirenToggle_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *SirenToggle
		want   string
	}{{
		fields: NewSirenToggle(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"siren\",\"service\":\"toggle\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestSirenTurnOff_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *SirenTurnOff
		want   string
	}{{
		fields: NewSirenTurnOff(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"siren\",\"service\":\"turn_off\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestSirenTurnOn_JSON(t *testing.T) {
	duration := "data"
	tone := "data"
	volumeLevel := 1.2

	tests := []struct {
		name   string
		fields *SirenTurnOn
		want   string
	}{{
		fields: NewSirenTurnOn(Targets("climate.kitchen")).Duration(duration).Tone(tone).VolumeLevel(volumeLevel),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"siren\",\"service\":\"turn_on\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"duration\":\"data\",\"tone\":\"data\",\"volume_level\":1.2}}",
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
