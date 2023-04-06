package model

import (
	"strings"
	"time"
)

type Result struct {
	EntityId    *string                `json:"entity_id,omitempty"`
	EntityState *string                `json:"state,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
	LastChanged *time.Time             `json:"last_changed,omitempty"`
	LastUpdated *time.Time             `json:"last_updated,omitempty"`
	Context     *Context               `json:"context,omitempty"`

	AreaId         *string     `json:"area_id"`
	ConfigEntryId  *string     `json:"config_entry_id"`
	DeviceId       *string     `json:"device_id"`
	DisabledBy     *string     `json:"disabled_by"`
	EntityCategory *string     `json:"entity_category"`
	HasEntityName  bool        `json:"has_entity_name"`
	HiddenBy       interface{} `json:"hidden_by"`
	Icon           *string     `json:"icon"`
	Id             string      `json:"id"`
	Name           *string     `json:"name"`
	Options        struct {
		SensorPrivate struct {
			SuggestedUnitOfMeasurement string `json:"suggested_unit_of_measurement"`
		} `json:"sensor.private,omitempty"`
	} `json:"options"`
	OriginalName   *string `json:"original_name"`
	Platform       string  `json:"platform"`
	TranslationKey *string `json:"translation_key"`
	UniqueId       string  `json:"unique_id"`

	ConfigurationUrl *string         `json:"configuration_url"`
	ConfigEntries    []string        `json:"config_entries"`
	Connections      [][]string      `json:"connections"`
	EntryType        *string         `json:"entry_type"`
	HwVersion        *string         `json:"hw_version"`
	Identifiers      [][]interface{} `json:"identifiers"`
	Manufacturer     *string         `json:"manufacturer"`
	Model            *string         `json:"model"`
	NameByUser       *string         `json:"name_by_user"`
	SwVersion        *string         `json:"sw_version"`
	ViaDeviceId      *string         `json:"via_device_id"`
}

func (r *Result) Domain() string {
	if r.EntityId != nil {
		return strings.Split(*r.EntityId, ".")[0]
	}
	return ""
}

func (r *Result) EntityID() string {
	if r.EntityId != nil {
		return strings.Join(strings.Split(*r.EntityId, ".")[1:], "")
	}

	return ""
}
func (r *Result) DomainEntity() string {
	if r.EntityId != nil {
		return *r.EntityId
	}

	return ""
}
func (r *Result) State() string {
	if r.EntityState != nil {
		return *r.EntityState
	}
	return ""
}
