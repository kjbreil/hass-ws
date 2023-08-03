package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewThermalComfortReload creates the object that can be sent to Home Assistant for domain thermal_comfort, service reload
// "Reload all Thermal Comfort entities."
func NewThermalComfortReload(target Target) *ThermalComfortReload {
	serviceDomain := "thermal_comfort"
	serviceType := "call_service"
	serviceService := "reload"
	t := &ThermalComfortReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return t
}

type ThermalComfortReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (t *ThermalComfortReload) JSON() string {
	data, _ := gojson.Marshal(t)
	return string(data)
}
func (t *ThermalComfortReload) Targets() []string {
	return t.Target.EntityId
}
func (t *ThermalComfortReload) Name() string {
	return fmt.Sprintf("%s.%s", *t.Domain, *t.Service)
}
func (t *ThermalComfortReload) SetID(id *int) {
	t.Id = id
}
