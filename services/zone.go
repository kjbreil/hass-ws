package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewZoneReload creates the object that can be sent to Home Assistant for domain zone, service reload
// "Reload the YAML-based zone configuration."
func NewZoneReload(target Target, zoneReloadParams ZoneReloadParams) *ZoneReload {
	serviceDomain := "zone"
	serviceType := "call_service"
	serviceService := "reload"
	z := &ZoneReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: zoneReloadParams,
	}
	return z
}

type ZoneReload struct {
	ServiceBase
	ServiceData ZoneReloadParams `json:"service_data,omitempty"`
}
type ZoneReloadParams struct{}

func (z *ZoneReload) JSON() string {
	data, _ := json.Marshal(z)
	return string(data)
}
func (z *ZoneReload) SetID(id *int) {
	z.Id = id
}
