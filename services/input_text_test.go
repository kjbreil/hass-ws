package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestInputTextReload_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *InputTextReload
		want   string
	}{{
		fields: NewInputTextReload(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_text\",\"service\":\"reload\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestInputTextSetValue_JSON(t *testing.T) {
	value := "data"

	tests := []struct {
		name   string
		fields *InputTextSetValue
		want   string
	}{{
		fields: NewInputTextSetValue(Targets("climate.kitchen"), &InputTextSetValueParams{Value: &value}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_text\",\"service\":\"set_value\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"value\":\"data\"}}",
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
