package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewLoggerSetDefaultLevel creates the object that can be sent to Home Assistant for domain logger, service set_default_level
// "Set the default log level for integrations."
func NewLoggerSetDefaultLevel(entities []string, level *Level) *LoggerSetDefaultLevel {
	serviceDomain := "logger"
	serviceType := "call_service"
	serviceService := "set_default_level"
	l := &LoggerSetDefaultLevel{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Level *Level `json:"level,omitempty"`
		}{Level: level},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return l
}

type LoggerSetDefaultLevel struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Level *Level `json:"level,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (l *LoggerSetDefaultLevel) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LoggerSetDefaultLevel) SetID(id *int) {
	l.Id = id
}

// NewLoggerSetLevel creates the object that can be sent to Home Assistant for domain logger, service set_level
// "Set log level for integrations."
func NewLoggerSetLevel(entities []string) *LoggerSetLevel {
	serviceDomain := "logger"
	serviceType := "call_service"
	serviceService := "set_level"
	l := &LoggerSetLevel{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return l
}

type LoggerSetLevel struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (l *LoggerSetLevel) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LoggerSetLevel) SetID(id *int) {
	l.Id = id
}
