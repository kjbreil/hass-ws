package rest

import "time"

type Api struct {
	Message string `json:"message"`
}

type Config struct {
	Components   []string `json:"components"`
	ConfigDir    string   `json:"config_dir"`
	Elevation    int      `json:"elevation"`
	Latitude     float64  `json:"latitude"`
	LocationName string   `json:"location_name"`
	Longitude    float64  `json:"longitude"`
	TimeZone     string   `json:"time_zone"`
	UnitSystem   struct {
		Length      string `json:"length"`
		Mass        string `json:"mass"`
		Temperature string `json:"temperature"`
		Volume      string `json:"volume"`
	} `json:"unit_system"`
	Version               string   `json:"version"`
	WhitelistExternalDirs []string `json:"whitelist_external_dirs"`
}

type Events []Event
type Event struct {
	Event         string `json:"event"`
	ListenerCount int    `json:"listener_count"`
}

type Service []Service
type Services []struct {
	Domain   string   `json:"domain"`
	Services []string `json:"services"`
}
type Histories []*History

type History struct {
	Attributes struct {
		FriendlyName      string `json:"friendly_name"`
		UnitOfMeasurement string `json:"unit_of_measurement"`
	} `json:"attributes"`
	EntityId    string    `json:"entity_id"`
	LastChanged time.Time `json:"last_changed"`
	LastUpdated time.Time `json:"last_updated"`
	State       string    `json:"state"`
}

type LogBooks []LogBook
type LogBook struct {
	ContextUserId interface{} `json:"context_user_id"`
	Domain        string      `json:"domain"`
	EntityId      string      `json:"entity_id"`
	Message       string      `json:"message"`
	Name          string      `json:"name"`
	When          time.Time   `json:"when"`
}
