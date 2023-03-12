package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewAlarmControlPanelAlarmArmAway creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_arm_away
// "Send the alarm the command for arm away."
func NewAlarmControlPanelAlarmArmAway(target Target, alarmControlPanelAlarmArmAwayParams *AlarmControlPanelAlarmArmAwayParams) *AlarmControlPanelAlarmArmAway {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_arm_away"
	a := &AlarmControlPanelAlarmArmAway{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *alarmControlPanelAlarmArmAwayParams,
	}
	return a
}

type AlarmControlPanelAlarmArmAway struct {
	ServiceBase
	ServiceData AlarmControlPanelAlarmArmAwayParams `json:"service_data,omitempty"`
}
type AlarmControlPanelAlarmArmAwayParams struct {
	Code *string `json:"code,omitempty"`
}

func (a *AlarmControlPanelAlarmArmAway) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmArmAway) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}
func (a *AlarmControlPanelAlarmArmAway) SetID(id *int) {
	a.Id = id
}

// NewAlarmControlPanelAlarmArmCustomBypass creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_arm_custom_bypass
// "Send arm custom bypass command."
func NewAlarmControlPanelAlarmArmCustomBypass(target Target, alarmControlPanelAlarmArmCustomBypassParams *AlarmControlPanelAlarmArmCustomBypassParams) *AlarmControlPanelAlarmArmCustomBypass {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_arm_custom_bypass"
	a := &AlarmControlPanelAlarmArmCustomBypass{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *alarmControlPanelAlarmArmCustomBypassParams,
	}
	return a
}

type AlarmControlPanelAlarmArmCustomBypass struct {
	ServiceBase
	ServiceData AlarmControlPanelAlarmArmCustomBypassParams `json:"service_data,omitempty"`
}
type AlarmControlPanelAlarmArmCustomBypassParams struct {
	Code *string `json:"code,omitempty"`
}

func (a *AlarmControlPanelAlarmArmCustomBypass) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmArmCustomBypass) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}
func (a *AlarmControlPanelAlarmArmCustomBypass) SetID(id *int) {
	a.Id = id
}

// NewAlarmControlPanelAlarmArmHome creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_arm_home
// "Send the alarm the command for arm home."
func NewAlarmControlPanelAlarmArmHome(target Target, alarmControlPanelAlarmArmHomeParams *AlarmControlPanelAlarmArmHomeParams) *AlarmControlPanelAlarmArmHome {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_arm_home"
	a := &AlarmControlPanelAlarmArmHome{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *alarmControlPanelAlarmArmHomeParams,
	}
	return a
}

type AlarmControlPanelAlarmArmHome struct {
	ServiceBase
	ServiceData AlarmControlPanelAlarmArmHomeParams `json:"service_data,omitempty"`
}
type AlarmControlPanelAlarmArmHomeParams struct {
	Code *string `json:"code,omitempty"`
}

func (a *AlarmControlPanelAlarmArmHome) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmArmHome) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}
func (a *AlarmControlPanelAlarmArmHome) SetID(id *int) {
	a.Id = id
}

// NewAlarmControlPanelAlarmArmNight creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_arm_night
// "Send the alarm the command for arm night."
func NewAlarmControlPanelAlarmArmNight(target Target, alarmControlPanelAlarmArmNightParams *AlarmControlPanelAlarmArmNightParams) *AlarmControlPanelAlarmArmNight {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_arm_night"
	a := &AlarmControlPanelAlarmArmNight{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *alarmControlPanelAlarmArmNightParams,
	}
	return a
}

type AlarmControlPanelAlarmArmNight struct {
	ServiceBase
	ServiceData AlarmControlPanelAlarmArmNightParams `json:"service_data,omitempty"`
}
type AlarmControlPanelAlarmArmNightParams struct {
	Code *string `json:"code,omitempty"`
}

func (a *AlarmControlPanelAlarmArmNight) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmArmNight) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}
func (a *AlarmControlPanelAlarmArmNight) SetID(id *int) {
	a.Id = id
}

// NewAlarmControlPanelAlarmArmVacation creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_arm_vacation
// "Send the alarm the command for arm vacation."
func NewAlarmControlPanelAlarmArmVacation(target Target, alarmControlPanelAlarmArmVacationParams *AlarmControlPanelAlarmArmVacationParams) *AlarmControlPanelAlarmArmVacation {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_arm_vacation"
	a := &AlarmControlPanelAlarmArmVacation{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *alarmControlPanelAlarmArmVacationParams,
	}
	return a
}

type AlarmControlPanelAlarmArmVacation struct {
	ServiceBase
	ServiceData AlarmControlPanelAlarmArmVacationParams `json:"service_data,omitempty"`
}
type AlarmControlPanelAlarmArmVacationParams struct {
	Code *string `json:"code,omitempty"`
}

func (a *AlarmControlPanelAlarmArmVacation) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmArmVacation) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}
func (a *AlarmControlPanelAlarmArmVacation) SetID(id *int) {
	a.Id = id
}

// NewAlarmControlPanelAlarmDisarm creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_disarm
// "Send the alarm the command for disarm."
func NewAlarmControlPanelAlarmDisarm(target Target, alarmControlPanelAlarmDisarmParams *AlarmControlPanelAlarmDisarmParams) *AlarmControlPanelAlarmDisarm {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_disarm"
	a := &AlarmControlPanelAlarmDisarm{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *alarmControlPanelAlarmDisarmParams,
	}
	return a
}

type AlarmControlPanelAlarmDisarm struct {
	ServiceBase
	ServiceData AlarmControlPanelAlarmDisarmParams `json:"service_data,omitempty"`
}
type AlarmControlPanelAlarmDisarmParams struct {
	Code *string `json:"code,omitempty"`
}

func (a *AlarmControlPanelAlarmDisarm) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmDisarm) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}
func (a *AlarmControlPanelAlarmDisarm) SetID(id *int) {
	a.Id = id
}

// NewAlarmControlPanelAlarmTrigger creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_trigger
// "Send the alarm the command for trigger."
func NewAlarmControlPanelAlarmTrigger(target Target, alarmControlPanelAlarmTriggerParams *AlarmControlPanelAlarmTriggerParams) *AlarmControlPanelAlarmTrigger {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_trigger"
	a := &AlarmControlPanelAlarmTrigger{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *alarmControlPanelAlarmTriggerParams,
	}
	return a
}

type AlarmControlPanelAlarmTrigger struct {
	ServiceBase
	ServiceData AlarmControlPanelAlarmTriggerParams `json:"service_data,omitempty"`
}
type AlarmControlPanelAlarmTriggerParams struct {
	Code *string `json:"code,omitempty"`
}

func (a *AlarmControlPanelAlarmTrigger) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmTrigger) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}
func (a *AlarmControlPanelAlarmTrigger) SetID(id *int) {
	a.Id = id
}
