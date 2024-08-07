package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewScriptReload creates the object that can be sent to Home Assistant for domain script, service reload
// "Reload all the available scripts"
func NewScriptReload(target Target) *ScriptReload {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "reload"
	s := &ScriptReload{
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

type ScriptReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *ScriptReload) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *ScriptReload) Targets() []string {
	return s.Target.EntityId
}
func (s *ScriptReload) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}

// NewScriptTestScript creates the object that can be sent to Home Assistant for domain script, service test_script
// ""
func NewScriptTestScript(target Target) *ScriptTestScript {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "test_script"
	s := &ScriptTestScript{
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

type ScriptTestScript struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *ScriptTestScript) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *ScriptTestScript) Targets() []string {
	return s.Target.EntityId
}
func (s *ScriptTestScript) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}

// NewScriptToggle creates the object that can be sent to Home Assistant for domain script, service toggle
// "Toggle script"
func NewScriptToggle(target Target) *ScriptToggle {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "toggle"
	s := &ScriptToggle{
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

type ScriptToggle struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *ScriptToggle) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *ScriptToggle) Targets() []string {
	return s.Target.EntityId
}
func (s *ScriptToggle) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}

// NewScriptTurnOff creates the object that can be sent to Home Assistant for domain script, service turn_off
// "Turn off script"
func NewScriptTurnOff(target Target) *ScriptTurnOff {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "turn_off"
	s := &ScriptTurnOff{
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

type ScriptTurnOff struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *ScriptTurnOff) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *ScriptTurnOff) Targets() []string {
	return s.Target.EntityId
}
func (s *ScriptTurnOff) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}

// NewScriptTurnOn creates the object that can be sent to Home Assistant for domain script, service turn_on
// "Turn on script"
func NewScriptTurnOn(target Target) *ScriptTurnOn {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "turn_on"
	s := &ScriptTurnOn{
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

type ScriptTurnOn struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *ScriptTurnOn) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *ScriptTurnOn) Targets() []string {
	return s.Target.EntityId
}
func (s *ScriptTurnOn) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
