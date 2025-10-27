package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestHomeassistantCheckConfig_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *HomeassistantCheckConfig
		want   string
	}{{
		fields: NewHomeassistantCheckConfig(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"homeassistant\",\"service\":\"check_config\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestHomeassistantReloadConfigEntry_JSON(t *testing.T) {
	entryId := "data"

	tests := []struct {
		name   string
		fields *HomeassistantReloadConfigEntry
		want   string
	}{{
		fields: NewHomeassistantReloadConfigEntry(Targets("climate.kitchen")).EntryId(entryId),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"homeassistant\",\"service\":\"reload_config_entry\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"entry_id\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestHomeassistantReloadCoreConfig_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *HomeassistantReloadCoreConfig
		want   string
	}{{
		fields: NewHomeassistantReloadCoreConfig(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"homeassistant\",\"service\":\"reload_core_config\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestHomeassistantRestart_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *HomeassistantRestart
		want   string
	}{{
		fields: NewHomeassistantRestart(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"homeassistant\",\"service\":\"restart\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestHomeassistantSavePersistentStates_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *HomeassistantSavePersistentStates
		want   string
	}{{
		fields: NewHomeassistantSavePersistentStates(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"homeassistant\",\"service\":\"save_persistent_states\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestHomeassistantSetLocation_JSON(t *testing.T) {
	latitude := "data"
	longitude := "data"

	tests := []struct {
		name   string
		fields *HomeassistantSetLocation
		want   string
	}{{
		fields: NewHomeassistantSetLocation(Targets("climate.kitchen")).Latitude(latitude).Longitude(longitude),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"homeassistant\",\"service\":\"set_location\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"latitude\":\"data\",\"longitude\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestHomeassistantStop_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *HomeassistantStop
		want   string
	}{{
		fields: NewHomeassistantStop(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"homeassistant\",\"service\":\"stop\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestHomeassistantToggle_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *HomeassistantToggle
		want   string
	}{{
		fields: NewHomeassistantToggle(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"homeassistant\",\"service\":\"toggle\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestHomeassistantTurnOff_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *HomeassistantTurnOff
		want   string
	}{{
		fields: NewHomeassistantTurnOff(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"homeassistant\",\"service\":\"turn_off\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestHomeassistantTurnOn_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *HomeassistantTurnOn
		want   string
	}{{
		fields: NewHomeassistantTurnOn(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"homeassistant\",\"service\":\"turn_on\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestHomeassistantUpdateEntity_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *HomeassistantUpdateEntity
		want   string
	}{{
		fields: NewHomeassistantUpdateEntity(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"homeassistant\",\"service\":\"update_entity\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
