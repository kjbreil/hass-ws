package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSchedulerAdd creates the object that can be sent to Home Assistant for domain scheduler, service add
// "Create a new schedule entity"
func NewSchedulerAdd(target Target) *SchedulerAdd {
	serviceDomain := "scheduler"
	serviceType := "call_service"
	serviceService := "add"
	s := &SchedulerAdd{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: SchedulerAddParams{},
	}
	return s
}

type SchedulerAdd struct {
	ServiceBase
	ServiceData SchedulerAddParams `json:"service_data,omitempty"`
}
type SchedulerAddParams struct {
	Name       *string     `json:"name,omitempty"`
	RepeatType *RepeatType `json:"repeat_type,omitempty"`
}

func (s *SchedulerAdd) SchedulerAddName(schedulerAddName string) *SchedulerAdd {
	s.ServiceData.Name = &schedulerAddName
	return s
}
func (s *SchedulerAdd) RepeatType(repeatType RepeatType) *SchedulerAdd {
	s.ServiceData.RepeatType = &repeatType
	return s
}
func (s *SchedulerAdd) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SchedulerAdd) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *SchedulerAdd) SetID(id *int) {
	s.Id = id
}

// NewSchedulerCopy creates the object that can be sent to Home Assistant for domain scheduler, service copy
// "Duplicate a schedule entity"
func NewSchedulerCopy(target Target) *SchedulerCopy {
	serviceDomain := "scheduler"
	serviceType := "call_service"
	serviceService := "copy"
	s := &SchedulerCopy{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: SchedulerCopyParams{},
	}
	return s
}

type SchedulerCopy struct {
	ServiceBase
	ServiceData SchedulerCopyParams `json:"service_data,omitempty"`
}
type SchedulerCopyParams struct {
	Name *string `json:"name,omitempty"`
}

func (s *SchedulerCopy) SchedulerCopyName(schedulerCopyName string) *SchedulerCopy {
	s.ServiceData.Name = &schedulerCopyName
	return s
}
func (s *SchedulerCopy) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SchedulerCopy) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *SchedulerCopy) SetID(id *int) {
	s.Id = id
}

// NewSchedulerEdit creates the object that can be sent to Home Assistant for domain scheduler, service edit
// "Edit a schedule entity"
func NewSchedulerEdit(target Target) *SchedulerEdit {
	serviceDomain := "scheduler"
	serviceType := "call_service"
	serviceService := "edit"
	s := &SchedulerEdit{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: SchedulerEditParams{},
	}
	return s
}

type SchedulerEdit struct {
	ServiceBase
	ServiceData SchedulerEditParams `json:"service_data,omitempty"`
}
type SchedulerEditParams struct {
	Name       *string     `json:"name,omitempty"`
	RepeatType *RepeatType `json:"repeat_type,omitempty"`
}

func (s *SchedulerEdit) SchedulerEditName(schedulerEditName string) *SchedulerEdit {
	s.ServiceData.Name = &schedulerEditName
	return s
}
func (s *SchedulerEdit) RepeatType(repeatType RepeatType) *SchedulerEdit {
	s.ServiceData.RepeatType = &repeatType
	return s
}
func (s *SchedulerEdit) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SchedulerEdit) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *SchedulerEdit) SetID(id *int) {
	s.Id = id
}

// NewSchedulerRemove creates the object that can be sent to Home Assistant for domain scheduler, service remove
// "Remove a schedule entity"
func NewSchedulerRemove(target Target) *SchedulerRemove {
	serviceDomain := "scheduler"
	serviceType := "call_service"
	serviceService := "remove"
	s := &SchedulerRemove{
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

type SchedulerRemove struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *SchedulerRemove) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SchedulerRemove) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *SchedulerRemove) SetID(id *int) {
	s.Id = id
}

// NewSchedulerRunAction creates the object that can be sent to Home Assistant for domain scheduler, service run_action
// "Execute the action of a schedule, optionally at a given time."
func NewSchedulerRunAction(target Target) *SchedulerRunAction {
	serviceDomain := "scheduler"
	serviceType := "call_service"
	serviceService := "run_action"
	s := &SchedulerRunAction{
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

type SchedulerRunAction struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *SchedulerRunAction) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SchedulerRunAction) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *SchedulerRunAction) SetID(id *int) {
	s.Id = id
}
