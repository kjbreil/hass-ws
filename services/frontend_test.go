package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestFrontendReloadThemes_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *FrontendReloadThemes
		want   string
	}{{
		fields: NewFrontendReloadThemes(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"frontend\",\"service\":\"reload_themes\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestFrontendSetTheme_JSON(t *testing.T) {
	mode := Modedark

	tests := []struct {
		name   string
		fields *FrontendSetTheme
		want   string
	}{{
		fields: NewFrontendSetTheme(Targets("climate.kitchen")).Mode(mode),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"frontend\",\"service\":\"set_theme\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"mode\":\"dark\"}}",
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
