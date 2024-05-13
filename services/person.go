package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewPersonReload creates the object that can be sent to Home Assistant for domain person, service reload
// "Reload the person configuration."
func NewPersonReload(target Target) *PersonReload {
	serviceDomain := "person"
	serviceType := "call_service"
	serviceService := "reload"
	p := &PersonReload{
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
	return p
}

type PersonReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (p *PersonReload) JSON() string {
	data, _ := gojson.Marshal(p)
	return string(data)
}
func (p *PersonReload) Targets() []string {
	return p.Target.EntityId
}
func (p *PersonReload) Name() string {
	return fmt.Sprintf("%s.%s", *p.Domain, *p.Service)
}
