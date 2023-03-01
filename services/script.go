package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewScriptReload creates the object that can be sent to Home Assistant for domain script, service reload
// "Reload all the available scripts"
func NewScriptReload(entities []string) *ScriptReload {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "reload"
	s := &ScriptReload{
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

type ScriptReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *ScriptReload) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ScriptReload) SetID(id *int) {
	s.Id = id
}

// NewScriptTestScript creates the object that can be sent to Home Assistant for domain script, service test_script
// ""
func NewScriptTestScript(entities []string) *ScriptTestScript {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "test_script"
	s := &ScriptTestScript{
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

type ScriptTestScript struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *ScriptTestScript) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ScriptTestScript) SetID(id *int) {
	s.Id = id
}

// NewScriptToggle creates the object that can be sent to Home Assistant for domain script, service toggle
// "Toggle script"
func NewScriptToggle(entities []string) *ScriptToggle {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "toggle"
	s := &ScriptToggle{
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

type ScriptToggle struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *ScriptToggle) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ScriptToggle) SetID(id *int) {
	s.Id = id
}

// NewScriptTurnOff creates the object that can be sent to Home Assistant for domain script, service turn_off
// "Turn off script"
func NewScriptTurnOff(entities []string) *ScriptTurnOff {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "turn_off"
	s := &ScriptTurnOff{
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

type ScriptTurnOff struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *ScriptTurnOff) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ScriptTurnOff) SetID(id *int) {
	s.Id = id
}

// NewScriptTurnOn creates the object that can be sent to Home Assistant for domain script, service turn_on
// "Turn on script"
func NewScriptTurnOn(entities []string) *ScriptTurnOn {
	serviceDomain := "script"
	serviceType := "call_service"
	serviceService := "turn_on"
	s := &ScriptTurnOn{
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

type ScriptTurnOn struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *ScriptTurnOn) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *ScriptTurnOn) SetID(id *int) {
	s.Id = id
}
