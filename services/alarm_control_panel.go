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
func NewAlarmControlPanelAlarmArmAway(target Target) *AlarmControlPanelAlarmArmAway {
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
		ServiceData: AlarmControlPanelAlarmArmAwayParams{},
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

func (a *AlarmControlPanelAlarmArmAway) Code(code string) *AlarmControlPanelAlarmArmAway {
	a.ServiceData.Code = &code
	return a
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
func NewAlarmControlPanelAlarmArmCustomBypass(target Target) *AlarmControlPanelAlarmArmCustomBypass {
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
		ServiceData: AlarmControlPanelAlarmArmCustomBypassParams{},
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

func (a *AlarmControlPanelAlarmArmCustomBypass) Code(code string) *AlarmControlPanelAlarmArmCustomBypass {
	a.ServiceData.Code = &code
	return a
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
func NewAlarmControlPanelAlarmArmHome(target Target) *AlarmControlPanelAlarmArmHome {
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
		ServiceData: AlarmControlPanelAlarmArmHomeParams{},
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

func (a *AlarmControlPanelAlarmArmHome) Code(code string) *AlarmControlPanelAlarmArmHome {
	a.ServiceData.Code = &code
	return a
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
func NewAlarmControlPanelAlarmArmNight(target Target) *AlarmControlPanelAlarmArmNight {
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
		ServiceData: AlarmControlPanelAlarmArmNightParams{},
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

func (a *AlarmControlPanelAlarmArmNight) Code(code string) *AlarmControlPanelAlarmArmNight {
	a.ServiceData.Code = &code
	return a
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
func NewAlarmControlPanelAlarmArmVacation(target Target) *AlarmControlPanelAlarmArmVacation {
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
		ServiceData: AlarmControlPanelAlarmArmVacationParams{},
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

func (a *AlarmControlPanelAlarmArmVacation) Code(code string) *AlarmControlPanelAlarmArmVacation {
	a.ServiceData.Code = &code
	return a
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
func NewAlarmControlPanelAlarmDisarm(target Target) *AlarmControlPanelAlarmDisarm {
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
		ServiceData: AlarmControlPanelAlarmDisarmParams{},
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

func (a *AlarmControlPanelAlarmDisarm) Code(code string) *AlarmControlPanelAlarmDisarm {
	a.ServiceData.Code = &code
	return a
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
func NewAlarmControlPanelAlarmTrigger(target Target) *AlarmControlPanelAlarmTrigger {
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
		ServiceData: AlarmControlPanelAlarmTriggerParams{},
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

func (a *AlarmControlPanelAlarmTrigger) Code(code string) *AlarmControlPanelAlarmTrigger {
	a.ServiceData.Code = &code
	return a
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
