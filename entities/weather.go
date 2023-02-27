package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Weather struct {
	Condition                *string   `json:"condition,omitempty"`
	Datetime                 *string   `json:"datetime,omitempty"`
	Forecast                 *[]string `json:"forecast,omitempty"`
	Humidity                 *float64  `json:"humidity,omitempty"`
	NativePrecipitation      *float64  `json:"native_precipitation,omitempty"`
	NativePrecipitationUnit  *string   `json:"native_precipitation_unit,omitempty"`
	NativePressure           *float64  `json:"native_pressure,omitempty"`
	NativePressureUnit       *string   `json:"native_pressure_unit,omitempty"`
	NativeTemperature        *float64  `json:"native_temperature,omitempty"`
	NativeTemperatureUnit    *string   `json:"native_temperature_unit,omitempty"`
	NativeTemplow            *float64  `json:"native_templow,omitempty"`
	NativeVisibility         *float64  `json:"native_visibility,omitempty"`
	NativeVisibilityUnit     *string   `json:"native_visibility_unit,omitempty"`
	NativeWindSpeed          *int      `json:"native_wind_speed,omitempty"`
	NativeWindSpeedUnit      *string   `json:"native_wind_speed_unit,omitempty"`
	Ozone                    *float64  `json:"ozone,omitempty"`
	PrecipitationProbability *int      `json:"precipitation_probability,omitempty"`
	WindBearing              *string   `json:"wind_bearing,omitempty"`
}
