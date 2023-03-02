package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewHomekitReload creates the object that can be sent to Home Assistant for domain homekit, service reload
// "Reload homekit and re-process YAML configuration"
func NewHomekitReload(entities []string) *HomekitReload {
	serviceDomain := "homekit"
	serviceType := "call_service"
	serviceService := "reload"
	h := &HomekitReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return h
}

type HomekitReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HomekitReload) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomekitReload) SetID(id *int) {
	h.Id = id
}

// NewHomekitResetAccessory creates the object that can be sent to Home Assistant for domain homekit, service reset_accessory
// "Reset a HomeKit accessory"
func NewHomekitResetAccessory(entities []string) *HomekitResetAccessory {
	serviceDomain := "homekit"
	serviceType := "call_service"
	serviceService := "reset_accessory"
	h := &HomekitResetAccessory{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return h
}

type HomekitResetAccessory struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HomekitResetAccessory) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomekitResetAccessory) SetID(id *int) {
	h.Id = id
}

// NewHomekitUnpair creates the object that can be sent to Home Assistant for domain homekit, service unpair
// "Forcefully remove all pairings from an accessory to allow re-pairing. Use this service if the accessory is no longer responsive, and you want to avoid deleting and re-adding the entry. Room locations, and accessory preferences will be lost."
func NewHomekitUnpair(entities []string) *HomekitUnpair {
	serviceDomain := "homekit"
	serviceType := "call_service"
	serviceService := "unpair"
	h := &HomekitUnpair{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return h
}

type HomekitUnpair struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HomekitUnpair) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomekitUnpair) SetID(id *int) {
	h.Id = id
}
