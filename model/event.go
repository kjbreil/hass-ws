package model

import (
	"time"
)

type EventType string

const (
	EventTypeAll              EventType = "all"
	EventTypeState            EventType = "state"
	EventTypeStateChanged     EventType = "state_changed"
	EventTypeCallService      EventType = "call_service"
	EventTypePysScriptRunning EventType = "pyscript_running"
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
	EventType EventType  `json:"event_type,omitempty"`
	TimeFired *time.Time `json:"time_fired,omitempty"`
	Origin    *string    `json:"origin,omitempty"`
	Context   *Context   `json:"context,omitempty"`
}

type Data struct {
	EntityId *string `json:"entity_id"`
	NewState *State  `json:"new_state,omitempty"`
	OldState *State  `json:"old_state,omitempty"`
}

type Context struct {
	Id       *string `json:"id"`
	ParentId *string `json:"parent_id"`
	UserId   *string `json:"user_id"`
}
