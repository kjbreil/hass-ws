package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewMolekuleUpdateDevices creates the object that can be sent to Home Assistant for domain molekule, service update_devices
// "Add new Molekule devices to Home Assistant"
func NewMolekuleUpdateDevices(entities []string) *MolekuleUpdateDevices {
	serviceDomain := "molekule"
	serviceType := "call_service"
	serviceService := "update_devices"
	m := &MolekuleUpdateDevices{
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

type MolekuleUpdateDevices struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MolekuleUpdateDevices) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MolekuleUpdateDevices) SetID(id *int) {
	m.Id = id
}
