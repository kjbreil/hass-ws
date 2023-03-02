package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewOpnsenseCloseNotice creates the object that can be sent to Home Assistant for domain opnsense, service close_notice
// "Closes a notice(s)."
func NewOpnsenseCloseNotice(entities []string) *OpnsenseCloseNotice {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "close_notice"
	o := &OpnsenseCloseNotice{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return o
}

type OpnsenseCloseNotice struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (o *OpnsenseCloseNotice) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseCloseNotice) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseFileNotice creates the object that can be sent to Home Assistant for domain opnsense, service file_notice
// "Files a notice(s)."
func NewOpnsenseFileNotice(entities []string) *OpnsenseFileNotice {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "file_notice"
	o := &OpnsenseFileNotice{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return o
}

type OpnsenseFileNotice struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (o *OpnsenseFileNotice) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseFileNotice) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseRestartService creates the object that can be sent to Home Assistant for domain opnsense, service restart_service
// "Restarts a service."
func NewOpnsenseRestartService(entities []string) *OpnsenseRestartService {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "restart_service"
	o := &OpnsenseRestartService{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return o
}

type OpnsenseRestartService struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (o *OpnsenseRestartService) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseRestartService) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseSendWol creates the object that can be sent to Home Assistant for domain opnsense, service send_wol
// "Sends wake-on-lan magic packet."
func NewOpnsenseSendWol(entities []string) *OpnsenseSendWol {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "send_wol"
	o := &OpnsenseSendWol{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return o
}

type OpnsenseSendWol struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (o *OpnsenseSendWol) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseSendWol) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseStartService creates the object that can be sent to Home Assistant for domain opnsense, service start_service
// "Starts a service."
func NewOpnsenseStartService(entities []string) *OpnsenseStartService {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "start_service"
	o := &OpnsenseStartService{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return o
}

type OpnsenseStartService struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (o *OpnsenseStartService) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseStartService) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseStopService creates the object that can be sent to Home Assistant for domain opnsense, service stop_service
// "Stops a service."
func NewOpnsenseStopService(entities []string) *OpnsenseStopService {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "stop_service"
	o := &OpnsenseStopService{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return o
}

type OpnsenseStopService struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (o *OpnsenseStopService) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseStopService) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseSystemHalt creates the object that can be sent to Home Assistant for domain opnsense, service system_halt
// "Halts the system."
func NewOpnsenseSystemHalt(entities []string) *OpnsenseSystemHalt {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "system_halt"
	o := &OpnsenseSystemHalt{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return o
}

type OpnsenseSystemHalt struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (o *OpnsenseSystemHalt) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseSystemHalt) SetID(id *int) {
	o.Id = id
}

// NewOpnsenseSystemReboot creates the object that can be sent to Home Assistant for domain opnsense, service system_reboot
// "Reboots the system."
func NewOpnsenseSystemReboot(entities []string) *OpnsenseSystemReboot {
	serviceDomain := "opnsense"
	serviceType := "call_service"
	serviceService := "system_reboot"
	o := &OpnsenseSystemReboot{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return o
}

type OpnsenseSystemReboot struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (o *OpnsenseSystemReboot) JSON() string {
	data, _ := json.Marshal(o)
	return string(data)
}
func (o *OpnsenseSystemReboot) SetID(id *int) {
	o.Id = id
}
