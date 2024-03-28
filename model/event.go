package model

import (
	"github.com/goccy/go-json"
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
	EntityId    *string      `json:"entity_id"`
	NewState    *State       `json:"new_state,omitempty"`
	OldState    *State       `json:"old_state,omitempty"`
	Domain      *string      `json:"domain,omitempty"`
	Service     *string      `json:"service,omitempty"`
	ServiceData *ServiceData `json:"service_data,omitempty"`
}

type ServiceData struct {
	EntityId []string `json:"entity_id,omitempty"`
}

func (sd *ServiceData) UnmarshalJSON(b []byte) error {
	var d map[string]interface{}
	err := json.Unmarshal(b, &d)
	if err != nil {
		return err
	}
	switch d["entity_id"].(type) {
	case string:
		sd.EntityId = append(sd.EntityId, d["entity_id"].(string))
	case []string:
		sd.EntityId = d["entity_id"].([]string)
	}

	return nil
}

type Context struct {
	Id       *string `json:"id"`
	ParentId *string `json:"parent_id"`
	UserId   *string `json:"user_id"`
}

type T struct {
	Type  string `json:"type"`
	Event struct {
		EventType string `json:"event_type"`
		Data      struct {
			Domain      string `json:"domain"`
			Service     string `json:"service"`
			ServiceData struct {
				BrightnessPct float64  `json:"brightness_pct"`
				ColorTemp     float64  `json:"color_temp"`
				Transition    float64  `json:"transition"`
				EntityId      []string `json:"entity_id"`
			} `json:"service_data"`
		} `json:"data"`
		Origin    string    `json:"origin"`
		TimeFired time.Time `json:"time_fired"`
		Context   struct {
			Id       string      `json:"id"`
			ParentId interface{} `json:"parent_id"`
			UserId   string      `json:"user_id"`
		} `json:"context"`
	} `json:"event"`
	Id int `json:"id"`
}
