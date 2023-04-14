package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewInputDatetimeReload creates the object that can be sent to Home Assistant for domain input_datetime, service reload
// "Reload the input_datetime configuration."
func NewInputDatetimeReload(target Target) *InputDatetimeReload {
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
		ServiceData: nil,
	}
	return i
}

type InputDatetimeReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputDatetimeReload) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputDatetimeReload) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputDatetimeReload) SetID(id *int) {
	i.Id = id
}

// NewInputDatetimeSetDatetime creates the object that can be sent to Home Assistant for domain input_datetime, service set_datetime
// "This can be used to dynamically set the date and/or time."
func NewInputDatetimeSetDatetime(target Target) *InputDatetimeSetDatetime {
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
		ServiceData: InputDatetimeSetDatetimeParams{},
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

func (i *InputDatetimeSetDatetime) Date(date string) *InputDatetimeSetDatetime {
	i.ServiceData.Date = &date
	return i
}
func (i *InputDatetimeSetDatetime) Datetime(datetime string) *InputDatetimeSetDatetime {
	i.ServiceData.Datetime = &datetime
	return i
}
func (i *InputDatetimeSetDatetime) Timestamp(timestamp float64) *InputDatetimeSetDatetime {
	i.ServiceData.Timestamp = &timestamp
	return i
}
func (i *InputDatetimeSetDatetime) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputDatetimeSetDatetime) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
func (i *InputDatetimeSetDatetime) SetID(id *int) {
	i.Id = id
}
