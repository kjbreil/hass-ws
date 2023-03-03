package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewLoggerSetDefaultLevel creates the object that can be sent to Home Assistant for domain logger, service set_default_level
// "Set the default log level for integrations."
func NewLoggerSetDefaultLevel(target Target, loggerSetDefaultLevelParams LoggerSetDefaultLevelParams) *LoggerSetDefaultLevel {
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
		ServiceData: loggerSetDefaultLevelParams,
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

func (l *LoggerSetDefaultLevel) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LoggerSetDefaultLevel) SetID(id *int) {
	l.Id = id
}

// NewLoggerSetLevel creates the object that can be sent to Home Assistant for domain logger, service set_level
// "Set log level for integrations."
func NewLoggerSetLevel(target Target, loggerSetLevelParams LoggerSetLevelParams) *LoggerSetLevel {
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
		ServiceData: loggerSetLevelParams,
	}
	return l
}

type LoggerSetLevel struct {
	ServiceBase
	ServiceData LoggerSetLevelParams `json:"service_data,omitempty"`
}
type LoggerSetLevelParams struct{}

func (l *LoggerSetLevel) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LoggerSetLevel) SetID(id *int) {
	l.Id = id
}
