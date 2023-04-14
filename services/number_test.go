package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestNumberSetValue_JSON(t *testing.T) {
	value := "data"

	tests := []struct {
		name   string
		fields *NumberSetValue
		want   string
	}{{
		fields: NewNumberSetValue(Targets("climate.kitchen")).Value(value),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"number\",\"service\":\"set_value\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"value\":\"data\"}}",
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
