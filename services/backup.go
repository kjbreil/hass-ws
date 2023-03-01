package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewBackupCreate creates the object that can be sent to Home Assistant for domain backup, service create
// "Create a new backup."
func NewBackupCreate(entities []string) *BackupCreate {
	serviceDomain := "backup"
	serviceType := "call_service"
	serviceService := "create"
	b := &BackupCreate{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return b
}

type BackupCreate struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (b *BackupCreate) JSON() string {
	data, _ := json.Marshal(b)
	return string(data)
}
func (b *BackupCreate) SetID(id *int) {
	b.Id = id
}
