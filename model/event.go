package model

import (
	"time"
)

type EventType string

const (
	EventTypeAll              EventType = "all"
	EventTypeStateChanged               = "state_changed"
	EventTypeCallService                = "call_service"
	EventTypePysScriptRunning           = "pyscript_running"
)

func (et EventType) Valid() bool {
	switch et {
	case EventTypeAll,
		EventTypeStateChanged,
		EventTypeCallService,
		EventTypePysScriptRunning:
		return true
	}

	return false
}

type Event struct {
	Data      *Data      `json:"data,omitempty"`
	EventType *EventType `json:"event_type,omitempty"`
	TimeFired *time.Time `json:"time_fired,omitempty"`
	Origin    *string    `json:"origin,omitempty"`
	Context   *Context   `json:"context,omitempty"`
}

type Data struct {
	EntityId *string `json:"entity_id"`
	NewState *State  `json:"new_state,omitempty"`
	OldState *State  `json:"old_state,omitempty"`
}

type State struct {
	EntityId    *string                `json:"entity_id,omitempty"`
	LastChanged *time.Time             `json:"last_changed,omitempty"`
	State       *string                `json:"state,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
	LastUpdated *time.Time             `json:"last_updated,omitempty"`
	Context     *Context               `json:"context,omitempty"`
}

type Context struct {
	Id       *string `json:"id"`
	ParentId *string `json:"parent_id"`
	UserId   *string `json:"user_id"`
}
