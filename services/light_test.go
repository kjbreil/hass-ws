package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestLightToggle_JSON(t *testing.T) {
	brightness := 1.2
	brightnessPct := 1.2
	colorName := ColorNamealiceblue
	colorTemp := 1.2
	effect := "data"
	flash := Flashlong
	kelvin := 1.2
	profile := "data"
	transition := 1.2

	tests := []struct {
		name   string
		fields *LightToggle
		want   string
	}{{
		fields: NewLightToggle(Targets("climate.kitchen"), &LightToggleParams{
			Brightness:    &brightness,
			BrightnessPct: &brightnessPct,
			ColorName:     &colorName,
			ColorTemp:     &colorTemp,
			Effect:        &effect,
			Flash:         &flash,
			Kelvin:        &kelvin,
			Profile:       &profile,
			Transition:    &transition,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"light\",\"service\":\"toggle\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"brightness\":1.2,\"brightness_pct\":1.2,\"color_name\":\"aliceblue\",\"color_temp\":1.2,\"effect\":\"data\",\"flash\":\"long\",\"kelvin\":1.2,\"profile\":\"data\",\"transition\":1.2}}",
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
	flash := Flashlong
	transition := 1.2

	tests := []struct {
		name   string
		fields *LightTurnOff
		want   string
	}{{
		fields: NewLightTurnOff(Targets("climate.kitchen"), &LightTurnOffParams{
			Flash:      &flash,
			Transition: &transition,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"light\",\"service\":\"turn_off\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"flash\":\"long\",\"transition\":1.2}}",
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
	brightness := 1.2
	brightnessPct := 1.2
	brightnessStep := 1.2
	brightnessStepPct := 1.2
	colorName := ColorNamealiceblue
	colorTemp := 1.2
	effect := "data"
	flash := Flashlong
	kelvin := 1.2
	profile := "data"
	transition := 1.2
	white := 1.2

	tests := []struct {
		name   string
		fields *LightTurnOn
		want   string
	}{{
		fields: NewLightTurnOn(Targets("climate.kitchen"), &LightTurnOnParams{
			Brightness:        &brightness,
			BrightnessPct:     &brightnessPct,
			BrightnessStep:    &brightnessStep,
			BrightnessStepPct: &brightnessStepPct,
			ColorName:         &colorName,
			ColorTemp:         &colorTemp,
			Effect:            &effect,
			Flash:             &flash,
			Kelvin:            &kelvin,
			Profile:           &profile,
			Transition:        &transition,
			White:             &white,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"light\",\"service\":\"turn_on\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"brightness\":1.2,\"brightness_pct\":1.2,\"brightness_step\":1.2,\"brightness_step_pct\":1.2,\"color_name\":\"aliceblue\",\"color_temp\":1.2,\"effect\":\"data\",\"flash\":\"long\",\"kelvin\":1.2,\"profile\":\"data\",\"transition\":1.2,\"white\":1.2}}",
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
