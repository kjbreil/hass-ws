package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestInputSelectReload_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *InputSelectReload
		want   string
	}{{
		fields: NewInputSelectReload(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_select\",\"service\":\"reload\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestInputSelectSelectFirst_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *InputSelectSelectFirst
		want   string
	}{{
		fields: NewInputSelectSelectFirst(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_select\",\"service\":\"select_first\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestInputSelectSelectLast_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *InputSelectSelectLast
		want   string
	}{{
		fields: NewInputSelectSelectLast(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_select\",\"service\":\"select_last\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestInputSelectSelectNext_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *InputSelectSelectNext
		want   string
	}{{
		fields: NewInputSelectSelectNext(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_select\",\"service\":\"select_next\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestInputSelectSelectOption_JSON(t *testing.T) {
	option := "data"

	tests := []struct {
		name   string
		fields *InputSelectSelectOption
		want   string
	}{{
		fields: NewInputSelectSelectOption(Targets("climate.kitchen"), &InputSelectSelectOptionParams{Option: &option}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_select\",\"service\":\"select_option\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"option\":\"data\"}}",
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
func TestInputSelectSelectPrevious_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *InputSelectSelectPrevious
		want   string
	}{{
		fields: NewInputSelectSelectPrevious(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_select\",\"service\":\"select_previous\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestInputSelectSetOptions_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *InputSelectSetOptions
		want   string
	}{{
		fields: NewInputSelectSetOptions(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_select\",\"service\":\"set_options\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
