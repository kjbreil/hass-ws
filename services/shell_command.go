package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewShellCommandShutdownTyr creates the object that can be sent to Home Assistant for domain shell_command, service shutdown_tyr
// ""
func NewShellCommandShutdownTyr(target Target, shellCommandShutdownTyrParams ShellCommandShutdownTyrParams) *ShellCommandShutdownTyr {
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
		ServiceData: shellCommandShutdownTyrParams,
	}
	return s
}

type ShellCommandShutdownTyr struct {
	ServiceBase
	ServiceData ShellCommandShutdownTyrParams `json:"service_data,omitempty"`
}
type ShellCommandShutdownTyrParams struct{}

func (s *ShellCommandShutdownTyr) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ShellCommandShutdownTyr) SetID(id *int) {
	s.Id = id
}

// NewShellCommandStartNcbp creates the object that can be sent to Home Assistant for domain shell_command, service start_ncbp
// ""
func NewShellCommandStartNcbp(target Target, shellCommandStartNcbpParams ShellCommandStartNcbpParams) *ShellCommandStartNcbp {
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
		ServiceData: shellCommandStartNcbpParams,
	}
	return s
}

type ShellCommandStartNcbp struct {
	ServiceBase
	ServiceData ShellCommandStartNcbpParams `json:"service_data,omitempty"`
}
type ShellCommandStartNcbpParams struct{}

func (s *ShellCommandStartNcbp) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ShellCommandStartNcbp) SetID(id *int) {
	s.Id = id
}

// NewShellCommandStopNcbp creates the object that can be sent to Home Assistant for domain shell_command, service stop_ncbp
// ""
func NewShellCommandStopNcbp(target Target, shellCommandStopNcbpParams ShellCommandStopNcbpParams) *ShellCommandStopNcbp {
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
		ServiceData: shellCommandStopNcbpParams,
	}
	return s
}

type ShellCommandStopNcbp struct {
	ServiceBase
	ServiceData ShellCommandStopNcbpParams `json:"service_data,omitempty"`
}
type ShellCommandStopNcbpParams struct{}

func (s *ShellCommandStopNcbp) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ShellCommandStopNcbp) SetID(id *int) {
	s.Id = id
}
