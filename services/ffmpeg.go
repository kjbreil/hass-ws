package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewFfmpegRestart creates the object that can be sent to Home Assistant for domain ffmpeg, service restart
// "Send a restart command to a ffmpeg based sensor."
func NewFfmpegRestart(target Target) *FfmpegRestart {
	serviceDomain := "ffmpeg"
	serviceType := "call_service"
	serviceService := "restart"
	f := &FfmpegRestart{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return f
}

type FfmpegRestart struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (f *FfmpegRestart) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FfmpegRestart) Name() string {
	return fmt.Sprintf("%s.%s", *f.Domain, *f.Service)
}
func (f *FfmpegRestart) SetID(id *int) {
	f.Id = id
}

// NewFfmpegStart creates the object that can be sent to Home Assistant for domain ffmpeg, service start
// "Send a start command to a ffmpeg based sensor."
func NewFfmpegStart(target Target) *FfmpegStart {
	serviceDomain := "ffmpeg"
	serviceType := "call_service"
	serviceService := "start"
	f := &FfmpegStart{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return f
}

type FfmpegStart struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (f *FfmpegStart) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FfmpegStart) Name() string {
	return fmt.Sprintf("%s.%s", *f.Domain, *f.Service)
}
func (f *FfmpegStart) SetID(id *int) {
	f.Id = id
}

// NewFfmpegStop creates the object that can be sent to Home Assistant for domain ffmpeg, service stop
// "Send a stop command to a ffmpeg based sensor."
func NewFfmpegStop(target Target) *FfmpegStop {
	serviceDomain := "ffmpeg"
	serviceType := "call_service"
	serviceService := "stop"
	f := &FfmpegStop{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return f
}

type FfmpegStop struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (f *FfmpegStop) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FfmpegStop) Name() string {
	return fmt.Sprintf("%s.%s", *f.Domain, *f.Service)
}
func (f *FfmpegStop) SetID(id *int) {
	f.Id = id
}
