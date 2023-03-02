package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSchedulerAdd creates the object that can be sent to Home Assistant for domain scheduler, service add
// "Create a new schedule entity"
func NewSchedulerAdd(entities []string, name *string, repeatType *RepeatType) *SchedulerAdd {
	serviceDomain := "scheduler"
	serviceType := "call_service"
	serviceService := "add"
	s := &SchedulerAdd{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Name       *string     `json:"name,omitempty"`
			RepeatType *RepeatType `json:"repeat_type,omitempty"`
		}{
			Name:       name,
			RepeatType: repeatType,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SchedulerAdd struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Name       *string     `json:"name,omitempty"`
		RepeatType *RepeatType `json:"repeat_type,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SchedulerAdd) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SchedulerAdd) SetID(id *int) {
	s.Id = id
}

// NewSchedulerCopy creates the object that can be sent to Home Assistant for domain scheduler, service copy
// "Duplicate a schedule entity"
func NewSchedulerCopy(entities []string, name *string) *SchedulerCopy {
	serviceDomain := "scheduler"
	serviceType := "call_service"
	serviceService := "copy"
	s := &SchedulerCopy{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Name *string `json:"name,omitempty"`
		}{Name: name},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SchedulerCopy struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Name *string `json:"name,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SchedulerCopy) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SchedulerCopy) SetID(id *int) {
	s.Id = id
}

// NewSchedulerEdit creates the object that can be sent to Home Assistant for domain scheduler, service edit
// "Edit a schedule entity"
func NewSchedulerEdit(entities []string, name *string, repeatType *RepeatType) *SchedulerEdit {
	serviceDomain := "scheduler"
	serviceType := "call_service"
	serviceService := "edit"
	s := &SchedulerEdit{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Name       *string     `json:"name,omitempty"`
			RepeatType *RepeatType `json:"repeat_type,omitempty"`
		}{
			Name:       name,
			RepeatType: repeatType,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SchedulerEdit struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Name       *string     `json:"name,omitempty"`
		RepeatType *RepeatType `json:"repeat_type,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SchedulerEdit) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SchedulerEdit) SetID(id *int) {
	s.Id = id
}

// NewSchedulerRemove creates the object that can be sent to Home Assistant for domain scheduler, service remove
// "Remove a schedule entity"
func NewSchedulerRemove(entities []string) *SchedulerRemove {
	serviceDomain := "scheduler"
	serviceType := "call_service"
	serviceService := "remove"
	s := &SchedulerRemove{
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

type SchedulerRemove struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SchedulerRemove) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SchedulerRemove) SetID(id *int) {
	s.Id = id
}

// NewSchedulerRunAction creates the object that can be sent to Home Assistant for domain scheduler, service run_action
// "Execute the action of a schedule, optionally at a given time."
func NewSchedulerRunAction(entities []string) *SchedulerRunAction {
	serviceDomain := "scheduler"
	serviceType := "call_service"
	serviceService := "run_action"
	s := &SchedulerRunAction{
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

type SchedulerRunAction struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SchedulerRunAction) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SchedulerRunAction) SetID(id *int) {
	s.Id = id
}
