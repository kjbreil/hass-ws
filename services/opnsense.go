package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewOpnsenseCloseNotice creates the object that can be sent to Home Assistant for domain opnsense, service close_notice
// "Closes a notice(s)."
func NewOpnsenseCloseNotice(target Target) *OpnsenseCloseNotice {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "close_notice"
	o := &OpnsenseCloseNotice{
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
	return o
}

type OpnsenseCloseNotice struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (o *OpnsenseCloseNotice) JSON() string {
	data, _ := gojson.Marshal(o)
	return string(data)
}
func (o *OpnsenseCloseNotice) Targets() []string {
	return o.Target.EntityId
}
func (o *OpnsenseCloseNotice) Name() string {
	return fmt.Sprintf("%s.%s", *o.Domain, *o.Service)
}

// NewOpnsenseFileNotice creates the object that can be sent to Home Assistant for domain opnsense, service file_notice
// "Files a notice(s)."
func NewOpnsenseFileNotice(target Target) *OpnsenseFileNotice {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "file_notice"
	o := &OpnsenseFileNotice{
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
	return o
}

type OpnsenseFileNotice struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (o *OpnsenseFileNotice) JSON() string {
	data, _ := gojson.Marshal(o)
	return string(data)
}
func (o *OpnsenseFileNotice) Targets() []string {
	return o.Target.EntityId
}
func (o *OpnsenseFileNotice) Name() string {
	return fmt.Sprintf("%s.%s", *o.Domain, *o.Service)
}

// NewOpnsenseRestartService creates the object that can be sent to Home Assistant for domain opnsense, service restart_service
// "Restarts a service."
func NewOpnsenseRestartService(target Target) *OpnsenseRestartService {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "restart_service"
	o := &OpnsenseRestartService{
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
	return o
}

type OpnsenseRestartService struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (o *OpnsenseRestartService) JSON() string {
	data, _ := gojson.Marshal(o)
	return string(data)
}
func (o *OpnsenseRestartService) Targets() []string {
	return o.Target.EntityId
}
func (o *OpnsenseRestartService) Name() string {
	return fmt.Sprintf("%s.%s", *o.Domain, *o.Service)
}

// NewOpnsenseSendWol creates the object that can be sent to Home Assistant for domain opnsense, service send_wol
// "Sends wake-on-lan magic packet."
func NewOpnsenseSendWol(target Target) *OpnsenseSendWol {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "send_wol"
	o := &OpnsenseSendWol{
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
	return o
}

type OpnsenseSendWol struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (o *OpnsenseSendWol) JSON() string {
	data, _ := gojson.Marshal(o)
	return string(data)
}
func (o *OpnsenseSendWol) Targets() []string {
	return o.Target.EntityId
}
func (o *OpnsenseSendWol) Name() string {
	return fmt.Sprintf("%s.%s", *o.Domain, *o.Service)
}

// NewOpnsenseStartService creates the object that can be sent to Home Assistant for domain opnsense, service start_service
// "Starts a service."
func NewOpnsenseStartService(target Target) *OpnsenseStartService {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "start_service"
	o := &OpnsenseStartService{
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
	return o
}

type OpnsenseStartService struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (o *OpnsenseStartService) JSON() string {
	data, _ := gojson.Marshal(o)
	return string(data)
}
func (o *OpnsenseStartService) Targets() []string {
	return o.Target.EntityId
}
func (o *OpnsenseStartService) Name() string {
	return fmt.Sprintf("%s.%s", *o.Domain, *o.Service)
}

// NewOpnsenseStopService creates the object that can be sent to Home Assistant for domain opnsense, service stop_service
// "Stops a service."
func NewOpnsenseStopService(target Target) *OpnsenseStopService {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "stop_service"
	o := &OpnsenseStopService{
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
	return o
}

type OpnsenseStopService struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (o *OpnsenseStopService) JSON() string {
	data, _ := gojson.Marshal(o)
	return string(data)
}
func (o *OpnsenseStopService) Targets() []string {
	return o.Target.EntityId
}
func (o *OpnsenseStopService) Name() string {
	return fmt.Sprintf("%s.%s", *o.Domain, *o.Service)
}

// NewOpnsenseSystemHalt creates the object that can be sent to Home Assistant for domain opnsense, service system_halt
// "Halts the system."
func NewOpnsenseSystemHalt(target Target) *OpnsenseSystemHalt {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "system_halt"
	o := &OpnsenseSystemHalt{
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
	return o
}

type OpnsenseSystemHalt struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (o *OpnsenseSystemHalt) JSON() string {
	data, _ := gojson.Marshal(o)
	return string(data)
}
func (o *OpnsenseSystemHalt) Targets() []string {
	return o.Target.EntityId
}
func (o *OpnsenseSystemHalt) Name() string {
	return fmt.Sprintf("%s.%s", *o.Domain, *o.Service)
}

// NewOpnsenseSystemReboot creates the object that can be sent to Home Assistant for domain opnsense, service system_reboot
// "Reboots the system."
func NewOpnsenseSystemReboot(target Target) *OpnsenseSystemReboot {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "system_reboot"
	o := &OpnsenseSystemReboot{
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
	return o
}

type OpnsenseSystemReboot struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (o *OpnsenseSystemReboot) JSON() string {
	data, _ := gojson.Marshal(o)
	return string(data)
}
func (o *OpnsenseSystemReboot) Targets() []string {
	return o.Target.EntityId
}
func (o *OpnsenseSystemReboot) Name() string {
	return fmt.Sprintf("%s.%s", *o.Domain, *o.Service)
}
