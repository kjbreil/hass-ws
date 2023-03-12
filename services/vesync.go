package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewVesyncUpdateDevices creates the object that can be sent to Home Assistant for domain vesync, service update_devices
// "Add new VeSync devices to Home Assistant"
func NewVesyncUpdateDevices(target Target) *VesyncUpdateDevices {
	serviceDomain := "vesync"
	serviceType := "call_service"
	serviceService := "update_devices"
	v := &VesyncUpdateDevices{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return v
}

type VesyncUpdateDevices struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (v *VesyncUpdateDevices) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VesyncUpdateDevices) Name() string {
	return fmt.Sprintf("%s.%s", *v.Domain, *v.Service)
}
func (v *VesyncUpdateDevices) SetID(id *int) {
	v.Id = id
}
