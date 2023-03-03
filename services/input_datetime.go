package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewInputDatetimeReload creates the object that can be sent to Home Assistant for domain input_datetime, service reload
// "Reload the input_datetime configuration."
func NewInputDatetimeReload(target Target, inputDatetimeReloadParams InputDatetimeReloadParams) *InputDatetimeReload {
	serviceDomain := "input_datetime"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputDatetimeReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: inputDatetimeReloadParams,
	}
	return i
}

type InputDatetimeReload struct {
	ServiceBase
	ServiceData InputDatetimeReloadParams `json:"service_data,omitempty"`
}
type InputDatetimeReloadParams struct{}

func (i *InputDatetimeReload) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputDatetimeReload) SetID(id *int) {
	i.Id = id
}

// NewInputDatetimeSetDatetime creates the object that can be sent to Home Assistant for domain input_datetime, service set_datetime
// "This can be used to dynamically set the date and/or time."
func NewInputDatetimeSetDatetime(target Target, inputDatetimeSetDatetimeParams InputDatetimeSetDatetimeParams) *InputDatetimeSetDatetime {
	serviceDomain := "input_datetime"
	serviceType := "call_service"
	serviceService := "set_datetime"
	i := &InputDatetimeSetDatetime{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: inputDatetimeSetDatetimeParams,
	}
	return i
}

type InputDatetimeSetDatetime struct {
	ServiceBase
	ServiceData InputDatetimeSetDatetimeParams `json:"service_data,omitempty"`
}
type InputDatetimeSetDatetimeParams struct {
	Date      *string  `json:"date,omitempty"`
	Datetime  *string  `json:"datetime,omitempty"`
	Timestamp *float64 `json:"timestamp,omitempty"`
}

func (i *InputDatetimeSetDatetime) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputDatetimeSetDatetime) SetID(id *int) {
	i.Id = id
}
