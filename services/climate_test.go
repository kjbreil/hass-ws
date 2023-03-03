package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestClimateSetAuxHeat_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *ClimateSetAuxHeat
		want   string
	}{{
		fields: NewClimateSetAuxHeat(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"climate\",\"service\":\"set_aux_heat\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestClimateSetFanMode_JSON(t *testing.T) {
	fanMode := "data"

	tests := []struct {
		name   string
		fields *ClimateSetFanMode
		want   string
	}{{
		fields: NewClimateSetFanMode(Targets("climate.kitchen"), &ClimateSetFanModeParams{FanMode: &fanMode}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"climate\",\"service\":\"set_fan_mode\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"fan_mode\":\"data\"}}",
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
func TestClimateSetHumidity_JSON(t *testing.T) {
	humidity := 1.2

	tests := []struct {
		name   string
		fields *ClimateSetHumidity
		want   string
	}{{
		fields: NewClimateSetHumidity(Targets("climate.kitchen"), &ClimateSetHumidityParams{Humidity: &humidity}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"climate\",\"service\":\"set_humidity\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"humidity\":1.2}}",
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
func TestClimateSetHvacMode_JSON(t *testing.T) {
	hvacMode := HvacModeauto

	tests := []struct {
		name   string
		fields *ClimateSetHvacMode
		want   string
	}{{
		fields: NewClimateSetHvacMode(Targets("climate.kitchen"), &ClimateSetHvacModeParams{HvacMode: &hvacMode}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"climate\",\"service\":\"set_hvac_mode\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"hvac_mode\":\"auto\"}}",
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
func TestClimateSetPresetMode_JSON(t *testing.T) {
	presetMode := "data"

	tests := []struct {
		name   string
		fields *ClimateSetPresetMode
		want   string
	}{{
		fields: NewClimateSetPresetMode(Targets("climate.kitchen"), &ClimateSetPresetModeParams{PresetMode: &presetMode}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"climate\",\"service\":\"set_preset_mode\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"preset_mode\":\"data\"}}",
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
func TestClimateSetSwingMode_JSON(t *testing.T) {
	swingMode := "data"

	tests := []struct {
		name   string
		fields *ClimateSetSwingMode
		want   string
	}{{
		fields: NewClimateSetSwingMode(Targets("climate.kitchen"), &ClimateSetSwingModeParams{SwingMode: &swingMode}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"climate\",\"service\":\"set_swing_mode\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"swing_mode\":\"data\"}}",
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
func TestClimateSetTemperature_JSON(t *testing.T) {
	hvacMode := HvacModeauto
	targetTempHigh := 1.2
	targetTempLow := 1.2
	temperature := 1.2

	tests := []struct {
		name   string
		fields *ClimateSetTemperature
		want   string
	}{{
		fields: NewClimateSetTemperature(Targets("climate.kitchen"), &ClimateSetTemperatureParams{
			HvacMode:       &hvacMode,
			TargetTempHigh: &targetTempHigh,
			TargetTempLow:  &targetTempLow,
			Temperature:    &temperature,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"climate\",\"service\":\"set_temperature\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"hvac_mode\":\"auto\",\"target_temp_high\":1.2,\"target_temp_low\":1.2,\"temperature\":1.2}}",
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
func TestClimateTurnOff_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *ClimateTurnOff
		want   string
	}{{
		fields: NewClimateTurnOff(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"climate\",\"service\":\"turn_off\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestClimateTurnOn_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *ClimateTurnOn
		want   string
	}{{
		fields: NewClimateTurnOn(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"climate\",\"service\":\"turn_on\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
