package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewInputSelectReload creates the object that can be sent to Home Assistant for domain input_select, service reload
// "Reloads helpers from the YAML-configuration."
func NewInputSelectReload(target Target) *InputSelectReload {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputSelectReload{
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
	return i
}

type InputSelectReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputSelectReload) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputSelectReload) Targets() []string {
	return i.Target.EntityId
}
func (i *InputSelectReload) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}

// NewInputSelectSelectFirst creates the object that can be sent to Home Assistant for domain input_select, service select_first
// "Selects the first option."
func NewInputSelectSelectFirst(target Target) *InputSelectSelectFirst {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_first"
	i := &InputSelectSelectFirst{
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
	return i
}

type InputSelectSelectFirst struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputSelectSelectFirst) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectFirst) Targets() []string {
	return i.Target.EntityId
}
func (i *InputSelectSelectFirst) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}

// NewInputSelectSelectLast creates the object that can be sent to Home Assistant for domain input_select, service select_last
// "Selects the last option."
func NewInputSelectSelectLast(target Target) *InputSelectSelectLast {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_last"
	i := &InputSelectSelectLast{
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
	return i
}

type InputSelectSelectLast struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputSelectSelectLast) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectLast) Targets() []string {
	return i.Target.EntityId
}
func (i *InputSelectSelectLast) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}

// NewInputSelectSelectNext creates the object that can be sent to Home Assistant for domain input_select, service select_next
// "Selects the next option."
func NewInputSelectSelectNext(target Target) *InputSelectSelectNext {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_next"
	i := &InputSelectSelectNext{
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
	return i
}

type InputSelectSelectNext struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputSelectSelectNext) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectNext) Targets() []string {
	return i.Target.EntityId
}
func (i *InputSelectSelectNext) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}

// NewInputSelectSelectOption creates the object that can be sent to Home Assistant for domain input_select, service select_option
// "Selects an option."
func NewInputSelectSelectOption(target Target) *InputSelectSelectOption {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_option"
	i := &InputSelectSelectOption{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: InputSelectSelectOptionParams{},
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

func (i *InputSelectSelectOption) Option(option string) *InputSelectSelectOption {
	i.ServiceData.Option = &option
	return i
}
func (i *InputSelectSelectOption) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectOption) Targets() []string {
	return i.Target.EntityId
}
func (i *InputSelectSelectOption) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}

// NewInputSelectSelectPrevious creates the object that can be sent to Home Assistant for domain input_select, service select_previous
// "Selects the previous option."
func NewInputSelectSelectPrevious(target Target) *InputSelectSelectPrevious {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "select_previous"
	i := &InputSelectSelectPrevious{
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
	return i
}

type InputSelectSelectPrevious struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (i *InputSelectSelectPrevious) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputSelectSelectPrevious) Targets() []string {
	return i.Target.EntityId
}
func (i *InputSelectSelectPrevious) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}

// NewInputSelectSetOptions creates the object that can be sent to Home Assistant for domain input_select, service set_options
// "Sets the options."
func NewInputSelectSetOptions(target Target) *InputSelectSetOptions {
	serviceDomain := "input_select"
	serviceType := "call_service"
	serviceService := "set_options"
	i := &InputSelectSetOptions{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: InputSelectSetOptionsParams{},
	}
	return i
}

type InputSelectSetOptions struct {
	ServiceBase
	ServiceData InputSelectSetOptionsParams `json:"service_data,omitempty"`
}
type InputSelectSetOptionsParams struct {
	Options *string `json:"options,omitempty"`
}

func (i *InputSelectSetOptions) Options(options string) *InputSelectSetOptions {
	i.ServiceData.Options = &options
	return i
}
func (i *InputSelectSetOptions) JSON() string {
	data, _ := gojson.Marshal(i)
	return string(data)
}
func (i *InputSelectSetOptions) Targets() []string {
	return i.Target.EntityId
}
func (i *InputSelectSetOptions) Name() string {
	return fmt.Sprintf("%s.%s", *i.Domain, *i.Service)
}
