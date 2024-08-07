package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewRecorderDisable creates the object that can be sent to Home Assistant for domain recorder, service disable
// "Stop the recording of events and state changes"
func NewRecorderDisable(target Target) *RecorderDisable {
	serviceDomain := "recorder"
	serviceType := "call_service"
	serviceService := "disable"
	r := &RecorderDisable{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return r
}

type RecorderDisable struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (r *RecorderDisable) JSON() string {
	data, _ := gojson.Marshal(r)
	return string(data)
}
func (r *RecorderDisable) Targets() []string {
	return r.Target.EntityId
}
func (r *RecorderDisable) Name() string {
	return fmt.Sprintf("%s.%s", *r.Domain, *r.Service)
}

// NewRecorderEnable creates the object that can be sent to Home Assistant for domain recorder, service enable
// "Start the recording of events and state changes"
func NewRecorderEnable(target Target) *RecorderEnable {
	serviceDomain := "recorder"
	serviceType := "call_service"
	serviceService := "enable"
	r := &RecorderEnable{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return r
}

type RecorderEnable struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (r *RecorderEnable) JSON() string {
	data, _ := gojson.Marshal(r)
	return string(data)
}
func (r *RecorderEnable) Targets() []string {
	return r.Target.EntityId
}
func (r *RecorderEnable) Name() string {
	return fmt.Sprintf("%s.%s", *r.Domain, *r.Service)
}

// NewRecorderPurge creates the object that can be sent to Home Assistant for domain recorder, service purge
// "Start purge task - to clean up old data from your database."
func NewRecorderPurge(target Target) *RecorderPurge {
	serviceDomain := "recorder"
	serviceType := "call_service"
	serviceService := "purge"
	r := &RecorderPurge{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: RecorderPurgeParams{},
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

func (r *RecorderPurge) KeepDays(keepDays float64) *RecorderPurge {
	r.ServiceData.KeepDays = &keepDays
	return r
}
func (r *RecorderPurge) JSON() string {
	data, _ := gojson.Marshal(r)
	return string(data)
}
func (r *RecorderPurge) Targets() []string {
	return r.Target.EntityId
}
func (r *RecorderPurge) Name() string {
	return fmt.Sprintf("%s.%s", *r.Domain, *r.Service)
}

// NewRecorderPurgeEntities creates the object that can be sent to Home Assistant for domain recorder, service purge_entities
// "Start purge task to remove specific entities from your database."
func NewRecorderPurgeEntities(target Target) *RecorderPurgeEntities {
	serviceDomain := "recorder"
	serviceType := "call_service"
	serviceService := "purge_entities"
	r := &RecorderPurgeEntities{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return r
}

type RecorderPurgeEntities struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (r *RecorderPurgeEntities) JSON() string {
	data, _ := gojson.Marshal(r)
	return string(data)
}
func (r *RecorderPurgeEntities) Targets() []string {
	return r.Target.EntityId
}
func (r *RecorderPurgeEntities) Name() string {
	return fmt.Sprintf("%s.%s", *r.Domain, *r.Service)
}
