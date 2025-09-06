package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestLightToggle_JSON(t *testing.T) {
	brightnessPct := 1.2
	colorTempKelvin := 1.2
	effect := "data"
	transition := 1.2

	tests := []struct {
		name   string
		fields *LightToggle
		want   string
	}{{
		fields: NewLightToggle(Targets("climate.kitchen")).BrightnessPct(brightnessPct).ColorTempKelvin(colorTempKelvin).Effect(effect).Transition(transition),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"light\",\"service\":\"toggle\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"brightness_pct\":1.2,\"color_temp_kelvin\":1.2,\"effect\":\"data\",\"transition\":1.2}}",
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
func TestLightTurnOff_JSON(t *testing.T) {
	transition := 1.2

	tests := []struct {
		name   string
		fields *LightTurnOff
		want   string
	}{{
		fields: NewLightTurnOff(Targets("climate.kitchen")).Transition(transition),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"light\",\"service\":\"turn_off\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"transition\":1.2}}",
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
func TestLightTurnOn_JSON(t *testing.T) {
	brightnessPct := 1.2
	brightnessStepPct := 1.2
	colorTempKelvin := 1.2
	effect := "data"
	transition := 1.2

	tests := []struct {
		name   string
		fields *LightTurnOn
		want   string
	}{{
		fields: NewLightTurnOn(Targets("climate.kitchen")).BrightnessPct(brightnessPct).BrightnessStepPct(brightnessStepPct).ColorTempKelvin(colorTempKelvin).Effect(effect).Transition(transition),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"light\",\"service\":\"turn_on\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"brightness_pct\":1.2,\"brightness_step_pct\":1.2,\"color_temp_kelvin\":1.2,\"effect\":\"data\",\"transition\":1.2}}",
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
