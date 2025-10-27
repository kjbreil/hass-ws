package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewPersistentNotificationCreate creates the object that can be sent to Home Assistant for domain persistent_notification, service create
// "Show a notification in the frontend."
func NewPersistentNotificationCreate(target Target) *PersistentNotificationCreate {
	serviceDomain := "persistent_notification"
	serviceType := "call_service"
	serviceService := "create"
	p := &PersistentNotificationCreate{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: PersistentNotificationCreateParams{},
	}
	return p
}

type PersistentNotificationCreate struct {
	ServiceBase
	ServiceData PersistentNotificationCreateParams `json:"service_data,omitempty"`
}
type PersistentNotificationCreateParams struct {
	Message        *string `json:"message,omitempty"`
	NotificationId *string `json:"notification_id,omitempty"`
	Title          *string `json:"title,omitempty"`
}

func (p *PersistentNotificationCreate) Message(message string) *PersistentNotificationCreate {
	p.ServiceData.Message = &message
	return p
}
func (p *PersistentNotificationCreate) NotificationId(notificationId string) *PersistentNotificationCreate {
	p.ServiceData.NotificationId = &notificationId
	return p
}
func (p *PersistentNotificationCreate) Title(title string) *PersistentNotificationCreate {
	p.ServiceData.Title = &title
	return p
}
func (p *PersistentNotificationCreate) JSON() string {
	data, _ := gojson.Marshal(p)
	return string(data)
}
func (p *PersistentNotificationCreate) Targets() []string {
	return p.Target.EntityId
}
func (p *PersistentNotificationCreate) Name() string {
	return fmt.Sprintf("%s.%s", *p.Domain, *p.Service)
}

// NewPersistentNotificationDismiss creates the object that can be sent to Home Assistant for domain persistent_notification, service dismiss
// "Remove a notification from the frontend."
func NewPersistentNotificationDismiss(target Target) *PersistentNotificationDismiss {
	serviceDomain := "persistent_notification"
	serviceType := "call_service"
	serviceService := "dismiss"
	p := &PersistentNotificationDismiss{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: PersistentNotificationDismissParams{},
	}
	return p
}

type PersistentNotificationDismiss struct {
	ServiceBase
	ServiceData PersistentNotificationDismissParams `json:"service_data,omitempty"`
}
type PersistentNotificationDismissParams struct {
	NotificationId *string `json:"notification_id,omitempty"`
}

func (p *PersistentNotificationDismiss) NotificationId(notificationId string) *PersistentNotificationDismiss {
	p.ServiceData.NotificationId = &notificationId
	return p
}
func (p *PersistentNotificationDismiss) JSON() string {
	data, _ := gojson.Marshal(p)
	return string(data)
}
func (p *PersistentNotificationDismiss) Targets() []string {
	return p.Target.EntityId
}
func (p *PersistentNotificationDismiss) Name() string {
	return fmt.Sprintf("%s.%s", *p.Domain, *p.Service)
}

// NewPersistentNotificationMarkRead creates the object that can be sent to Home Assistant for domain persistent_notification, service mark_read
// "Mark a notification read."
func NewPersistentNotificationMarkRead(target Target) *PersistentNotificationMarkRead {
	serviceDomain := "persistent_notification"
	serviceType := "call_service"
	serviceService := "mark_read"
	p := &PersistentNotificationMarkRead{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: PersistentNotificationMarkReadParams{},
	}
	return p
}

type PersistentNotificationMarkRead struct {
	ServiceBase
	ServiceData PersistentNotificationMarkReadParams `json:"service_data,omitempty"`
}
type PersistentNotificationMarkReadParams struct {
	NotificationId *string `json:"notification_id,omitempty"`
}

func (p *PersistentNotificationMarkRead) NotificationId(notificationId string) *PersistentNotificationMarkRead {
	p.ServiceData.NotificationId = &notificationId
	return p
}
func (p *PersistentNotificationMarkRead) JSON() string {
	data, _ := gojson.Marshal(p)
	return string(data)
}
func (p *PersistentNotificationMarkRead) Targets() []string {
	return p.Target.EntityId
}
func (p *PersistentNotificationMarkRead) Name() string {
	return fmt.Sprintf("%s.%s", *p.Domain, *p.Service)
}
