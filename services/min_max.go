package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewMinMaxReload creates the object that can be sent to Home Assistant for domain min_max, service reload
// "Reload all min_max entities."
func NewMinMaxReload(entities []string) *MinMaxReload {
	serviceDomain := "min_max"
	serviceType := "call_service"
	serviceService := "reload"
	m := &MinMaxReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return m
}

type MinMaxReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MinMaxReload) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MinMaxReload) SetID(id *int) {
	m.Id = id
}
