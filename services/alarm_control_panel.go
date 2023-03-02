package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewAlarmControlPanelAlarmArmAway creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_arm_away
// "Send the alarm the command for arm away."
func NewAlarmControlPanelAlarmArmAway(entities []string, code *string) *AlarmControlPanelAlarmArmAway {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_arm_away"
	a := &AlarmControlPanelAlarmArmAway{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Code *string `json:"code,omitempty"`
		}{Code: code},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AlarmControlPanelAlarmArmAway struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Code *string `json:"code,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AlarmControlPanelAlarmArmAway) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmArmAway) SetID(id *int) {
	a.Id = id
}

// NewAlarmControlPanelAlarmArmCustomBypass creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_arm_custom_bypass
// "Send arm custom bypass command."
func NewAlarmControlPanelAlarmArmCustomBypass(entities []string, code *string) *AlarmControlPanelAlarmArmCustomBypass {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_arm_custom_bypass"
	a := &AlarmControlPanelAlarmArmCustomBypass{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Code *string `json:"code,omitempty"`
		}{Code: code},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AlarmControlPanelAlarmArmCustomBypass struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Code *string `json:"code,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AlarmControlPanelAlarmArmCustomBypass) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmArmCustomBypass) SetID(id *int) {
	a.Id = id
}

// NewAlarmControlPanelAlarmArmHome creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_arm_home
// "Send the alarm the command for arm home."
func NewAlarmControlPanelAlarmArmHome(entities []string, code *string) *AlarmControlPanelAlarmArmHome {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_arm_home"
	a := &AlarmControlPanelAlarmArmHome{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Code *string `json:"code,omitempty"`
		}{Code: code},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AlarmControlPanelAlarmArmHome struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Code *string `json:"code,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AlarmControlPanelAlarmArmHome) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmArmHome) SetID(id *int) {
	a.Id = id
}

// NewAlarmControlPanelAlarmArmNight creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_arm_night
// "Send the alarm the command for arm night."
func NewAlarmControlPanelAlarmArmNight(entities []string, code *string) *AlarmControlPanelAlarmArmNight {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_arm_night"
	a := &AlarmControlPanelAlarmArmNight{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Code *string `json:"code,omitempty"`
		}{Code: code},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AlarmControlPanelAlarmArmNight struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Code *string `json:"code,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AlarmControlPanelAlarmArmNight) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmArmNight) SetID(id *int) {
	a.Id = id
}

// NewAlarmControlPanelAlarmArmVacation creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_arm_vacation
// "Send the alarm the command for arm vacation."
func NewAlarmControlPanelAlarmArmVacation(entities []string, code *string) *AlarmControlPanelAlarmArmVacation {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_arm_vacation"
	a := &AlarmControlPanelAlarmArmVacation{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Code *string `json:"code,omitempty"`
		}{Code: code},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AlarmControlPanelAlarmArmVacation struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Code *string `json:"code,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AlarmControlPanelAlarmArmVacation) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmArmVacation) SetID(id *int) {
	a.Id = id
}

// NewAlarmControlPanelAlarmDisarm creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_disarm
// "Send the alarm the command for disarm."
func NewAlarmControlPanelAlarmDisarm(entities []string, code *string) *AlarmControlPanelAlarmDisarm {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_disarm"
	a := &AlarmControlPanelAlarmDisarm{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Code *string `json:"code,omitempty"`
		}{Code: code},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AlarmControlPanelAlarmDisarm struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Code *string `json:"code,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AlarmControlPanelAlarmDisarm) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmDisarm) SetID(id *int) {
	a.Id = id
}

// NewAlarmControlPanelAlarmTrigger creates the object that can be sent to Home Assistant for domain alarm_control_panel, service alarm_trigger
// "Send the alarm the command for trigger."
func NewAlarmControlPanelAlarmTrigger(entities []string, code *string) *AlarmControlPanelAlarmTrigger {
	serviceDomain := "alarm_control_panel"
	serviceType := "call_service"
	serviceService := "alarm_trigger"
	a := &AlarmControlPanelAlarmTrigger{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Code *string `json:"code,omitempty"`
		}{Code: code},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AlarmControlPanelAlarmTrigger struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Code *string `json:"code,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AlarmControlPanelAlarmTrigger) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AlarmControlPanelAlarmTrigger) SetID(id *int) {
	a.Id = id
}
