package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewShellCommandShutdownTyr creates the object that can be sent to Home Assistant for domain shell_command, service shutdown_tyr
// ""
func NewShellCommandShutdownTyr(entities []string) *ShellCommandShutdownTyr {
	serviceDomain := "shell_command"
	serviceType := "call_service"
	serviceService := "shutdown_tyr"
	s := &ShellCommandShutdownTyr{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type ShellCommandShutdownTyr struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *ShellCommandShutdownTyr) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ShellCommandShutdownTyr) SetID(id *int) {
	s.Id = id
}

// NewShellCommandStartNcbp creates the object that can be sent to Home Assistant for domain shell_command, service start_ncbp
// ""
func NewShellCommandStartNcbp(entities []string) *ShellCommandStartNcbp {
	serviceDomain := "shell_command"
	serviceType := "call_service"
	serviceService := "start_ncbp"
	s := &ShellCommandStartNcbp{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type ShellCommandStartNcbp struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *ShellCommandStartNcbp) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ShellCommandStartNcbp) SetID(id *int) {
	s.Id = id
}

// NewShellCommandStopNcbp creates the object that can be sent to Home Assistant for domain shell_command, service stop_ncbp
// ""
func NewShellCommandStopNcbp(entities []string) *ShellCommandStopNcbp {
	serviceDomain := "shell_command"
	serviceType := "call_service"
	serviceService := "stop_ncbp"
	s := &ShellCommandStopNcbp{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type ShellCommandStopNcbp struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *ShellCommandStopNcbp) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ShellCommandStopNcbp) SetID(id *int) {
	s.Id = id
}
