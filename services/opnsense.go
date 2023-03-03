package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewOpnsenseCloseNotice creates the object that can be sent to Home Assistant for domain opnsense, service close_notice
// "Closes a notice(s)."
func NewOpnsenseCloseNotice(target Target, opnsenseCloseNoticeParams OpnsenseCloseNoticeParams) *OpnsenseCloseNotice {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "close_notice"
	o := &OpnsenseCloseNotice{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: opnsenseCloseNoticeParams,
	}
	return o
}

type OpnsenseCloseNotice struct {
	ServiceBase
	ServiceData OpnsenseCloseNoticeParams `json:"service_data,omitempty"`
}
type OpnsenseCloseNoticeParams struct{}

func (o *OpnsenseCloseNotice) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseCloseNotice) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseFileNotice creates the object that can be sent to Home Assistant for domain opnsense, service file_notice
// "Files a notice(s)."
func NewOpnsenseFileNotice(target Target, opnsenseFileNoticeParams OpnsenseFileNoticeParams) *OpnsenseFileNotice {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "file_notice"
	o := &OpnsenseFileNotice{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: opnsenseFileNoticeParams,
	}
	return o
}

type OpnsenseFileNotice struct {
	ServiceBase
	ServiceData OpnsenseFileNoticeParams `json:"service_data,omitempty"`
}
type OpnsenseFileNoticeParams struct{}

func (o *OpnsenseFileNotice) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseFileNotice) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseRestartService creates the object that can be sent to Home Assistant for domain opnsense, service restart_service
// "Restarts a service."
func NewOpnsenseRestartService(target Target, opnsenseRestartServiceParams OpnsenseRestartServiceParams) *OpnsenseRestartService {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "restart_service"
	o := &OpnsenseRestartService{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: opnsenseRestartServiceParams,
	}
	return o
}

type OpnsenseRestartService struct {
	ServiceBase
	ServiceData OpnsenseRestartServiceParams `json:"service_data,omitempty"`
}
type OpnsenseRestartServiceParams struct{}

func (o *OpnsenseRestartService) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseRestartService) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseSendWol creates the object that can be sent to Home Assistant for domain opnsense, service send_wol
// "Sends wake-on-lan magic packet."
func NewOpnsenseSendWol(target Target, opnsenseSendWolParams OpnsenseSendWolParams) *OpnsenseSendWol {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "send_wol"
	o := &OpnsenseSendWol{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: opnsenseSendWolParams,
	}
	return o
}

type OpnsenseSendWol struct {
	ServiceBase
	ServiceData OpnsenseSendWolParams `json:"service_data,omitempty"`
}
type OpnsenseSendWolParams struct{}

func (o *OpnsenseSendWol) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseSendWol) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseStartService creates the object that can be sent to Home Assistant for domain opnsense, service start_service
// "Starts a service."
func NewOpnsenseStartService(target Target, opnsenseStartServiceParams OpnsenseStartServiceParams) *OpnsenseStartService {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "start_service"
	o := &OpnsenseStartService{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: opnsenseStartServiceParams,
	}
	return o
}

type OpnsenseStartService struct {
	ServiceBase
	ServiceData OpnsenseStartServiceParams `json:"service_data,omitempty"`
}
type OpnsenseStartServiceParams struct{}

func (o *OpnsenseStartService) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseStartService) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseStopService creates the object that can be sent to Home Assistant for domain opnsense, service stop_service
// "Stops a service."
func NewOpnsenseStopService(target Target, opnsenseStopServiceParams OpnsenseStopServiceParams) *OpnsenseStopService {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "stop_service"
	o := &OpnsenseStopService{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: opnsenseStopServiceParams,
	}
	return o
}

type OpnsenseStopService struct {
	ServiceBase
	ServiceData OpnsenseStopServiceParams `json:"service_data,omitempty"`
}
type OpnsenseStopServiceParams struct{}

func (o *OpnsenseStopService) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseStopService) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseSystemHalt creates the object that can be sent to Home Assistant for domain opnsense, service system_halt
// "Halts the system."
func NewOpnsenseSystemHalt(target Target, opnsenseSystemHaltParams OpnsenseSystemHaltParams) *OpnsenseSystemHalt {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "system_halt"
	o := &OpnsenseSystemHalt{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: opnsenseSystemHaltParams,
	}
	return o
}

type OpnsenseSystemHalt struct {
	ServiceBase
	ServiceData OpnsenseSystemHaltParams `json:"service_data,omitempty"`
}
type OpnsenseSystemHaltParams struct{}

func (o *OpnsenseSystemHalt) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseSystemHalt) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseSystemReboot creates the object that can be sent to Home Assistant for domain opnsense, service system_reboot
// "Reboots the system."
func NewOpnsenseSystemReboot(target Target, opnsenseSystemRebootParams OpnsenseSystemRebootParams) *OpnsenseSystemReboot {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "system_reboot"
	o := &OpnsenseSystemReboot{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: opnsenseSystemRebootParams,
	}
	return o
}

type OpnsenseSystemReboot struct {
	ServiceBase
	ServiceData OpnsenseSystemRebootParams `json:"service_data,omitempty"`
}
type OpnsenseSystemRebootParams struct{}

func (o *OpnsenseSystemReboot) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseSystemReboot) SetID(id *int) {
	o.Id = id
}
