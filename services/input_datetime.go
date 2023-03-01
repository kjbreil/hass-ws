package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewInputDatetimeReload creates the object that can be sent to Home Assistant for domain input_datetime, service reload
// "Reload the input_datetime configuration."
func NewInputDatetimeReload(entities []string) *InputDatetimeReload {
	serviceDomain := "input_datetime"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputDatetimeReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return i
}

type InputDatetimeReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputDatetimeReload) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputDatetimeReload) SetID(id *int) {
	i.Id = id
}

// NewInputDatetimeSetDatetime creates the object that can be sent to Home Assistant for domain input_datetime, service set_datetime
// "This can be used to dynamically set the date and/or time."
func NewInputDatetimeSetDatetime(entities []string, date *string, datetime *string, timestamp *int) *InputDatetimeSetDatetime {
	serviceDomain := "input_datetime"
	serviceType := "call_service"
	serviceService := "set_datetime"
	i := &InputDatetimeSetDatetime{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Date      *string `json:"date,omitempty"`
			Datetime  *string `json:"datetime,omitempty"`
			Timestamp *int    `json:"timestamp,omitempty"`
		}{
			Date:      date,
			Datetime:  datetime,
			Timestamp: timestamp,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return i
}

type InputDatetimeSetDatetime struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Date      *string `json:"date,omitempty"`
		Datetime  *string `json:"datetime,omitempty"`
		Timestamp *int    `json:"timestamp,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputDatetimeSetDatetime) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputDatetimeSetDatetime) SetID(id *int) {
	i.Id = id
}
