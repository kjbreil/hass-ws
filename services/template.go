package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewTemplateReload creates the object that can be sent to Home Assistant for domain template, service reload
// "Reload all template entities."
func NewTemplateReload(entities []string) *TemplateReload {
	serviceDomain := "template"
	serviceType := "call_service"
	serviceService := "reload"
	t := &TemplateReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return t
}

type TemplateReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (t *TemplateReload) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TemplateReload) SetID(id *int) {
	t.Id = id
}
