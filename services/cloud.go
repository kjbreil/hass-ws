package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCloudRemoteConnect creates the object that can be sent to Home Assistant for domain cloud, service remote_connect
// "Makes the instance UI accessible from outside of the local network by enabling your Home Assistant Cloud connection."
func NewCloudRemoteConnect(target Target) *CloudRemoteConnect {
	serviceDomain := "cloud"
	serviceType := "call_service"
	serviceService := "remote_connect"
	c := &CloudRemoteConnect{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CloudRemoteConnect) Targets() []string {
	return c.Target.EntityId
}
func (c *CloudRemoteConnect) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCloudRemoteDisconnect creates the object that can be sent to Home Assistant for domain cloud, service remote_disconnect
// "Disconnects the instance UI from Home Assistant Cloud. This disables access to it from outside your local network."
func NewCloudRemoteDisconnect(target Target) *CloudRemoteDisconnect {
	serviceDomain := "cloud"
	serviceType := "call_service"
	serviceService := "remote_disconnect"
	c := &CloudRemoteDisconnect{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CloudRemoteDisconnect) Targets() []string {
	return c.Target.EntityId
}
func (c *CloudRemoteDisconnect) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}
