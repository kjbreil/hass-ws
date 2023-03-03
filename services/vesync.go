package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewVesyncUpdateDevices creates the object that can be sent to Home Assistant for domain vesync, service update_devices
// "Add new VeSync devices to Home Assistant"
func NewVesyncUpdateDevices(target Target, vesyncUpdateDevicesParams VesyncUpdateDevicesParams) *VesyncUpdateDevices {
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
		ServiceData: vesyncUpdateDevicesParams,
	}
	return v
}

type VesyncUpdateDevices struct {
	ServiceBase
	ServiceData VesyncUpdateDevicesParams `json:"service_data,omitempty"`
}
type VesyncUpdateDevicesParams struct{}

func (v *VesyncUpdateDevices) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VesyncUpdateDevices) SetID(id *int) {
	v.Id = id
}
