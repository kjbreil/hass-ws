package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestCounterConfigure_JSON(t *testing.T) {
	initial := 1.2
	maximum := 1.2
	minimum := 1.2
	step := 1.2
	value := 1.2

	tests := []struct {
		name   string
		fields *CounterConfigure
		want   string
	}{{
		fields: NewCounterConfigure(Targets("climate.kitchen")).Initial(initial).Maximum(maximum).Minimum(minimum).Step(step).Value(value),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"counter\",\"service\":\"configure\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"initial\":1.2,\"maximum\":1.2,\"minimum\":1.2,\"step\":1.2,\"value\":1.2}}",
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
