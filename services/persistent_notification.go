package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewPersistentNotificationCreate creates the object that can be sent to Home Assistant for domain persistent_notification, service create
// "Shows a notification on the notifications panel."
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
// "Deletes a notification from the notifications panel."
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

// NewPersistentNotificationDismissAll creates the object that can be sent to Home Assistant for domain persistent_notification, service dismiss_all
// "Deletes all notifications from the notifications panel."
func NewPersistentNotificationDismissAll(target Target) *PersistentNotificationDismissAll {
	serviceDomain := "persistent_notification"
	serviceType := "call_service"
	serviceService := "dismiss_all"
	p := &PersistentNotificationDismissAll{
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
	return p
}

type PersistentNotificationDismissAll struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (p *PersistentNotificationDismissAll) JSON() string {
	data, _ := gojson.Marshal(p)
	return string(data)
}
func (p *PersistentNotificationDismissAll) Targets() []string {
	return p.Target.EntityId
}
func (p *PersistentNotificationDismissAll) Name() string {
	return fmt.Sprintf("%s.%s", *p.Domain, *p.Service)
}
