package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestInputNumberDecrement_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *InputNumberDecrement
		want   string
	}{{
		fields: NewInputNumberDecrement(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_number\",\"service\":\"decrement\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestInputNumberIncrement_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *InputNumberIncrement
		want   string
	}{{
		fields: NewInputNumberIncrement(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_number\",\"service\":\"increment\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestInputNumberReload_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *InputNumberReload
		want   string
	}{{
		fields: NewInputNumberReload(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_number\",\"service\":\"reload\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestInputNumberSetValue_JSON(t *testing.T) {
	value := 1.2

	tests := []struct {
		name   string
		fields *InputNumberSetValue
		want   string
	}{{
		fields: NewInputNumberSetValue(Targets("climate.kitchen")).Value(value),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_number\",\"service\":\"set_value\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"value\":1.2}}",
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
