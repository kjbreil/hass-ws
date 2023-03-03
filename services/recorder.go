package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewRecorderDisable creates the object that can be sent to Home Assistant for domain recorder, service disable
// "Stop the recording of events and state changes"
func NewRecorderDisable(target Target, recorderDisableParams RecorderDisableParams) *RecorderDisable {
	serviceDomain := "recorder"
	serviceType := "call_service"
	serviceService := "disable"
	r := &RecorderDisable{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: recorderDisableParams,
	}
	return r
}

type RecorderDisable struct {
	ServiceBase
	ServiceData RecorderDisableParams `json:"service_data,omitempty"`
}
type RecorderDisableParams struct{}

func (r *RecorderDisable) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RecorderDisable) SetID(id *int) {
	r.Id = id
}

// NewRecorderEnable creates the object that can be sent to Home Assistant for domain recorder, service enable
// "Start the recording of events and state changes"
func NewRecorderEnable(target Target, recorderEnableParams RecorderEnableParams) *RecorderEnable {
	serviceDomain := "recorder"
	serviceType := "call_service"
	serviceService := "enable"
	r := &RecorderEnable{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: recorderEnableParams,
	}
	return r
}

type RecorderEnable struct {
	ServiceBase
	ServiceData RecorderEnableParams `json:"service_data,omitempty"`
}
type RecorderEnableParams struct{}

func (r *RecorderEnable) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RecorderEnable) SetID(id *int) {
	r.Id = id
}

// NewRecorderPurge creates the object that can be sent to Home Assistant for domain recorder, service purge
// "Start purge task - to clean up old data from your database."
func NewRecorderPurge(target Target, recorderPurgeParams RecorderPurgeParams) *RecorderPurge {
	serviceDomain := "recorder"
	serviceType := "call_service"
	serviceService := "purge"
	r := &RecorderPurge{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: recorderPurgeParams,
	}
	return r
}

type RecorderPurge struct {
	ServiceBase
	ServiceData RecorderPurgeParams `json:"service_data,omitempty"`
}
type RecorderPurgeParams struct {
	KeepDays *float64 `json:"keep_days,omitempty"`
}

func (r *RecorderPurge) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RecorderPurge) SetID(id *int) {
	r.Id = id
}

// NewRecorderPurgeEntities creates the object that can be sent to Home Assistant for domain recorder, service purge_entities
// "Start purge task to remove specific entities from your database."
func NewRecorderPurgeEntities(target Target, recorderPurgeEntitiesParams RecorderPurgeEntitiesParams) *RecorderPurgeEntities {
	serviceDomain := "recorder"
	serviceType := "call_service"
	serviceService := "purge_entities"
	r := &RecorderPurgeEntities{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: recorderPurgeEntitiesParams,
	}
	return r
}

type RecorderPurgeEntities struct {
	ServiceBase
	ServiceData RecorderPurgeEntitiesParams `json:"service_data,omitempty"`
}
type RecorderPurgeEntitiesParams struct{}

func (r *RecorderPurgeEntities) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RecorderPurgeEntities) SetID(id *int) {
	r.Id = id
}
