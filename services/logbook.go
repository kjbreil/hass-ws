package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewLogbookLog creates the object that can be sent to Home Assistant for domain logbook, service log
// "Create a custom entry in your logbook."
func NewLogbookLog(entities []string, domain *string, message *string, name *string) *LogbookLog {
	serviceDomain := "logbook"
	serviceType := "call_service"
	serviceService := "log"
	l := &LogbookLog{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Domain  *string `json:"domain,omitempty"`
			Message *string `json:"message,omitempty"`
			Name    *string `json:"name,omitempty"`
		}{
			Domain:  domain,
			Message: message,
			Name:    name,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return l
}

type LogbookLog struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Domain  *string `json:"domain,omitempty"`
		Message *string `json:"message,omitempty"`
		Name    *string `json:"name,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (l *LogbookLog) JSON() string {
	data, _ := json.Marshal(l)
	return string(data)
}
func (l *LogbookLog) SetID(id *int) {
	l.Id = id
}
