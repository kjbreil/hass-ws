package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewHomeassistantCheckConfig creates the object that can be sent to Home Assistant for domain homeassistant, service check_config
// "Check the Home Assistant configuration files for errors. Errors will be displayed in the Home Assistant log."
func NewHomeassistantCheckConfig(target Target) *HomeassistantCheckConfig {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "check_config"
	h := &HomeassistantCheckConfig{
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
	return h
}

type HomeassistantCheckConfig struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (h *HomeassistantCheckConfig) JSON() string {
	data, _ := gojson.Marshal(h)
	return string(data)
}
func (h *HomeassistantCheckConfig) Targets() []string {
	return h.Target.EntityId
}
func (h *HomeassistantCheckConfig) Name() string {
	return fmt.Sprintf("%s.%s", *h.Domain, *h.Service)
}

// NewHomeassistantReloadConfigEntry creates the object that can be sent to Home Assistant for domain homeassistant, service reload_config_entry
// "Reload a config entry that matches a target."
func NewHomeassistantReloadConfigEntry(target Target) *HomeassistantReloadConfigEntry {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "reload_config_entry"
	h := &HomeassistantReloadConfigEntry{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: HomeassistantReloadConfigEntryParams{},
	}
	return h
}

type HomeassistantReloadConfigEntry struct {
	ServiceBase
	ServiceData HomeassistantReloadConfigEntryParams `json:"service_data,omitempty"`
}
type HomeassistantReloadConfigEntryParams struct {
	EntryId *string `json:"entry_id,omitempty"`
}

func (h *HomeassistantReloadConfigEntry) EntryId(entryId string) *HomeassistantReloadConfigEntry {
	h.ServiceData.EntryId = &entryId
	return h
}
func (h *HomeassistantReloadConfigEntry) JSON() string {
	data, _ := gojson.Marshal(h)
	return string(data)
}
func (h *HomeassistantReloadConfigEntry) Targets() []string {
	return h.Target.EntityId
}
func (h *HomeassistantReloadConfigEntry) Name() string {
	return fmt.Sprintf("%s.%s", *h.Domain, *h.Service)
}

// NewHomeassistantReloadCoreConfig creates the object that can be sent to Home Assistant for domain homeassistant, service reload_core_config
// "Reload the core configuration."
func NewHomeassistantReloadCoreConfig(target Target) *HomeassistantReloadCoreConfig {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "reload_core_config"
	h := &HomeassistantReloadCoreConfig{
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
	return h
}

type HomeassistantReloadCoreConfig struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (h *HomeassistantReloadCoreConfig) JSON() string {
	data, _ := gojson.Marshal(h)
	return string(data)
}
func (h *HomeassistantReloadCoreConfig) Targets() []string {
	return h.Target.EntityId
}
func (h *HomeassistantReloadCoreConfig) Name() string {
	return fmt.Sprintf("%s.%s", *h.Domain, *h.Service)
}

// NewHomeassistantRestart creates the object that can be sent to Home Assistant for domain homeassistant, service restart
// "Restart the Home Assistant service."
func NewHomeassistantRestart(target Target) *HomeassistantRestart {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "restart"
	h := &HomeassistantRestart{
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
	return h
}

type HomeassistantRestart struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (h *HomeassistantRestart) JSON() string {
	data, _ := gojson.Marshal(h)
	return string(data)
}
func (h *HomeassistantRestart) Targets() []string {
	return h.Target.EntityId
}
func (h *HomeassistantRestart) Name() string {
	return fmt.Sprintf("%s.%s", *h.Domain, *h.Service)
}

// NewHomeassistantSavePersistentStates creates the object that can be sent to Home Assistant for domain homeassistant, service save_persistent_states
// "Save the persistent states (for entities derived from RestoreEntity) immediately. Maintain the normal periodic saving interval."
func NewHomeassistantSavePersistentStates(target Target) *HomeassistantSavePersistentStates {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "save_persistent_states"
	h := &HomeassistantSavePersistentStates{
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
	return h
}

type HomeassistantSavePersistentStates struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (h *HomeassistantSavePersistentStates) JSON() string {
	data, _ := gojson.Marshal(h)
	return string(data)
}
func (h *HomeassistantSavePersistentStates) Targets() []string {
	return h.Target.EntityId
}
func (h *HomeassistantSavePersistentStates) Name() string {
	return fmt.Sprintf("%s.%s", *h.Domain, *h.Service)
}

// NewHomeassistantSetLocation creates the object that can be sent to Home Assistant for domain homeassistant, service set_location
// "Update the Home Assistant location."
func NewHomeassistantSetLocation(target Target) *HomeassistantSetLocation {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "set_location"
	h := &HomeassistantSetLocation{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: HomeassistantSetLocationParams{},
	}
	return h
}

type HomeassistantSetLocation struct {
	ServiceBase
	ServiceData HomeassistantSetLocationParams `json:"service_data,omitempty"`
}
type HomeassistantSetLocationParams struct {
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
}

func (h *HomeassistantSetLocation) Latitude(latitude string) *HomeassistantSetLocation {
	h.ServiceData.Latitude = &latitude
	return h
}
func (h *HomeassistantSetLocation) Longitude(longitude string) *HomeassistantSetLocation {
	h.ServiceData.Longitude = &longitude
	return h
}
func (h *HomeassistantSetLocation) JSON() string {
	data, _ := gojson.Marshal(h)
	return string(data)
}
func (h *HomeassistantSetLocation) Targets() []string {
	return h.Target.EntityId
}
func (h *HomeassistantSetLocation) Name() string {
	return fmt.Sprintf("%s.%s", *h.Domain, *h.Service)
}

// NewHomeassistantStop creates the object that can be sent to Home Assistant for domain homeassistant, service stop
// "Stop the Home Assistant service."
func NewHomeassistantStop(target Target) *HomeassistantStop {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "stop"
	h := &HomeassistantStop{
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
	return h
}

type HomeassistantStop struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (h *HomeassistantStop) JSON() string {
	data, _ := gojson.Marshal(h)
	return string(data)
}
func (h *HomeassistantStop) Targets() []string {
	return h.Target.EntityId
}
func (h *HomeassistantStop) Name() string {
	return fmt.Sprintf("%s.%s", *h.Domain, *h.Service)
}

// NewHomeassistantToggle creates the object that can be sent to Home Assistant for domain homeassistant, service toggle
// "Generic service to toggle devices on/off under any domain"
func NewHomeassistantToggle(target Target) *HomeassistantToggle {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "toggle"
	h := &HomeassistantToggle{
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
	return h
}

type HomeassistantToggle struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (h *HomeassistantToggle) JSON() string {
	data, _ := gojson.Marshal(h)
	return string(data)
}
func (h *HomeassistantToggle) Targets() []string {
	return h.Target.EntityId
}
func (h *HomeassistantToggle) Name() string {
	return fmt.Sprintf("%s.%s", *h.Domain, *h.Service)
}

// NewHomeassistantTurnOff creates the object that can be sent to Home Assistant for domain homeassistant, service turn_off
// "Generic service to turn devices off under any domain."
func NewHomeassistantTurnOff(target Target) *HomeassistantTurnOff {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "turn_off"
	h := &HomeassistantTurnOff{
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
	return h
}

type HomeassistantTurnOff struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (h *HomeassistantTurnOff) JSON() string {
	data, _ := gojson.Marshal(h)
	return string(data)
}
func (h *HomeassistantTurnOff) Targets() []string {
	return h.Target.EntityId
}
func (h *HomeassistantTurnOff) Name() string {
	return fmt.Sprintf("%s.%s", *h.Domain, *h.Service)
}

// NewHomeassistantTurnOn creates the object that can be sent to Home Assistant for domain homeassistant, service turn_on
// "Generic service to turn devices on under any domain."
func NewHomeassistantTurnOn(target Target) *HomeassistantTurnOn {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "turn_on"
	h := &HomeassistantTurnOn{
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
	return h
}

type HomeassistantTurnOn struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (h *HomeassistantTurnOn) JSON() string {
	data, _ := gojson.Marshal(h)
	return string(data)
}
func (h *HomeassistantTurnOn) Targets() []string {
	return h.Target.EntityId
}
func (h *HomeassistantTurnOn) Name() string {
	return fmt.Sprintf("%s.%s", *h.Domain, *h.Service)
}

// NewHomeassistantUpdateEntity creates the object that can be sent to Home Assistant for domain homeassistant, service update_entity
// "Force one or more entities to update its data"
func NewHomeassistantUpdateEntity(target Target) *HomeassistantUpdateEntity {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "update_entity"
	h := &HomeassistantUpdateEntity{
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
	return h
}

type HomeassistantUpdateEntity struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (h *HomeassistantUpdateEntity) JSON() string {
	data, _ := gojson.Marshal(h)
	return string(data)
}
func (h *HomeassistantUpdateEntity) Targets() []string {
	return h.Target.EntityId
}
func (h *HomeassistantUpdateEntity) Name() string {
	return fmt.Sprintf("%s.%s", *h.Domain, *h.Service)
}
