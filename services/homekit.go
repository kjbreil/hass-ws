package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewHomekitReload creates the object that can be sent to Home Assistant for domain homekit, service reload
// "Reload homekit and re-process YAML configuration"
func NewHomekitReload(target Target, homekitReloadParams HomekitReloadParams) *HomekitReload {
	serviceDomain := "homekit"
	serviceType := "call_service"
	serviceService := "reload"
	h := &HomekitReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homekitReloadParams,
	}
	return h
}

type HomekitReload struct {
	ServiceBase
	ServiceData HomekitReloadParams `json:"service_data,omitempty"`
}
type HomekitReloadParams struct{}

func (h *HomekitReload) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomekitReload) SetID(id *int) {
	h.Id = id
}

// NewHomekitResetAccessory creates the object that can be sent to Home Assistant for domain homekit, service reset_accessory
// "Reset a HomeKit accessory"
func NewHomekitResetAccessory(target Target, homekitResetAccessoryParams HomekitResetAccessoryParams) *HomekitResetAccessory {
	serviceDomain := "homekit"
	serviceType := "call_service"
	serviceService := "reset_accessory"
	h := &HomekitResetAccessory{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homekitResetAccessoryParams,
	}
	return h
}

type HomekitResetAccessory struct {
	ServiceBase
	ServiceData HomekitResetAccessoryParams `json:"service_data,omitempty"`
}
type HomekitResetAccessoryParams struct{}

func (h *HomekitResetAccessory) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomekitResetAccessory) SetID(id *int) {
	h.Id = id
}

// NewHomekitUnpair creates the object that can be sent to Home Assistant for domain homekit, service unpair
// "Forcefully remove all pairings from an accessory to allow re-pairing. Use this service if the accessory is no longer responsive, and you want to avoid deleting and re-adding the entry. Room locations, and accessory preferences will be lost."
func NewHomekitUnpair(target Target, homekitUnpairParams HomekitUnpairParams) *HomekitUnpair {
	serviceDomain := "homekit"
	serviceType := "call_service"
	serviceService := "unpair"
	h := &HomekitUnpair{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homekitUnpairParams,
	}
	return h
}

type HomekitUnpair struct {
	ServiceBase
	ServiceData HomekitUnpairParams `json:"service_data,omitempty"`
}
type HomekitUnpairParams struct{}

func (h *HomekitUnpair) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomekitUnpair) SetID(id *int) {
	h.Id = id
}
