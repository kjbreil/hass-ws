package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewPersistentNotificationCreate creates the object that can be sent to Home Assistant for domain persistent_notification, service create
// "Show a notification in the frontend."
func NewPersistentNotificationCreate(target Target, persistentNotificationCreateParams *PersistentNotificationCreateParams) *PersistentNotificationCreate {
	serviceDomain := "persistent_notification"
	serviceType := "call_service"
	serviceService := "create"
	p := &PersistentNotificationCreate{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *persistentNotificationCreateParams,
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

func (p *PersistentNotificationCreate) JSON() string {
	data, _ := json.Marshal(p)
	return string(data)
}
func (p *PersistentNotificationCreate) SetID(id *int) {
	p.Id = id
}

// NewPersistentNotificationDismiss creates the object that can be sent to Home Assistant for domain persistent_notification, service dismiss
// "Remove a notification from the frontend."
func NewPersistentNotificationDismiss(target Target, persistentNotificationDismissParams *PersistentNotificationDismissParams) *PersistentNotificationDismiss {
	serviceDomain := "persistent_notification"
	serviceType := "call_service"
	serviceService := "dismiss"
	p := &PersistentNotificationDismiss{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *persistentNotificationDismissParams,
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

func (p *PersistentNotificationDismiss) JSON() string {
	data, _ := json.Marshal(p)
	return string(data)
}
func (p *PersistentNotificationDismiss) SetID(id *int) {
	p.Id = id
}

// NewPersistentNotificationMarkRead creates the object that can be sent to Home Assistant for domain persistent_notification, service mark_read
// "Mark a notification read."
func NewPersistentNotificationMarkRead(target Target, persistentNotificationMarkReadParams *PersistentNotificationMarkReadParams) *PersistentNotificationMarkRead {
	serviceDomain := "persistent_notification"
	serviceType := "call_service"
	serviceService := "mark_read"
	p := &PersistentNotificationMarkRead{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *persistentNotificationMarkReadParams,
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

func (p *PersistentNotificationMarkRead) JSON() string {
	data, _ := json.Marshal(p)
	return string(data)
}
func (p *PersistentNotificationMarkRead) SetID(id *int) {
	p.Id = id
}
