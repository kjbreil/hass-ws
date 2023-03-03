package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCloudRemoteConnect creates the object that can be sent to Home Assistant for domain cloud, service remote_connect
// "Make instance UI available outside over NabuCasa cloud"
func NewCloudRemoteConnect(target Target) *CloudRemoteConnect {
	serviceDomain := "cloud"
	serviceType := "call_service"
	serviceService := "remote_connect"
	c := &CloudRemoteConnect{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type CloudRemoteConnect struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CloudRemoteConnect) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CloudRemoteConnect) SetID(id *int) {
	c.Id = id
}

// NewCloudRemoteDisconnect creates the object that can be sent to Home Assistant for domain cloud, service remote_disconnect
// "Disconnect UI from NabuCasa cloud"
func NewCloudRemoteDisconnect(target Target) *CloudRemoteDisconnect {
	serviceDomain := "cloud"
	serviceType := "call_service"
	serviceService := "remote_disconnect"
	c := &CloudRemoteDisconnect{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type CloudRemoteDisconnect struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CloudRemoteDisconnect) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CloudRemoteDisconnect) SetID(id *int) {
	c.Id = id
}
