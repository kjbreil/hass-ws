package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewScheduleReload creates the object that can be sent to Home Assistant for domain schedule, service reload
// "Reload the schedule configuration"
func NewScheduleReload(entities []string) *ScheduleReload {
	serviceDomain := "schedule"
	serviceType := "call_service"
	serviceService := "reload"
	s := &ScheduleReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type ScheduleReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *ScheduleReload) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ScheduleReload) SetID(id *int) {
	s.Id = id
}
