package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewLogbookLog creates the object that can be sent to Home Assistant for domain logbook, service log
// "Create a custom entry in your logbook."
func NewLogbookLog(target Target) *LogbookLog {
	serviceDomain := "logbook"
	serviceType := "call_service"
	serviceService := "log"
	l := &LogbookLog{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: LogbookLogParams{},
	}
	return l
}

type LogbookLog struct {
	ServiceBase
	ServiceData LogbookLogParams `json:"service_data,omitempty"`
}
type LogbookLogParams struct {
	Domain  *string `json:"domain,omitempty"`
	Message *string `json:"message,omitempty"`
	Name    *string `json:"name,omitempty"`
}

func (l *LogbookLog) LogbookLogDomain(logbookLogDomain string) *LogbookLog {
	l.ServiceData.Domain = &logbookLogDomain
	return l
}
func (l *LogbookLog) Message(message string) *LogbookLog {
	l.ServiceData.Message = &message
	return l
}
func (l *LogbookLog) LogbookLogName(logbookLogName string) *LogbookLog {
	l.ServiceData.Name = &logbookLogName
	return l
}
func (l *LogbookLog) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LogbookLog) Name() string {
	return fmt.Sprintf("%s.%s", *l.Domain, *l.Service)
}
func (l *LogbookLog) SetID(id *int) {
	l.Id = id
}
