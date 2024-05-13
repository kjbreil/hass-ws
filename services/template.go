package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewTemplateReload creates the object that can be sent to Home Assistant for domain template, service reload
// "Reload all template entities."
func NewTemplateReload(target Target) *TemplateReload {
	serviceDomain := "template"
	serviceType := "call_service"
	serviceService := "reload"
	t := &TemplateReload{
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
	return t
}

type TemplateReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (t *TemplateReload) JSON() string {
	data, _ := gojson.Marshal(t)
	return string(data)
}
func (t *TemplateReload) Targets() []string {
	return t.Target.EntityId
}
func (t *TemplateReload) Name() string {
	return fmt.Sprintf("%s.%s", *t.Domain, *t.Service)
}
