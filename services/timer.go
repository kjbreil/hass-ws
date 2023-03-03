package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewTimerCancel creates the object that can be sent to Home Assistant for domain timer, service cancel
// "Cancel a timer."
func NewTimerCancel(target Target, timerCancelParams TimerCancelParams) *TimerCancel {
	serviceDomain := "timer"
	serviceType := "call_service"
	serviceService := "cancel"
	t := &TimerCancel{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: timerCancelParams,
	}
	return t
}

type TimerCancel struct {
	ServiceBase
	ServiceData TimerCancelParams `json:"service_data,omitempty"`
}
type TimerCancelParams struct{}

func (t *TimerCancel) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TimerCancel) SetID(id *int) {
	t.Id = id
}

// NewTimerFinish creates the object that can be sent to Home Assistant for domain timer, service finish
// "Finish a timer."
func NewTimerFinish(target Target, timerFinishParams TimerFinishParams) *TimerFinish {
	serviceDomain := "timer"
	serviceType := "call_service"
	serviceService := "finish"
	t := &TimerFinish{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: timerFinishParams,
	}
	return t
}

type TimerFinish struct {
	ServiceBase
	ServiceData TimerFinishParams `json:"service_data,omitempty"`
}
type TimerFinishParams struct{}

func (t *TimerFinish) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TimerFinish) SetID(id *int) {
	t.Id = id
}

// NewTimerPause creates the object that can be sent to Home Assistant for domain timer, service pause
// "Pause a timer."
func NewTimerPause(target Target, timerPauseParams TimerPauseParams) *TimerPause {
	serviceDomain := "timer"
	serviceType := "call_service"
	serviceService := "pause"
	t := &TimerPause{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: timerPauseParams,
	}
	return t
}

type TimerPause struct {
	ServiceBase
	ServiceData TimerPauseParams `json:"service_data,omitempty"`
}
type TimerPauseParams struct{}

func (t *TimerPause) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TimerPause) SetID(id *int) {
	t.Id = id
}

// NewTimerReload creates the object that can be sent to Home Assistant for domain timer, service reload
// ""
func NewTimerReload(target Target, timerReloadParams TimerReloadParams) *TimerReload {
	serviceDomain := "timer"
	serviceType := "call_service"
	serviceService := "reload"
	t := &TimerReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: timerReloadParams,
	}
	return t
}

type TimerReload struct {
	ServiceBase
	ServiceData TimerReloadParams `json:"service_data,omitempty"`
}
type TimerReloadParams struct{}

func (t *TimerReload) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TimerReload) SetID(id *int) {
	t.Id = id
}

// NewTimerStart creates the object that can be sent to Home Assistant for domain timer, service start
// "Start a timer"
func NewTimerStart(target Target, timerStartParams TimerStartParams) *TimerStart {
	serviceDomain := "timer"
	serviceType := "call_service"
	serviceService := "start"
	t := &TimerStart{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: timerStartParams,
	}
	return t
}

type TimerStart struct {
	ServiceBase
	ServiceData TimerStartParams `json:"service_data,omitempty"`
}
type TimerStartParams struct {
	Duration *string `json:"duration,omitempty"`
}

func (t *TimerStart) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TimerStart) SetID(id *int) {
	t.Id = id
}
