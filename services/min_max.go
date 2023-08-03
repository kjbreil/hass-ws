package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewMinMaxReload creates the object that can be sent to Home Assistant for domain min_max, service reload
// "Reload all min_max entities."
func NewMinMaxReload(target Target) *MinMaxReload {
	serviceDomain := "min_max"
	serviceType := "call_service"
	serviceService := "reload"
	m := &MinMaxReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return m
}

type MinMaxReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MinMaxReload) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MinMaxReload) Targets() []string {
	return m.Target.EntityId
}
func (m *MinMaxReload) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}
func (m *MinMaxReload) SetID(id *int) {
	m.Id = id
}
