package model

import (
	"strings"
	"time"
)

type States []*State

type State struct {
	EntityId    *string                `json:"entity_id,omitempty"`
	LastChanged *time.Time             `json:"last_changed,omitempty"`
	State       *string                `json:"state,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
	LastUpdated *time.Time             `json:"last_updated,omitempty"`
	Context     *Context               `json:"context,omitempty"`
}

func (s *State) Domain() string {
	if s.EntityId == nil {
		return ""
	}
	split := strings.Split(*s.EntityId, ".")
	if len(split) != 2 {
		return ""
	}
	return split[0]
}
func (s *State) Entity() string {
	if s.EntityId == nil {
		return ""
	}
	split := strings.Split(*s.EntityId, ".")
	if len(split) != 2 {
		return ""
	}
	return split[1]
}

func (s *State) DomainEntity() string {
	if s.EntityId != nil {
		return *s.EntityId
	}
	return ""
}
