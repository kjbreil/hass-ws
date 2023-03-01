package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewTimerCancel creates the object that can be sent to Home Assistant for domain timer, service cancel
// "Cancel a timer."
func NewTimerCancel(entities []string) *TimerCancel {
	serviceDomain := "timer"
	serviceType := "call_service"
	serviceService := "cancel"
	t := &TimerCancel{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return t
}

type TimerCancel struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (t *TimerCancel) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TimerCancel) SetID(id *int) {
	t.Id = id
}

// NewTimerFinish creates the object that can be sent to Home Assistant for domain timer, service finish
// "Finish a timer."
func NewTimerFinish(entities []string) *TimerFinish {
	serviceDomain := "timer"
	serviceType := "call_service"
	serviceService := "finish"
	t := &TimerFinish{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return t
}

type TimerFinish struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (t *TimerFinish) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TimerFinish) SetID(id *int) {
	t.Id = id
}

// NewTimerPause creates the object that can be sent to Home Assistant for domain timer, service pause
// "Pause a timer."
func NewTimerPause(entities []string) *TimerPause {
	serviceDomain := "timer"
	serviceType := "call_service"
	serviceService := "pause"
	t := &TimerPause{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return t
}

type TimerPause struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (t *TimerPause) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TimerPause) SetID(id *int) {
	t.Id = id
}

// NewTimerReload creates the object that can be sent to Home Assistant for domain timer, service reload
// ""
func NewTimerReload(entities []string) *TimerReload {
	serviceDomain := "timer"
	serviceType := "call_service"
	serviceService := "reload"
	t := &TimerReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return t
}

type TimerReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (t *TimerReload) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TimerReload) SetID(id *int) {
	t.Id = id
}

// NewTimerStart creates the object that can be sent to Home Assistant for domain timer, service start
// "Start a timer"
func NewTimerStart(entities []string, duration *string) *TimerStart {
	serviceDomain := "timer"
	serviceType := "call_service"
	serviceService := "start"
	t := &TimerStart{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Duration *string `json:"duration,omitempty"`
		}{Duration: duration},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return t
}

type TimerStart struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Duration *string `json:"duration,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (t *TimerStart) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TimerStart) SetID(id *int) {
	t.Id = id
}
