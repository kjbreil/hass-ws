package hass

import (
	"os"

	"github.com/goccy/go-yaml"
)

// Config represents the Home Assistant WebSocket connection configuration
type Config struct {
	Host  string `json:"host"`
	SSL   bool   `json:"ssl"`
	Port  int    `json:"port"`
	Token string `json:"token"`
}

// ParseConfig reads and parses a configuration file
func ParseConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var c struct {
		Websocket Config
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	return &c.Websocket, nil
}
