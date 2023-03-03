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
