package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewFfmpegRestart creates the object that can be sent to Home Assistant for domain ffmpeg, service restart
// "Sends a restart command to an FFmpeg-based sensor."
func NewFfmpegRestart(target Target) *FfmpegRestart {
	serviceDomain := "ffmpeg"
	serviceType := "call_service"
	serviceService := "restart"
	f := &FfmpegRestart{
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
	return f
}

type FfmpegRestart struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (f *FfmpegRestart) JSON() string {
	data, _ := gojson.Marshal(f)
	return string(data)
}
func (f *FfmpegRestart) Targets() []string {
	return f.Target.EntityId
}
func (f *FfmpegRestart) Name() string {
	return fmt.Sprintf("%s.%s", *f.Domain, *f.Service)
}

// NewFfmpegStart creates the object that can be sent to Home Assistant for domain ffmpeg, service start
// "Sends a start command to an FFmpeg-based sensor."
func NewFfmpegStart(target Target) *FfmpegStart {
	serviceDomain := "ffmpeg"
	serviceType := "call_service"
	serviceService := "start"
	f := &FfmpegStart{
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
	return f
}

type FfmpegStart struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (f *FfmpegStart) JSON() string {
	data, _ := gojson.Marshal(f)
	return string(data)
}
func (f *FfmpegStart) Targets() []string {
	return f.Target.EntityId
}
func (f *FfmpegStart) Name() string {
	return fmt.Sprintf("%s.%s", *f.Domain, *f.Service)
}

// NewFfmpegStop creates the object that can be sent to Home Assistant for domain ffmpeg, service stop
// "Sends a stop command to an FFmpeg-based sensor."
func NewFfmpegStop(target Target) *FfmpegStop {
	serviceDomain := "ffmpeg"
	serviceType := "call_service"
	serviceService := "stop"
	f := &FfmpegStop{
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
	return f
}

type FfmpegStop struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (f *FfmpegStop) JSON() string {
	data, _ := gojson.Marshal(f)
	return string(data)
}
func (f *FfmpegStop) Targets() []string {
	return f.Target.EntityId
}
func (f *FfmpegStop) Name() string {
	return fmt.Sprintf("%s.%s", *f.Domain, *f.Service)
}
