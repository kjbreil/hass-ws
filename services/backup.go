package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewBackupCreate creates the object that can be sent to Home Assistant for domain backup, service create
// "Create a new backup."
func NewBackupCreate(target Target) *BackupCreate {
	serviceDomain := "backup"
	serviceType := "call_service"
	serviceService := "create"
	b := &BackupCreate{
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
	return b
}

type BackupCreate struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (b *BackupCreate) JSON() string {
	data, _ := gojson.Marshal(b)
	return string(data)
}
func (b *BackupCreate) Targets() []string {
	return b.Target.EntityId
}
func (b *BackupCreate) Name() string {
	return fmt.Sprintf("%s.%s", *b.Domain, *b.Service)
}
