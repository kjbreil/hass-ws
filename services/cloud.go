package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCloudRemoteConnect creates the object that can be sent to Home Assistant for domain cloud, service remote_connect
// "Make instance UI available outside over NabuCasa cloud"
func NewCloudRemoteConnect(target Target, cloudRemoteConnectParams CloudRemoteConnectParams) *CloudRemoteConnect {
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
		ServiceData: cloudRemoteConnectParams,
	}
	return c
}

type CloudRemoteConnect struct {
	ServiceBase
	ServiceData CloudRemoteConnectParams `json:"service_data,omitempty"`
}
type CloudRemoteConnectParams struct{}

func (c *CloudRemoteConnect) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CloudRemoteConnect) SetID(id *int) {
	c.Id = id
}

// NewCloudRemoteDisconnect creates the object that can be sent to Home Assistant for domain cloud, service remote_disconnect
// "Disconnect UI from NabuCasa cloud"
func NewCloudRemoteDisconnect(target Target, cloudRemoteDisconnectParams CloudRemoteDisconnectParams) *CloudRemoteDisconnect {
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
		ServiceData: cloudRemoteDisconnectParams,
	}
	return c
}

type CloudRemoteDisconnect struct {
	ServiceBase
	ServiceData CloudRemoteDisconnectParams `json:"service_data,omitempty"`
}
type CloudRemoteDisconnectParams struct{}

func (c *CloudRemoteDisconnect) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CloudRemoteDisconnect) SetID(id *int) {
	c.Id = id
}
