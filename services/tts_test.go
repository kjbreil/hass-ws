package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestTtsClearCache_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *TtsClearCache
		want   string
	}{{
		fields: NewTtsClearCache(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"tts\",\"service\":\"clear_cache\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestTtsCloudSay_JSON(t *testing.T) {
	language := "data"
	message := "data"

	tests := []struct {
		name   string
		fields *TtsCloudSay
		want   string
	}{{
		fields: NewTtsCloudSay(Targets("climate.kitchen"), &TtsCloudSayParams{
			Language: &language,
			Message:  &message,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"tts\",\"service\":\"cloud_say\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"language\":\"data\",\"message\":\"data\"}}",
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
func TestTtsGoogleTranslateSay_JSON(t *testing.T) {
	language := "data"
	message := "data"

	tests := []struct {
		name   string
		fields *TtsGoogleTranslateSay
		want   string
	}{{
		fields: NewTtsGoogleTranslateSay(Targets("climate.kitchen"), &TtsGoogleTranslateSayParams{
			Language: &language,
			Message:  &message,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"tts\",\"service\":\"google_translate_say\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"language\":\"data\",\"message\":\"data\"}}",
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
