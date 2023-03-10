package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSystemLogClear creates the object that can be sent to Home Assistant for domain system_log, service clear
// "Clear all log entries."
func NewSystemLogClear(target Target) *SystemLogClear {
	serviceDomain := "system_log"
	serviceType := "call_service"
	serviceService := "clear"
	s := &SystemLogClear{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return s
}

type SystemLogClear struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *SystemLogClear) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SystemLogClear) SetID(id *int) {
	s.Id = id
}

// NewSystemLogWrite creates the object that can be sent to Home Assistant for domain system_log, service write
// "Write log entry."
func NewSystemLogWrite(target Target, systemLogWriteParams *SystemLogWriteParams) *SystemLogWrite {
	serviceDomain := "system_log"
	serviceType := "call_service"
	serviceService := "write"
	s := &SystemLogWrite{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *systemLogWriteParams,
	}
	return s
}

type SystemLogWrite struct {
	ServiceBase
	ServiceData SystemLogWriteParams `json:"service_data,omitempty"`
}
type SystemLogWriteParams struct {
	Level   *Level  `json:"level,omitempty"`
	Logger  *string `json:"logger,omitempty"`
	Message *string `json:"message,omitempty"`
}

func (s *SystemLogWrite) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SystemLogWrite) SetID(id *int) {
	s.Id = id
}
