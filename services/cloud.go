package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCloudRemoteConnect creates the object that can be sent to Home Assistant for domain cloud, service remote_connect
// "Make instance UI available outside over NabuCasa cloud"
func NewCloudRemoteConnect(entities []string) *CloudRemoteConnect {
	serviceDomain := "cloud"
	serviceType := "call_service"
	serviceService := "remote_connect"
	c := &CloudRemoteConnect{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CloudRemoteConnect struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewCloudRemoteDisconnect(entities []string) *CloudRemoteDisconnect {
	serviceDomain := "cloud"
	serviceType := "call_service"
	serviceService := "remote_disconnect"
	c := &CloudRemoteDisconnect{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CloudRemoteDisconnect struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CloudRemoteDisconnect) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CloudRemoteDisconnect) SetID(id *int) {
	c.Id = id
}
