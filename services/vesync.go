package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewVesyncUpdateDevices creates the object that can be sent to Home Assistant for domain vesync, service update_devices
// "Add new VeSync devices to Home Assistant"
func NewVesyncUpdateDevices(entities []string) *VesyncUpdateDevices {
	serviceDomain := "vesync"
	serviceType := "call_service"
	serviceService := "update_devices"
	v := &VesyncUpdateDevices{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return v
}

type VesyncUpdateDevices struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (v *VesyncUpdateDevices) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VesyncUpdateDevices) SetID(id *int) {
	v.Id = id
}
