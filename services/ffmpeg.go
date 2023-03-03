package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewFfmpegRestart creates the object that can be sent to Home Assistant for domain ffmpeg, service restart
// "Send a restart command to a ffmpeg based sensor."
func NewFfmpegRestart(target Target, ffmpegRestartParams FfmpegRestartParams) *FfmpegRestart {
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
		ServiceData: ffmpegRestartParams,
	}
	return f
}

type FfmpegRestart struct {
	ServiceBase
	ServiceData FfmpegRestartParams `json:"service_data,omitempty"`
}
type FfmpegRestartParams struct{}

func (f *FfmpegRestart) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FfmpegRestart) SetID(id *int) {
	f.Id = id
}

// NewFfmpegStart creates the object that can be sent to Home Assistant for domain ffmpeg, service start
// "Send a start command to a ffmpeg based sensor."
func NewFfmpegStart(target Target, ffmpegStartParams FfmpegStartParams) *FfmpegStart {
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
		ServiceData: ffmpegStartParams,
	}
	return f
}

type FfmpegStart struct {
	ServiceBase
	ServiceData FfmpegStartParams `json:"service_data,omitempty"`
}
type FfmpegStartParams struct{}

func (f *FfmpegStart) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FfmpegStart) SetID(id *int) {
	f.Id = id
}

// NewFfmpegStop creates the object that can be sent to Home Assistant for domain ffmpeg, service stop
// "Send a stop command to a ffmpeg based sensor."
func NewFfmpegStop(target Target, ffmpegStopParams FfmpegStopParams) *FfmpegStop {
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
		ServiceData: ffmpegStopParams,
	}
	return f
}

type FfmpegStop struct {
	ServiceBase
	ServiceData FfmpegStopParams `json:"service_data,omitempty"`
}
type FfmpegStopParams struct{}

func (f *FfmpegStop) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FfmpegStop) SetID(id *int) {
	f.Id = id
}
