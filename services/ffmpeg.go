package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewFfmpegRestart creates the object that can be sent to Home Assistant for domain ffmpeg, service restart
// "Send a restart command to a ffmpeg based sensor."
func NewFfmpegRestart(entities []string) *FfmpegRestart {
	serviceDomain := "ffmpeg"
	serviceType := "call_service"
	serviceService := "restart"
	f := &FfmpegRestart{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FfmpegRestart struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (f *FfmpegRestart) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FfmpegRestart) SetID(id *int) {
	f.Id = id
}

// NewFfmpegStart creates the object that can be sent to Home Assistant for domain ffmpeg, service start
// "Send a start command to a ffmpeg based sensor."
func NewFfmpegStart(entities []string) *FfmpegStart {
	serviceDomain := "ffmpeg"
	serviceType := "call_service"
	serviceService := "start"
	f := &FfmpegStart{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FfmpegStart struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (f *FfmpegStart) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FfmpegStart) SetID(id *int) {
	f.Id = id
}

// NewFfmpegStop creates the object that can be sent to Home Assistant for domain ffmpeg, service stop
// "Send a stop command to a ffmpeg based sensor."
func NewFfmpegStop(entities []string) *FfmpegStop {
	serviceDomain := "ffmpeg"
	serviceType := "call_service"
	serviceService := "stop"
	f := &FfmpegStop{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FfmpegStop struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (f *FfmpegStop) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FfmpegStop) SetID(id *int) {
	f.Id = id
}
