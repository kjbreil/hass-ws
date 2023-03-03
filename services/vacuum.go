package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewVacuumCleanSpot creates the object that can be sent to Home Assistant for domain vacuum, service clean_spot
// "Tell the vacuum cleaner to do a spot clean-up."
func NewVacuumCleanSpot(target Target) *VacuumCleanSpot {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "clean_spot"
	v := &VacuumCleanSpot{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return v
}

type VacuumCleanSpot struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (v *VacuumCleanSpot) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumCleanSpot) SetID(id *int) {
	v.Id = id
}

// NewVacuumLocate creates the object that can be sent to Home Assistant for domain vacuum, service locate
// "Locate the vacuum cleaner robot."
func NewVacuumLocate(target Target) *VacuumLocate {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "locate"
	v := &VacuumLocate{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return v
}

type VacuumLocate struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (v *VacuumLocate) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumLocate) SetID(id *int) {
	v.Id = id
}

// NewVacuumPause creates the object that can be sent to Home Assistant for domain vacuum, service pause
// "Pause the cleaning task."
func NewVacuumPause(target Target) *VacuumPause {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "pause"
	v := &VacuumPause{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return v
}

type VacuumPause struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (v *VacuumPause) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumPause) SetID(id *int) {
	v.Id = id
}

// NewVacuumReturnToBase creates the object that can be sent to Home Assistant for domain vacuum, service return_to_base
// "Tell the vacuum cleaner to return to its dock."
func NewVacuumReturnToBase(target Target) *VacuumReturnToBase {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "return_to_base"
	v := &VacuumReturnToBase{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return v
}

type VacuumReturnToBase struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (v *VacuumReturnToBase) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumReturnToBase) SetID(id *int) {
	v.Id = id
}

// NewVacuumSendCommand creates the object that can be sent to Home Assistant for domain vacuum, service send_command
// "Send a raw command to the vacuum cleaner."
func NewVacuumSendCommand(target Target, vacuumSendCommandParams *VacuumSendCommandParams) *VacuumSendCommand {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "send_command"
	v := &VacuumSendCommand{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *vacuumSendCommandParams,
	}
	return v
}

type VacuumSendCommand struct {
	ServiceBase
	ServiceData VacuumSendCommandParams `json:"service_data,omitempty"`
}
type VacuumSendCommandParams struct {
	Command *string `json:"command,omitempty"`
}

func (v *VacuumSendCommand) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumSendCommand) SetID(id *int) {
	v.Id = id
}

// NewVacuumSetFanSpeed creates the object that can be sent to Home Assistant for domain vacuum, service set_fan_speed
// "Set the fan speed of the vacuum cleaner."
func NewVacuumSetFanSpeed(target Target, vacuumSetFanSpeedParams *VacuumSetFanSpeedParams) *VacuumSetFanSpeed {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "set_fan_speed"
	v := &VacuumSetFanSpeed{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *vacuumSetFanSpeedParams,
	}
	return v
}

type VacuumSetFanSpeed struct {
	ServiceBase
	ServiceData VacuumSetFanSpeedParams `json:"service_data,omitempty"`
}
type VacuumSetFanSpeedParams struct {
	FanSpeed *string `json:"fan_speed,omitempty"`
}

func (v *VacuumSetFanSpeed) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumSetFanSpeed) SetID(id *int) {
	v.Id = id
}

// NewVacuumStart creates the object that can be sent to Home Assistant for domain vacuum, service start
// "Start or resume the cleaning task."
func NewVacuumStart(target Target) *VacuumStart {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "start"
	v := &VacuumStart{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return v
}

type VacuumStart struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (v *VacuumStart) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumStart) SetID(id *int) {
	v.Id = id
}

// NewVacuumStartPause creates the object that can be sent to Home Assistant for domain vacuum, service start_pause
// "Start, pause, or resume the cleaning task."
func NewVacuumStartPause(target Target) *VacuumStartPause {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "start_pause"
	v := &VacuumStartPause{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return v
}

type VacuumStartPause struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (v *VacuumStartPause) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumStartPause) SetID(id *int) {
	v.Id = id
}

// NewVacuumStop creates the object that can be sent to Home Assistant for domain vacuum, service stop
// "Stop the current cleaning task."
func NewVacuumStop(target Target) *VacuumStop {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "stop"
	v := &VacuumStop{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return v
}

type VacuumStop struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (v *VacuumStop) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumStop) SetID(id *int) {
	v.Id = id
}

// NewVacuumToggle creates the object that can be sent to Home Assistant for domain vacuum, service toggle
// ""
func NewVacuumToggle(target Target) *VacuumToggle {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "toggle"
	v := &VacuumToggle{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return v
}

type VacuumToggle struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (v *VacuumToggle) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumToggle) SetID(id *int) {
	v.Id = id
}

// NewVacuumTurnOff creates the object that can be sent to Home Assistant for domain vacuum, service turn_off
// "Stop the current cleaning task and return to home."
func NewVacuumTurnOff(target Target) *VacuumTurnOff {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "turn_off"
	v := &VacuumTurnOff{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return v
}

type VacuumTurnOff struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (v *VacuumTurnOff) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumTurnOff) SetID(id *int) {
	v.Id = id
}

// NewVacuumTurnOn creates the object that can be sent to Home Assistant for domain vacuum, service turn_on
// "Start a new cleaning task."
func NewVacuumTurnOn(target Target) *VacuumTurnOn {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "turn_on"
	v := &VacuumTurnOn{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return v
}

type VacuumTurnOn struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (v *VacuumTurnOn) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumTurnOn) SetID(id *int) {
	v.Id = id
}
