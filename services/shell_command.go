package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewShellCommandShutdownTyr creates the object that can be sent to Home Assistant for domain shell_command, service shutdown_tyr
// ""
func NewShellCommandShutdownTyr(target Target) *ShellCommandShutdownTyr {
	serviceDomain := "shell_command"
	serviceType := "call_service"
	serviceService := "shutdown_tyr"
	s := &ShellCommandShutdownTyr{
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

type ShellCommandShutdownTyr struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *ShellCommandShutdownTyr) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *ShellCommandShutdownTyr) Targets() []string {
	return s.Target.EntityId
}
func (s *ShellCommandShutdownTyr) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *ShellCommandShutdownTyr) SetID(id *int) {
	s.Id = id
}

// NewShellCommandStartNcbp creates the object that can be sent to Home Assistant for domain shell_command, service start_ncbp
// ""
func NewShellCommandStartNcbp(target Target) *ShellCommandStartNcbp {
	serviceDomain := "shell_command"
	serviceType := "call_service"
	serviceService := "start_ncbp"
	s := &ShellCommandStartNcbp{
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

type ShellCommandStartNcbp struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *ShellCommandStartNcbp) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *ShellCommandStartNcbp) Targets() []string {
	return s.Target.EntityId
}
func (s *ShellCommandStartNcbp) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *ShellCommandStartNcbp) SetID(id *int) {
	s.Id = id
}

// NewShellCommandStopNcbp creates the object that can be sent to Home Assistant for domain shell_command, service stop_ncbp
// ""
func NewShellCommandStopNcbp(target Target) *ShellCommandStopNcbp {
	serviceDomain := "shell_command"
	serviceType := "call_service"
	serviceService := "stop_ncbp"
	s := &ShellCommandStopNcbp{
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

type ShellCommandStopNcbp struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *ShellCommandStopNcbp) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *ShellCommandStopNcbp) Targets() []string {
	return s.Target.EntityId
}
func (s *ShellCommandStopNcbp) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *ShellCommandStopNcbp) SetID(id *int) {
	s.Id = id
}
