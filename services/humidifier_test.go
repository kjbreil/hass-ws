package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestHumidifierSetHumidity_JSON(t *testing.T) {
	humidity := 1.2

	tests := []struct {
		name   string
		fields *HumidifierSetHumidity
		want   string
	}{{
		fields: NewHumidifierSetHumidity(Targets("climate.kitchen"), &HumidifierSetHumidityParams{Humidity: &humidity}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"humidifier\",\"service\":\"set_humidity\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"humidity\":1.2}}",
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
func TestHumidifierSetMode_JSON(t *testing.T) {
	mode := "data"

	tests := []struct {
		name   string
		fields *HumidifierSetMode
		want   string
	}{{
		fields: NewHumidifierSetMode(Targets("climate.kitchen"), &HumidifierSetModeParams{Mode: &mode}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"humidifier\",\"service\":\"set_mode\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"mode\":\"data\"}}",
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
func TestHumidifierToggle_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *HumidifierToggle
		want   string
	}{{
		fields: NewHumidifierToggle(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"humidifier\",\"service\":\"toggle\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestHumidifierTurnOff_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *HumidifierTurnOff
		want   string
	}{{
		fields: NewHumidifierTurnOff(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"humidifier\",\"service\":\"turn_off\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestHumidifierTurnOn_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *HumidifierTurnOn
		want   string
	}{{
		fields: NewHumidifierTurnOn(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"humidifier\",\"service\":\"turn_on\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
