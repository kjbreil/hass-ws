package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
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
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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
	data, _ := gojson.Marshal(l)
	return string(data)
}
func (l *LoggerSetDefaultLevel) Targets() []string {
	return l.Target.EntityId
}
func (l *LoggerSetDefaultLevel) Name() string {
	return fmt.Sprintf("%s.%s", *l.Domain, *l.Service)
}

// NewLoggerSetLevel creates the object that can be sent to Home Assistant for domain logger, service set_level
// "Set log level for integrations."
func NewLoggerSetLevel(target Target) *LoggerSetLevel {
	serviceDomain := "logger"
	serviceType := "call_service"
	serviceService := "set_level"
	l := &LoggerSetLevel{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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
	data, _ := gojson.Marshal(l)
	return string(data)
}
func (l *LoggerSetLevel) Targets() []string {
	return l.Target.EntityId
}
func (l *LoggerSetLevel) Name() string {
	return fmt.Sprintf("%s.%s", *l.Domain, *l.Service)
}
