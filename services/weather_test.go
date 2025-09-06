package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestWeatherGetForecasts_JSON(t *testing.T) {
	getForecastsType := GetForecastsTypedaily

	tests := []struct {
		name   string
		fields *WeatherGetForecasts
		want   string
	}{{
		fields: NewWeatherGetForecasts(Targets("climate.kitchen")).GetForecastsType(getForecastsType),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"weather\",\"service\":\"get_forecasts\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"get_forecasts_type\":\"daily\"}}",
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
