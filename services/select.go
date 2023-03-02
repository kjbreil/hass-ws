package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSelectSelectOption creates the object that can be sent to Home Assistant for domain select, service select_option
// "Select an option of an select entity."
func NewSelectSelectOption(entities []string, option *string) *SelectSelectOption {
	serviceDomain := "select"
	serviceType := "call_service"
	serviceService := "select_option"
	s := &SelectSelectOption{
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
	return s
}

type SelectSelectOption struct {
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

func (s *SelectSelectOption) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SelectSelectOption) SetID(id *int) {
	s.Id = id
}
