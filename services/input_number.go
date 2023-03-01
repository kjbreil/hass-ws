package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewInputNumberDecrement creates the object that can be sent to Home Assistant for domain input_number, service decrement
// "Decrement the value of an input number entity by its stepping."
func NewInputNumberDecrement(entities []string) *InputNumberDecrement {
	serviceDomain := "input_number"
	serviceType := "call_service"
	serviceService := "decrement"
	i := &InputNumberDecrement{
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

type InputNumberDecrement struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputNumberDecrement) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputNumberDecrement) SetID(id *int) {
	i.Id = id
}

// NewInputNumberIncrement creates the object that can be sent to Home Assistant for domain input_number, service increment
// "Increment the value of an input number entity by its stepping."
func NewInputNumberIncrement(entities []string) *InputNumberIncrement {
	serviceDomain := "input_number"
	serviceType := "call_service"
	serviceService := "increment"
	i := &InputNumberIncrement{
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

type InputNumberIncrement struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputNumberIncrement) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputNumberIncrement) SetID(id *int) {
	i.Id = id
}

// NewInputNumberReload creates the object that can be sent to Home Assistant for domain input_number, service reload
// "Reload the input_number configuration."
func NewInputNumberReload(entities []string) *InputNumberReload {
	serviceDomain := "input_number"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputNumberReload{
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

type InputNumberReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputNumberReload) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputNumberReload) SetID(id *int) {
	i.Id = id
}

// NewInputNumberSetValue creates the object that can be sent to Home Assistant for domain input_number, service set_value
// "Set the value of an input number entity."
func NewInputNumberSetValue(entities []string, value *int) *InputNumberSetValue {
	serviceDomain := "input_number"
	serviceType := "call_service"
	serviceService := "set_value"
	i := &InputNumberSetValue{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Value *int `json:"value,omitempty"`
		}{Value: value},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return i
}

type InputNumberSetValue struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Value *int `json:"value,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputNumberSetValue) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputNumberSetValue) SetID(id *int) {
	i.Id = id
}
