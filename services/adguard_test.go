package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestAdguardAddUrl_JSON(t *testing.T) {
	name := "data"
	url := "data"

	tests := []struct {
		name   string
		fields *AdguardAddUrl
		want   string
	}{{
		fields: NewAdguardAddUrl(Targets("climate.kitchen"), &AdguardAddUrlParams{
			Name: &name,
			Url:  &url,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"adguard\",\"service\":\"add_url\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"name\":\"data\",\"url\":\"data\"}}",
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
func TestAdguardDisableUrl_JSON(t *testing.T) {
	url := "data"

	tests := []struct {
		name   string
		fields *AdguardDisableUrl
		want   string
	}{{
		fields: NewAdguardDisableUrl(Targets("climate.kitchen"), &AdguardDisableUrlParams{Url: &url}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"adguard\",\"service\":\"disable_url\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"url\":\"data\"}}",
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
func TestAdguardEnableUrl_JSON(t *testing.T) {
	url := "data"

	tests := []struct {
		name   string
		fields *AdguardEnableUrl
		want   string
	}{{
		fields: NewAdguardEnableUrl(Targets("climate.kitchen"), &AdguardEnableUrlParams{Url: &url}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"adguard\",\"service\":\"enable_url\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"url\":\"data\"}}",
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
func TestAdguardRefresh_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *AdguardRefresh
		want   string
	}{{
		fields: NewAdguardRefresh(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"adguard\",\"service\":\"refresh\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestAdguardRemoveUrl_JSON(t *testing.T) {
	url := "data"

	tests := []struct {
		name   string
		fields *AdguardRemoveUrl
		want   string
	}{{
		fields: NewAdguardRemoveUrl(Targets("climate.kitchen"), &AdguardRemoveUrlParams{Url: &url}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"adguard\",\"service\":\"remove_url\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"url\":\"data\"}}",
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
