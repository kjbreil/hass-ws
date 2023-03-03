package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewLogbookLog creates the object that can be sent to Home Assistant for domain logbook, service log
// "Create a custom entry in your logbook."
func NewLogbookLog(target Target, logbookLogParams *LogbookLogParams) *LogbookLog {
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
		ServiceData: *logbookLogParams,
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

func (l *LogbookLog) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LogbookLog) SetID(id *int) {
	l.Id = id
}
