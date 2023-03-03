package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewInputSelectReload creates the object that can be sent to Home Assistant for domain input_select, service reload
// "Reload the input_select configuration."
func NewInputSelectReload(target Target, inputSelectReloadParams InputSelectReloadParams) *InputSelectReload {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputSelectReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: inputSelectReloadParams,
	}
	return i
}

type InputSelectReload struct {
	ServiceBase
	ServiceData InputSelectReloadParams `json:"service_data,omitempty"`
}
type InputSelectReloadParams struct{}

func (i *InputSelectReload) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputSelectReload) SetID(id *int) {
	i.Id = id
}

// NewInputSelectSelectFirst creates the object that can be sent to Home Assistant for domain input_select, service select_first
// "Select the first option of an input select entity."
func NewInputSelectSelectFirst(target Target, inputSelectSelectFirstParams InputSelectSelectFirstParams) *InputSelectSelectFirst {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_first"
	i := &InputSelectSelectFirst{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: inputSelectSelectFirstParams,
	}
	return i
}

type InputSelectSelectFirst struct {
	ServiceBase
	ServiceData InputSelectSelectFirstParams `json:"service_data,omitempty"`
}
type InputSelectSelectFirstParams struct{}

func (i *InputSelectSelectFirst) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectFirst) SetID(id *int) {
	i.Id = id
}

// NewInputSelectSelectLast creates the object that can be sent to Home Assistant for domain input_select, service select_last
// "Select the last option of an input select entity."
func NewInputSelectSelectLast(target Target, inputSelectSelectLastParams InputSelectSelectLastParams) *InputSelectSelectLast {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_last"
	i := &InputSelectSelectLast{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: inputSelectSelectLastParams,
	}
	return i
}

type InputSelectSelectLast struct {
	ServiceBase
	ServiceData InputSelectSelectLastParams `json:"service_data,omitempty"`
}
type InputSelectSelectLastParams struct{}

func (i *InputSelectSelectLast) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectLast) SetID(id *int) {
	i.Id = id
}

// NewInputSelectSelectNext creates the object that can be sent to Home Assistant for domain input_select, service select_next
// "Select the next options of an input select entity."
func NewInputSelectSelectNext(target Target, inputSelectSelectNextParams InputSelectSelectNextParams) *InputSelectSelectNext {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_next"
	i := &InputSelectSelectNext{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: inputSelectSelectNextParams,
	}
	return i
}

type InputSelectSelectNext struct {
	ServiceBase
	ServiceData InputSelectSelectNextParams `json:"service_data,omitempty"`
}
type InputSelectSelectNextParams struct{}

func (i *InputSelectSelectNext) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectNext) SetID(id *int) {
	i.Id = id
}

// NewInputSelectSelectOption creates the object that can be sent to Home Assistant for domain input_select, service select_option
// "Select an option of an input select entity."
func NewInputSelectSelectOption(target Target, inputSelectSelectOptionParams InputSelectSelectOptionParams) *InputSelectSelectOption {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_option"
	i := &InputSelectSelectOption{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: inputSelectSelectOptionParams,
	}
	return i
}

type InputSelectSelectOption struct {
	ServiceBase
	ServiceData InputSelectSelectOptionParams `json:"service_data,omitempty"`
}
type InputSelectSelectOptionParams struct {
	Option *string `json:"option,omitempty"`
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
func NewInputSelectSelectPrevious(target Target, inputSelectSelectPreviousParams InputSelectSelectPreviousParams) *InputSelectSelectPrevious {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_previous"
	i := &InputSelectSelectPrevious{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: inputSelectSelectPreviousParams,
	}
	return i
}

type InputSelectSelectPrevious struct {
	ServiceBase
	ServiceData InputSelectSelectPreviousParams `json:"service_data,omitempty"`
}
type InputSelectSelectPreviousParams struct{}

func (i *InputSelectSelectPrevious) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectPrevious) SetID(id *int) {
	i.Id = id
}

// NewInputSelectSetOptions creates the object that can be sent to Home Assistant for domain input_select, service set_options
// "Set the options of an input select entity."
func NewInputSelectSetOptions(target Target, inputSelectSetOptionsParams InputSelectSetOptionsParams) *InputSelectSetOptions {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "set_options"
	i := &InputSelectSetOptions{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: inputSelectSetOptionsParams,
	}
	return i
}

type InputSelectSetOptions struct {
	ServiceBase
	ServiceData InputSelectSetOptionsParams `json:"service_data,omitempty"`
}
type InputSelectSetOptionsParams struct{}

func (i *InputSelectSetOptions) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputSelectSetOptions) SetID(id *int) {
	i.Id = id
}
