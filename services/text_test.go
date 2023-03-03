package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestTextSetValue_JSON(t *testing.T) {
	value := "data"

	tests := []struct {
		name   string
		fields *TextSetValue
		want   string
	}{{
		fields: NewTextSetValue(Targets("climate.kitchen"), &TextSetValueParams{Value: &value}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"text\",\"service\":\"set_value\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"value\":\"data\"}}",
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
