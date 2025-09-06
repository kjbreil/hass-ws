package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewScheduleGetSchedule creates the object that can be sent to Home Assistant for domain schedule, service get_schedule
// "Retrieves the configured time ranges of one or multiple schedules."
func NewScheduleGetSchedule(target Target) *ScheduleGetSchedule {
	serviceDomain := "schedule"
	serviceType := "call_service"
	serviceService := "get_schedule"
	s := &ScheduleGetSchedule{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: true,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return s
}

type ScheduleGetSchedule struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *ScheduleGetSchedule) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *ScheduleGetSchedule) Targets() []string {
	return s.Target.EntityId
}
func (s *ScheduleGetSchedule) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}

// NewScheduleReload creates the object that can be sent to Home Assistant for domain schedule, service reload
// "Reloads schedules from the YAML-configuration."
func NewScheduleReload(target Target) *ScheduleReload {
	serviceDomain := "schedule"
	serviceType := "call_service"
	serviceService := "reload"
	s := &ScheduleReload{
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
	return s
}

type ScheduleReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *ScheduleReload) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *ScheduleReload) Targets() []string {
	return s.Target.EntityId
}
func (s *ScheduleReload) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
