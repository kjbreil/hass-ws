package services

import "encoding/json"

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
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
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
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TemplateReload) SetID(id *int) {
	t.Id = id
}
