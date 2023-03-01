package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewZoneReload creates the object that can be sent to Home Assistant for domain zone, service reload
// "Reload the YAML-based zone configuration."
func NewZoneReload(entities []string) *ZoneReload {
	serviceDomain := "zone"
	serviceType := "call_service"
	serviceService := "reload"
	z := &ZoneReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return z
}

type ZoneReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (z *ZoneReload) JSON() string {
	data, _ := json.Marshal(z)
	return string(data)
}
func (z *ZoneReload) SetID(id *int) {
	z.Id = id
}
