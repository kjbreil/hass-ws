package model

import "time"

type EventType string

const (
	EventTypeAll          EventType = "all"
	EventTypeStateChanged           = "state_changed"
)

func (et EventType) Valid() bool {
	switch et {
	case EventTypeAll, EventTypeStateChanged:
		return true
	}

	return false
}

type Event struct {
	Data struct {
		EntityId string `json:"entity_id"`
		NewState struct {
			EntityId    string    `json:"entity_id"`
			LastChanged time.Time `json:"last_changed"`
			State       string    `json:"state"`
			Attributes  struct {
				RgbColor          []int     `json:"rgb_color"`
				ColorTemp         int       `json:"color_temp"`
				SupportedFeatures int       `json:"supported_features"`
				XyColor           []float64 `json:"xy_color"`
				Brightness        int       `json:"brightness"`
				WhiteValue        int       `json:"white_value"`
				FriendlyName      string    `json:"friendly_name"`
			} `json:"attributes"`
			LastUpdated time.Time `json:"last_updated"`
			Context     struct {
				Id       string      `json:"id"`
				ParentId interface{} `json:"parent_id"`
				UserId   string      `json:"user_id"`
			} `json:"context"`
		} `json:"new_state"`
		OldState struct {
			EntityId    string    `json:"entity_id"`
			LastChanged time.Time `json:"last_changed"`
			State       string    `json:"state"`
			Attributes  struct {
				SupportedFeatures int    `json:"supported_features"`
				FriendlyName      string `json:"friendly_name"`
			} `json:"attributes"`
			LastUpdated time.Time `json:"last_updated"`
			Context     struct {
				Id       string      `json:"id"`
				ParentId interface{} `json:"parent_id"`
				UserId   string      `json:"user_id"`
			} `json:"context"`
		} `json:"old_state"`
	} `json:"data"`
	EventType *EventType `json:"event_type,omitempty"`
	TimeFired time.Time  `json:"time_fired"`
	Origin    string     `json:"origin"`
	Context   struct {
		Id       string      `json:"id"`
		ParentId interface{} `json:"parent_id"`
		UserId   string      `json:"user_id"`
	} `json:"context"`
}
