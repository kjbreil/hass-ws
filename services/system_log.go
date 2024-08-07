package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

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
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SystemLogClear) Targets() []string {
	return s.Target.EntityId
}
func (s *SystemLogClear) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}

// NewSystemLogWrite creates the object that can be sent to Home Assistant for domain system_log, service write
// "Write log entry."
func NewSystemLogWrite(target Target) *SystemLogWrite {
	serviceDomain := "system_log"
	serviceType := "call_service"
	serviceService := "write"
	s := &SystemLogWrite{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: SystemLogWriteParams{},
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

func (s *SystemLogWrite) Level(level Level) *SystemLogWrite {
	s.ServiceData.Level = &level
	return s
}
func (s *SystemLogWrite) Logger(logger string) *SystemLogWrite {
	s.ServiceData.Logger = &logger
	return s
}
func (s *SystemLogWrite) Message(message string) *SystemLogWrite {
	s.ServiceData.Message = &message
	return s
}
func (s *SystemLogWrite) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SystemLogWrite) Targets() []string {
	return s.Target.EntityId
}
func (s *SystemLogWrite) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
