package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestSelectSelectOption_JSON(t *testing.T) {
	option := "data"

	tests := []struct {
		name   string
		fields *SelectSelectOption
		want   string
	}{{
		fields: NewSelectSelectOption(Targets("climate.kitchen")).Option(option),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"select\",\"service\":\"select_option\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"option\":\"data\"}}",
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
