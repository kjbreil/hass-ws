package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewHomeassistantCheckConfig creates the object that can be sent to Home Assistant for domain homeassistant, service check_config
// "Check the Home Assistant configuration files for errors. Errors will be displayed in the Home Assistant log."
func NewHomeassistantCheckConfig(entities []string) *HomeassistantCheckConfig {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "check_config"
	h := &HomeassistantCheckConfig{
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

type HomeassistantCheckConfig struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HomeassistantCheckConfig) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantCheckConfig) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantReloadConfigEntry creates the object that can be sent to Home Assistant for domain homeassistant, service reload_config_entry
// "Reload a config entry that matches a target."
func NewHomeassistantReloadConfigEntry(entities []string, entryId *string) *HomeassistantReloadConfigEntry {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "reload_config_entry"
	h := &HomeassistantReloadConfigEntry{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			EntryId *string `json:"entry_id,omitempty"`
		}{EntryId: entryId},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return h
}

type HomeassistantReloadConfigEntry struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		EntryId *string `json:"entry_id,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewHomeassistantReloadCoreConfig(entities []string) *HomeassistantReloadCoreConfig {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "reload_core_config"
	h := &HomeassistantReloadCoreConfig{
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

type HomeassistantReloadCoreConfig struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HomeassistantReloadCoreConfig) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantReloadCoreConfig) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantRestart creates the object that can be sent to Home Assistant for domain homeassistant, service restart
// "Restart the Home Assistant service."
func NewHomeassistantRestart(entities []string) *HomeassistantRestart {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "restart"
	h := &HomeassistantRestart{
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

type HomeassistantRestart struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HomeassistantRestart) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantRestart) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantSavePersistentStates creates the object that can be sent to Home Assistant for domain homeassistant, service save_persistent_states
// "Save the persistent states (for entities derived from RestoreEntity) immediately. Maintain the normal periodic saving interval."
func NewHomeassistantSavePersistentStates(entities []string) *HomeassistantSavePersistentStates {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "save_persistent_states"
	h := &HomeassistantSavePersistentStates{
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

type HomeassistantSavePersistentStates struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HomeassistantSavePersistentStates) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantSavePersistentStates) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantSetLocation creates the object that can be sent to Home Assistant for domain homeassistant, service set_location
// "Update the Home Assistant location."
func NewHomeassistantSetLocation(entities []string, latitude *string, longitude *string) *HomeassistantSetLocation {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "set_location"
	h := &HomeassistantSetLocation{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Latitude  *string `json:"latitude,omitempty"`
			Longitude *string `json:"longitude,omitempty"`
		}{
			Latitude:  latitude,
			Longitude: longitude,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return h
}

type HomeassistantSetLocation struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Latitude  *string `json:"latitude,omitempty"`
		Longitude *string `json:"longitude,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewHomeassistantStop(entities []string) *HomeassistantStop {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "stop"
	h := &HomeassistantStop{
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

type HomeassistantStop struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HomeassistantStop) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantStop) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantToggle creates the object that can be sent to Home Assistant for domain homeassistant, service toggle
// "Generic service to toggle devices on/off under any domain"
func NewHomeassistantToggle(entities []string) *HomeassistantToggle {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "toggle"
	h := &HomeassistantToggle{
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

type HomeassistantToggle struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HomeassistantToggle) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantToggle) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantTurnOff creates the object that can be sent to Home Assistant for domain homeassistant, service turn_off
// "Generic service to turn devices off under any domain."
func NewHomeassistantTurnOff(entities []string) *HomeassistantTurnOff {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "turn_off"
	h := &HomeassistantTurnOff{
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

type HomeassistantTurnOff struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HomeassistantTurnOff) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantTurnOff) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantTurnOn creates the object that can be sent to Home Assistant for domain homeassistant, service turn_on
// "Generic service to turn devices on under any domain."
func NewHomeassistantTurnOn(entities []string) *HomeassistantTurnOn {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "turn_on"
	h := &HomeassistantTurnOn{
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

type HomeassistantTurnOn struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HomeassistantTurnOn) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantTurnOn) SetID(id *int) {
	h.Id = id
}

// NewHomeassistantUpdateEntity creates the object that can be sent to Home Assistant for domain homeassistant, service update_entity
// "Force one or more entities to update its data"
func NewHomeassistantUpdateEntity(entities []string) *HomeassistantUpdateEntity {
	serviceDomain := "homeassistant"
	serviceType := "call_service"
	serviceService := "update_entity"
	h := &HomeassistantUpdateEntity{
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

type HomeassistantUpdateEntity struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HomeassistantUpdateEntity) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HomeassistantUpdateEntity) SetID(id *int) {
	h.Id = id
}
