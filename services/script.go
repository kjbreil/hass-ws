package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewScriptReload creates the object that can be sent to Home Assistant for domain script, service reload
// "Reload all the available scripts"
func NewScriptReload(target Target, scriptReloadParams ScriptReloadParams) *ScriptReload {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "reload"
	s := &ScriptReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: scriptReloadParams,
	}
	return s
}

type ScriptReload struct {
	ServiceBase
	ServiceData ScriptReloadParams `json:"service_data,omitempty"`
}
type ScriptReloadParams struct{}

func (s *ScriptReload) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ScriptReload) SetID(id *int) {
	s.Id = id
}

// NewScriptTestScript creates the object that can be sent to Home Assistant for domain script, service test_script
// ""
func NewScriptTestScript(target Target, scriptTestScriptParams ScriptTestScriptParams) *ScriptTestScript {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "test_script"
	s := &ScriptTestScript{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: scriptTestScriptParams,
	}
	return s
}

type ScriptTestScript struct {
	ServiceBase
	ServiceData ScriptTestScriptParams `json:"service_data,omitempty"`
}
type ScriptTestScriptParams struct{}

func (s *ScriptTestScript) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ScriptTestScript) SetID(id *int) {
	s.Id = id
}

// NewScriptToggle creates the object that can be sent to Home Assistant for domain script, service toggle
// "Toggle script"
func NewScriptToggle(target Target, scriptToggleParams ScriptToggleParams) *ScriptToggle {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "toggle"
	s := &ScriptToggle{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: scriptToggleParams,
	}
	return s
}

type ScriptToggle struct {
	ServiceBase
	ServiceData ScriptToggleParams `json:"service_data,omitempty"`
}
type ScriptToggleParams struct{}

func (s *ScriptToggle) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ScriptToggle) SetID(id *int) {
	s.Id = id
}

// NewScriptTurnOff creates the object that can be sent to Home Assistant for domain script, service turn_off
// "Turn off script"
func NewScriptTurnOff(target Target, scriptTurnOffParams ScriptTurnOffParams) *ScriptTurnOff {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "turn_off"
	s := &ScriptTurnOff{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: scriptTurnOffParams,
	}
	return s
}

type ScriptTurnOff struct {
	ServiceBase
	ServiceData ScriptTurnOffParams `json:"service_data,omitempty"`
}
type ScriptTurnOffParams struct{}

func (s *ScriptTurnOff) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ScriptTurnOff) SetID(id *int) {
	s.Id = id
}

// NewScriptTurnOn creates the object that can be sent to Home Assistant for domain script, service turn_on
// "Turn on script"
func NewScriptTurnOn(target Target, scriptTurnOnParams ScriptTurnOnParams) *ScriptTurnOn {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "turn_on"
	s := &ScriptTurnOn{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: scriptTurnOnParams,
	}
	return s
}

type ScriptTurnOn struct {
	ServiceBase
	ServiceData ScriptTurnOnParams `json:"service_data,omitempty"`
}
type ScriptTurnOnParams struct{}

func (s *ScriptTurnOn) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ScriptTurnOn) SetID(id *int) {
	s.Id = id
}
