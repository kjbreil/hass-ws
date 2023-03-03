package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestInputDatetimeReload_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *InputDatetimeReload
		want   string
	}{{
		fields: NewInputDatetimeReload(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_datetime\",\"service\":\"reload\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestInputDatetimeSetDatetime_JSON(t *testing.T) {
	date := "data"
	datetime := "data"
	timestamp := 1.2

	tests := []struct {
		name   string
		fields *InputDatetimeSetDatetime
		want   string
	}{{
		fields: NewInputDatetimeSetDatetime(Targets("climate.kitchen"), &InputDatetimeSetDatetimeParams{
			Date:      &date,
			Datetime:  &datetime,
			Timestamp: &timestamp,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"input_datetime\",\"service\":\"set_datetime\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"date\":\"data\",\"datetime\":\"data\",\"timestamp\":1.2}}",
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
