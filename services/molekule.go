package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewMolekuleUpdateDevices creates the object that can be sent to Home Assistant for domain molekule, service update_devices
// "Add new Molekule devices to Home Assistant"
func NewMolekuleUpdateDevices(target Target, molekuleUpdateDevicesParams MolekuleUpdateDevicesParams) *MolekuleUpdateDevices {
	serviceDomain := "molekule"
	serviceType := "call_service"
	serviceService := "update_devices"
	m := &MolekuleUpdateDevices{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: molekuleUpdateDevicesParams,
	}
	return m
}

type MolekuleUpdateDevices struct {
	ServiceBase
	ServiceData MolekuleUpdateDevicesParams `json:"service_data,omitempty"`
}
type MolekuleUpdateDevicesParams struct{}

func (m *MolekuleUpdateDevices) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MolekuleUpdateDevices) SetID(id *int) {
	m.Id = id
}
