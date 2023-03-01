package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewInputSelectReload creates the object that can be sent to Home Assistant for domain input_select, service reload
// "Reload the input_select configuration."
func NewInputSelectReload(entities []string) *InputSelectReload {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputSelectReload{
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

type InputSelectReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputSelectReload) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputSelectReload) SetID(id *int) {
	i.Id = id
}

// NewInputSelectSelectFirst creates the object that can be sent to Home Assistant for domain input_select, service select_first
// "Select the first option of an input select entity."
func NewInputSelectSelectFirst(entities []string) *InputSelectSelectFirst {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_first"
	i := &InputSelectSelectFirst{
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

type InputSelectSelectFirst struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputSelectSelectFirst) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectFirst) SetID(id *int) {
	i.Id = id
}

// NewInputSelectSelectLast creates the object that can be sent to Home Assistant for domain input_select, service select_last
// "Select the last option of an input select entity."
func NewInputSelectSelectLast(entities []string) *InputSelectSelectLast {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_last"
	i := &InputSelectSelectLast{
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

type InputSelectSelectLast struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputSelectSelectLast) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectLast) SetID(id *int) {
	i.Id = id
}

// NewInputSelectSelectNext creates the object that can be sent to Home Assistant for domain input_select, service select_next
// "Select the next options of an input select entity."
func NewInputSelectSelectNext(entities []string) *InputSelectSelectNext {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_next"
	i := &InputSelectSelectNext{
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

type InputSelectSelectNext struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputSelectSelectNext) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectNext) SetID(id *int) {
	i.Id = id
}

// NewInputSelectSelectOption creates the object that can be sent to Home Assistant for domain input_select, service select_option
// "Select an option of an input select entity."
func NewInputSelectSelectOption(entities []string, option *string) *InputSelectSelectOption {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_option"
	i := &InputSelectSelectOption{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Option *string `json:"option,omitempty"`
		}{Option: option},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return i
}

type InputSelectSelectOption struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Option *string `json:"option,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputSelectSelectOption) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectOption) SetID(id *int) {
	i.Id = id
}

// NewInputSelectSelectPrevious creates the object that can be sent to Home Assistant for domain input_select, service select_previous
// "Select the previous options of an input select entity."
func NewInputSelectSelectPrevious(entities []string) *InputSelectSelectPrevious {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_previous"
	i := &InputSelectSelectPrevious{
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

type InputSelectSelectPrevious struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputSelectSelectPrevious) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectPrevious) SetID(id *int) {
	i.Id = id
}

// NewInputSelectSetOptions creates the object that can be sent to Home Assistant for domain input_select, service set_options
// "Set the options of an input select entity."
func NewInputSelectSetOptions(entities []string) *InputSelectSetOptions {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "set_options"
	i := &InputSelectSetOptions{
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

type InputSelectSetOptions struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (i *InputSelectSetOptions) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputSelectSetOptions) SetID(id *int) {
	i.Id = id
}
