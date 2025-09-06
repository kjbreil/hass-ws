package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestCounterDecrement_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CounterDecrement
		want   string
	}{{
		fields: NewCounterDecrement(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"counter\",\"service\":\"decrement\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCounterIncrement_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CounterIncrement
		want   string
	}{{
		fields: NewCounterIncrement(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"counter\",\"service\":\"increment\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCounterReset_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CounterReset
		want   string
	}{{
		fields: NewCounterReset(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"counter\",\"service\":\"reset\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCounterSetValue_JSON(t *testing.T) {
	value := 1.2

	tests := []struct {
		name   string
		fields *CounterSetValue
		want   string
	}{{
		fields: NewCounterSetValue(Targets("climate.kitchen")).Value(value),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"counter\",\"service\":\"set_value\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"value\":1.2}}",
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
