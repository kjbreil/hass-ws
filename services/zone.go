package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewZoneReload creates the object that can be sent to Home Assistant for domain zone, service reload
// "Reload the YAML-based zone configuration."
func NewZoneReload(target Target) *ZoneReload {
	serviceDomain := "zone"
	serviceType := "call_service"
	serviceService := "reload"
	z := &ZoneReload{
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
	return z
}

type ZoneReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (z *ZoneReload) JSON() string {
	data, _ := gojson.Marshal(z)
	return string(data)
}
func (z *ZoneReload) Targets() []string {
	return z.Target.EntityId
}
func (z *ZoneReload) Name() string {
	return fmt.Sprintf("%s.%s", *z.Domain, *z.Service)
}
