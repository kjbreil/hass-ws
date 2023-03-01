package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewInputBooleanReload creates the object that can be sent to Home Assistant for domain input_boolean, service reload
// "Reload the input_boolean configuration"
func NewInputBooleanReload(entities []string) *InputBooleanReload {
	serviceDomain := "input_boolean"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputBooleanReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return i
}

type InputBooleanReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputBooleanReload) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputBooleanReload) SetID(id *int) {
	i.Id = id
}

// NewInputBooleanToggle creates the object that can be sent to Home Assistant for domain input_boolean, service toggle
// "Toggle an input boolean"
func NewInputBooleanToggle(entities []string) *InputBooleanToggle {
	serviceDomain := "input_boolean"
	serviceType := "call_service"
	serviceService := "toggle"
	i := &InputBooleanToggle{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return i
}

type InputBooleanToggle struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputBooleanToggle) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputBooleanToggle) SetID(id *int) {
	i.Id = id
}

// NewInputBooleanTurnOff creates the object that can be sent to Home Assistant for domain input_boolean, service turn_off
// "Turn off an input boolean"
func NewInputBooleanTurnOff(entities []string) *InputBooleanTurnOff {
	serviceDomain := "input_boolean"
	serviceType := "call_service"
	serviceService := "turn_off"
	i := &InputBooleanTurnOff{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return i
}

type InputBooleanTurnOff struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputBooleanTurnOff) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputBooleanTurnOff) SetID(id *int) {
	i.Id = id
}

// NewInputBooleanTurnOn creates the object that can be sent to Home Assistant for domain input_boolean, service turn_on
// "Turn on an input boolean"
func NewInputBooleanTurnOn(entities []string) *InputBooleanTurnOn {
	serviceDomain := "input_boolean"
	serviceType := "call_service"
	serviceService := "turn_on"
	i := &InputBooleanTurnOn{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return i
}

type InputBooleanTurnOn struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputBooleanTurnOn) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputBooleanTurnOn) SetID(id *int) {
	i.Id = id
}
