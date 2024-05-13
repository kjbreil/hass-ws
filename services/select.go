package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSelectSelectOption creates the object that can be sent to Home Assistant for domain select, service select_option
// "Select an option of an select entity."
func NewSelectSelectOption(target Target) *SelectSelectOption {
	serviceDomain := "select"
	serviceType := "call_service"
	serviceService := "select_option"
	s := &SelectSelectOption{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: SelectSelectOptionParams{},
	}
	return s
}

type SelectSelectOption struct {
	ServiceBase
	ServiceData SelectSelectOptionParams `json:"service_data,omitempty"`
}
type SelectSelectOptionParams struct {
	Option *string `json:"option,omitempty"`
}

func (s *SelectSelectOption) Option(option string) *SelectSelectOption {
	s.ServiceData.Option = &option
	return s
}
func (s *SelectSelectOption) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SelectSelectOption) Targets() []string {
	return s.Target.EntityId
}
func (s *SelectSelectOption) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
