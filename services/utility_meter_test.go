package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestUtilityMeterCalibrate_JSON(t *testing.T) {
	value := "data"

	tests := []struct {
		name   string
		fields *UtilityMeterCalibrate
		want   string
	}{{
		fields: NewUtilityMeterCalibrate(Targets("climate.kitchen")).Value(value),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"utility_meter\",\"service\":\"calibrate\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"value\":\"data\"}}",
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
func TestUtilityMeterReset_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *UtilityMeterReset
		want   string
	}{{
		fields: NewUtilityMeterReset(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"utility_meter\",\"service\":\"reset\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
