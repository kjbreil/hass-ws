package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewRecorderDisable creates the object that can be sent to Home Assistant for domain recorder, service disable
// "Stop the recording of events and state changes"
func NewRecorderDisable(entities []string) *RecorderDisable {
	serviceDomain := "recorder"
	serviceType := "call_service"
	serviceService := "disable"
	r := &RecorderDisable{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return r
}

type RecorderDisable struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (r *RecorderDisable) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RecorderDisable) SetID(id *int) {
	r.Id = id
}

// NewRecorderEnable creates the object that can be sent to Home Assistant for domain recorder, service enable
// "Start the recording of events and state changes"
func NewRecorderEnable(entities []string) *RecorderEnable {
	serviceDomain := "recorder"
	serviceType := "call_service"
	serviceService := "enable"
	r := &RecorderEnable{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return r
}

type RecorderEnable struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (r *RecorderEnable) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RecorderEnable) SetID(id *int) {
	r.Id = id
}

// NewRecorderPurge creates the object that can be sent to Home Assistant for domain recorder, service purge
// "Start purge task - to clean up old data from your database."
func NewRecorderPurge(entities []string, keepDays *int) *RecorderPurge {
	serviceDomain := "recorder"
	serviceType := "call_service"
	serviceService := "purge"
	r := &RecorderPurge{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			KeepDays *int `json:"keep_days,omitempty"`
		}{KeepDays: keepDays},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return r
}

type RecorderPurge struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		KeepDays *int `json:"keep_days,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (r *RecorderPurge) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RecorderPurge) SetID(id *int) {
	r.Id = id
}

// NewRecorderPurgeEntities creates the object that can be sent to Home Assistant for domain recorder, service purge_entities
// "Start purge task to remove specific entities from your database."
func NewRecorderPurgeEntities(entities []string) *RecorderPurgeEntities {
	serviceDomain := "recorder"
	serviceType := "call_service"
	serviceService := "purge_entities"
	r := &RecorderPurgeEntities{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return r
}

type RecorderPurgeEntities struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (r *RecorderPurgeEntities) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RecorderPurgeEntities) SetID(id *int) {
	r.Id = id
}
