package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewSystemLogClear creates the object that can be sent to Home Assistant for domain system_log, service clear
// "Clear all log entries."
func NewSystemLogClear(entities []string) *SystemLogClear {
	serviceDomain := "system_log"
	serviceType := "call_service"
	serviceService := "clear"
	s := &SystemLogClear{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SystemLogClear struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SystemLogClear) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SystemLogClear) SetID(id *int) {
	s.Id = id
}

// NewSystemLogWrite creates the object that can be sent to Home Assistant for domain system_log, service write
// "Write log entry."
func NewSystemLogWrite(entities []string, level *Level, logger *string, message *string) *SystemLogWrite {
	serviceDomain := "system_log"
	serviceType := "call_service"
	serviceService := "write"
	s := &SystemLogWrite{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Level   *Level  `json:"level,omitempty"`
			Logger  *string `json:"logger,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Level:   level,
			Logger:  logger,
			Message: message,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SystemLogWrite struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Level   *Level  `json:"level,omitempty"`
		Logger  *string `json:"logger,omitempty"`
		Message *string `json:"message,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SystemLogWrite) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SystemLogWrite) SetID(id *int) {
	s.Id = id
}
