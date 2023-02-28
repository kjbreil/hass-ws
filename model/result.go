package model

import "time"

type Result struct {
	EntityId    *string                `json:"entity_id,omitempty"`
	State       *string                `json:"state,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
	LastChanged *time.Time             `json:"last_changed,omitempty"`
	LastUpdated *time.Time             `json:"last_updated,omitempty"`
	Context     *Context               `json:"context,omitempty"`
}
