package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewScheduleReload creates the object that can be sent to Home Assistant for domain schedule, service reload
// "Reload the schedule configuration"
func NewScheduleReload(target Target) *ScheduleReload {
	serviceDomain := "schedule"
	serviceType := "call_service"
	serviceService := "reload"
	s := &ScheduleReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
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
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ScheduleReload) SetID(id *int) {
	s.Id = id
}
