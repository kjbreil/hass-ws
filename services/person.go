package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewPersonReload creates the object that can be sent to Home Assistant for domain person, service reload
// "Reload the person configuration."
func NewPersonReload(entities []string) *PersonReload {
	serviceDomain := "person"
	serviceType := "call_service"
	serviceService := "reload"
	p := &PersonReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return p
}

type PersonReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (p *PersonReload) JSON() string {
	data, _ := json.Marshal(p)
	return string(data)
}
func (p *PersonReload) SetID(id *int) {
	p.Id = id
}
