package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewLoggerSetDefaultLevel creates the object that can be sent to Home Assistant for domain logger, service set_default_level
// "Set the default log level for integrations."
func NewLoggerSetDefaultLevel(target Target) *LoggerSetDefaultLevel {
	serviceDomain := "logger"
	serviceType := "call_service"
	serviceService := "set_default_level"
	l := &LoggerSetDefaultLevel{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: LoggerSetDefaultLevelParams{},
	}
	return l
}

type LoggerSetDefaultLevel struct {
	ServiceBase
	ServiceData LoggerSetDefaultLevelParams `json:"service_data,omitempty"`
}
type LoggerSetDefaultLevelParams struct {
	Level *Level `json:"level,omitempty"`
}

func (l *LoggerSetDefaultLevel) Level(level Level) *LoggerSetDefaultLevel {
	l.ServiceData.Level = &level
	return l
}
func (l *LoggerSetDefaultLevel) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LoggerSetDefaultLevel) Name() string {
	return fmt.Sprintf("%s.%s", *l.Domain, *l.Service)
}
func (l *LoggerSetDefaultLevel) SetID(id *int) {
	l.Id = id
}

// NewLoggerSetLevel creates the object that can be sent to Home Assistant for domain logger, service set_level
// "Set log level for integrations."
func NewLoggerSetLevel(target Target) *LoggerSetLevel {
	serviceDomain := "logger"
	serviceType := "call_service"
	serviceService := "set_level"
	l := &LoggerSetLevel{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return l
}

type LoggerSetLevel struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (l *LoggerSetLevel) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LoggerSetLevel) Name() string {
	return fmt.Sprintf("%s.%s", *l.Domain, *l.Service)
}
func (l *LoggerSetLevel) SetID(id *int) {
	l.Id = id
}
