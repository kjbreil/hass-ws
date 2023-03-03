package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewHomeassistantCheckConfig creates the object that can be sent to Home Assistant for domain homeassistant, service check_config
// "Check the Home Assistant configuration files for errors. Errors will be displayed in the Home Assistant log."
func NewHomeassistantCheckConfig(target Target, homeassistantCheckConfigParams HomeassistantCheckConfigParams) *HomeassistantCheckConfig {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "check_config"
	h := &HomeassistantCheckConfig{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homeassistantCheckConfigParams,
	}
	return h
}

type HomeassistantCheckConfig struct {
	ServiceBase
	ServiceData HomeassistantCheckConfigParams `json:"service_data,omitempty"`
}
type HomeassistantCheckConfigParams struct{}

func (h *HomeassistantCheckConfig) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantCheckConfig) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantReloadConfigEntry creates the object that can be sent to Home Assistant for domain homeassistant, service reload_config_entry
// "Reload a config entry that matches a target."
func NewHomeassistantReloadConfigEntry(target Target, homeassistantReloadConfigEntryParams HomeassistantReloadConfigEntryParams) *HomeassistantReloadConfigEntry {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "reload_config_entry"
	h := &HomeassistantReloadConfigEntry{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homeassistantReloadConfigEntryParams,
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

func (h *HomeassistantReloadConfigEntry) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantReloadConfigEntry) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantReloadCoreConfig creates the object that can be sent to Home Assistant for domain homeassistant, service reload_core_config
// "Reload the core configuration."
func NewHomeassistantReloadCoreConfig(target Target, homeassistantReloadCoreConfigParams HomeassistantReloadCoreConfigParams) *HomeassistantReloadCoreConfig {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "reload_core_config"
	h := &HomeassistantReloadCoreConfig{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homeassistantReloadCoreConfigParams,
	}
	return h
}

type HomeassistantReloadCoreConfig struct {
	ServiceBase
	ServiceData HomeassistantReloadCoreConfigParams `json:"service_data,omitempty"`
}
type HomeassistantReloadCoreConfigParams struct{}

func (h *HomeassistantReloadCoreConfig) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantReloadCoreConfig) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantRestart creates the object that can be sent to Home Assistant for domain homeassistant, service restart
// "Restart the Home Assistant service."
func NewHomeassistantRestart(target Target, homeassistantRestartParams HomeassistantRestartParams) *HomeassistantRestart {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "restart"
	h := &HomeassistantRestart{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homeassistantRestartParams,
	}
	return h
}

type HomeassistantRestart struct {
	ServiceBase
	ServiceData HomeassistantRestartParams `json:"service_data,omitempty"`
}
type HomeassistantRestartParams struct{}

func (h *HomeassistantRestart) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantRestart) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantSavePersistentStates creates the object that can be sent to Home Assistant for domain homeassistant, service save_persistent_states
// "Save the persistent states (for entities derived from RestoreEntity) immediately. Maintain the normal periodic saving interval."
func NewHomeassistantSavePersistentStates(target Target, homeassistantSavePersistentStatesParams HomeassistantSavePersistentStatesParams) *HomeassistantSavePersistentStates {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "save_persistent_states"
	h := &HomeassistantSavePersistentStates{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homeassistantSavePersistentStatesParams,
	}
	return h
}

type HomeassistantSavePersistentStates struct {
	ServiceBase
	ServiceData HomeassistantSavePersistentStatesParams `json:"service_data,omitempty"`
}
type HomeassistantSavePersistentStatesParams struct{}

func (h *HomeassistantSavePersistentStates) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantSavePersistentStates) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantSetLocation creates the object that can be sent to Home Assistant for domain homeassistant, service set_location
// "Update the Home Assistant location."
func NewHomeassistantSetLocation(target Target, homeassistantSetLocationParams HomeassistantSetLocationParams) *HomeassistantSetLocation {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "set_location"
	h := &HomeassistantSetLocation{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homeassistantSetLocationParams,
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

func (h *HomeassistantSetLocation) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantSetLocation) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantStop creates the object that can be sent to Home Assistant for domain homeassistant, service stop
// "Stop the Home Assistant service."
func NewHomeassistantStop(target Target, homeassistantStopParams HomeassistantStopParams) *HomeassistantStop {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "stop"
	h := &HomeassistantStop{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homeassistantStopParams,
	}
	return h
}

type HomeassistantStop struct {
	ServiceBase
	ServiceData HomeassistantStopParams `json:"service_data,omitempty"`
}
type HomeassistantStopParams struct{}

func (h *HomeassistantStop) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantStop) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantToggle creates the object that can be sent to Home Assistant for domain homeassistant, service toggle
// "Generic service to toggle devices on/off under any domain"
func NewHomeassistantToggle(target Target, homeassistantToggleParams HomeassistantToggleParams) *HomeassistantToggle {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "toggle"
	h := &HomeassistantToggle{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homeassistantToggleParams,
	}
	return h
}

type HomeassistantToggle struct {
	ServiceBase
	ServiceData HomeassistantToggleParams `json:"service_data,omitempty"`
}
type HomeassistantToggleParams struct{}

func (h *HomeassistantToggle) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantToggle) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantTurnOff creates the object that can be sent to Home Assistant for domain homeassistant, service turn_off
// "Generic service to turn devices off under any domain."
func NewHomeassistantTurnOff(target Target, homeassistantTurnOffParams HomeassistantTurnOffParams) *HomeassistantTurnOff {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "turn_off"
	h := &HomeassistantTurnOff{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homeassistantTurnOffParams,
	}
	return h
}

type HomeassistantTurnOff struct {
	ServiceBase
	ServiceData HomeassistantTurnOffParams `json:"service_data,omitempty"`
}
type HomeassistantTurnOffParams struct{}

func (h *HomeassistantTurnOff) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantTurnOff) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantTurnOn creates the object that can be sent to Home Assistant for domain homeassistant, service turn_on
// "Generic service to turn devices on under any domain."
func NewHomeassistantTurnOn(target Target, homeassistantTurnOnParams HomeassistantTurnOnParams) *HomeassistantTurnOn {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "turn_on"
	h := &HomeassistantTurnOn{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homeassistantTurnOnParams,
	}
	return h
}

type HomeassistantTurnOn struct {
	ServiceBase
	ServiceData HomeassistantTurnOnParams `json:"service_data,omitempty"`
}
type HomeassistantTurnOnParams struct{}

func (h *HomeassistantTurnOn) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantTurnOn) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantUpdateEntity creates the object that can be sent to Home Assistant for domain homeassistant, service update_entity
// "Force one or more entities to update its data"
func NewHomeassistantUpdateEntity(target Target, homeassistantUpdateEntityParams HomeassistantUpdateEntityParams) *HomeassistantUpdateEntity {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "update_entity"
	h := &HomeassistantUpdateEntity{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: homeassistantUpdateEntityParams,
	}
	return h
}

type HomeassistantUpdateEntity struct {
	ServiceBase
	ServiceData HomeassistantUpdateEntityParams `json:"service_data,omitempty"`
}
type HomeassistantUpdateEntityParams struct{}

func (h *HomeassistantUpdateEntity) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantUpdateEntity) SetID(id *int) {
	h.Id = id
}
