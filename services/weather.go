package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewWeatherGetForecasts creates the object that can be sent to Home Assistant for domain weather, service get_forecasts
// "Retrieves the forecast from selected weather services."
func NewWeatherGetForecasts(target Target) *WeatherGetForecasts {
	serviceDomain := "weather"
	serviceType := "call_service"
	serviceService := "get_forecasts"
	w := &WeatherGetForecasts{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: true,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: WeatherGetForecastsParams{},
	}
	return w
}

type WeatherGetForecasts struct {
	ServiceBase
	ServiceData WeatherGetForecastsParams `json:"service_data,omitempty"`
}
type WeatherGetForecastsParams struct {
	GetForecastsType *GetForecastsType `json:"type,omitempty"`
}

func (w *WeatherGetForecasts) GetForecastsType(getForecastsType GetForecastsType) *WeatherGetForecasts {
	w.ServiceData.GetForecastsType = &getForecastsType
	return w
}
func (w *WeatherGetForecasts) JSON() string {
	data, _ := gojson.Marshal(w)
	return string(data)
}
func (w *WeatherGetForecasts) Targets() []string {
	return w.Target.EntityId
}
func (w *WeatherGetForecasts) Name() string {
	return fmt.Sprintf("%s.%s", *w.Domain, *w.Service)
}
