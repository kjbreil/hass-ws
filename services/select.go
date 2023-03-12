package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSelectSelectOption creates the object that can be sent to Home Assistant for domain select, service select_option
// "Select an option of an select entity."
func NewSelectSelectOption(target Target, selectSelectOptionParams *SelectSelectOptionParams) *SelectSelectOption {
	serviceDomain := "select"
	serviceType := "call_service"
	serviceService := "select_option"
	s := &SelectSelectOption{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *selectSelectOptionParams,
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

func (s *SelectSelectOption) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SelectSelectOption) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *SelectSelectOption) SetID(id *int) {
	s.Id = id
}
