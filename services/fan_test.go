package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestFanDecreaseSpeed_JSON(t *testing.T) {
	percentageStep := 1.2

	tests := []struct {
		name   string
		fields *FanDecreaseSpeed
		want   string
	}{{
		fields: NewFanDecreaseSpeed(Targets("climate.kitchen"), &FanDecreaseSpeedParams{PercentageStep: &percentageStep}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"fan\",\"service\":\"decrease_speed\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"percentage_step\":1.2}}",
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
func TestFanIncreaseSpeed_JSON(t *testing.T) {
	percentageStep := 1.2

	tests := []struct {
		name   string
		fields *FanIncreaseSpeed
		want   string
	}{{
		fields: NewFanIncreaseSpeed(Targets("climate.kitchen"), &FanIncreaseSpeedParams{PercentageStep: &percentageStep}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"fan\",\"service\":\"increase_speed\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"percentage_step\":1.2}}",
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
func TestFanOscillate_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *FanOscillate
		want   string
	}{{
		fields: NewFanOscillate(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"fan\",\"service\":\"oscillate\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestFanSetDirection_JSON(t *testing.T) {
	direction := Directionforward

	tests := []struct {
		name   string
		fields *FanSetDirection
		want   string
	}{{
		fields: NewFanSetDirection(Targets("climate.kitchen"), &FanSetDirectionParams{Direction: &direction}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"fan\",\"service\":\"set_direction\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"direction\":\"forward\"}}",
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
func TestFanSetPercentage_JSON(t *testing.T) {
	percentage := 1.2

	tests := []struct {
		name   string
		fields *FanSetPercentage
		want   string
	}{{
		fields: NewFanSetPercentage(Targets("climate.kitchen"), &FanSetPercentageParams{Percentage: &percentage}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"fan\",\"service\":\"set_percentage\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"percentage\":1.2}}",
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
func TestFanSetPresetMode_JSON(t *testing.T) {
	presetMode := "data"

	tests := []struct {
		name   string
		fields *FanSetPresetMode
		want   string
	}{{
		fields: NewFanSetPresetMode(Targets("climate.kitchen"), &FanSetPresetModeParams{PresetMode: &presetMode}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"fan\",\"service\":\"set_preset_mode\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"preset_mode\":\"data\"}}",
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
func TestFanToggle_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *FanToggle
		want   string
	}{{
		fields: NewFanToggle(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"fan\",\"service\":\"toggle\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestFanTurnOff_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *FanTurnOff
		want   string
	}{{
		fields: NewFanTurnOff(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"fan\",\"service\":\"turn_off\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestFanTurnOn_JSON(t *testing.T) {
	percentage := 1.2
	presetMode := "data"
	speed := "data"

	tests := []struct {
		name   string
		fields *FanTurnOn
		want   string
	}{{
		fields: NewFanTurnOn(Targets("climate.kitchen"), &FanTurnOnParams{
			Percentage: &percentage,
			PresetMode: &presetMode,
			Speed:      &speed,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"fan\",\"service\":\"turn_on\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"percentage\":1.2,\"preset_mode\":\"data\",\"speed\":\"data\"}}",
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
