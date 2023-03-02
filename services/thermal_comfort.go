package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewThermalComfortReload creates the object that can be sent to Home Assistant for domain thermal_comfort, service reload
// "Reload all Thermal Comfort entities."
func NewThermalComfortReload(entities []string) *ThermalComfortReload {
	serviceDomain := "thermal_comfort"
	serviceType := "call_service"
	serviceService := "reload"
	t := &ThermalComfortReload{
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

type ThermalComfortReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (t *ThermalComfortReload) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *ThermalComfortReload) SetID(id *int) {
	t.Id = id
}
