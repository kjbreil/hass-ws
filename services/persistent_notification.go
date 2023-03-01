package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewPersistentNotificationCreate creates the object that can be sent to Home Assistant for domain persistent_notification, service create
// "Show a notification in the frontend."
func NewPersistentNotificationCreate(entities []string, message *string, notificationId *string, title *string) *PersistentNotificationCreate {
	serviceDomain := "persistent_notification"
	serviceType := "call_service"
	serviceService := "create"
	p := &PersistentNotificationCreate{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Message        *string `json:"message,omitempty"`
			NotificationId *string `json:"notification_id,omitempty"`
			Title          *string `json:"title,omitempty"`
		}{
			Message:        message,
			NotificationId: notificationId,
			Title:          title,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return p
}

type PersistentNotificationCreate struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Message        *string `json:"message,omitempty"`
		NotificationId *string `json:"notification_id,omitempty"`
		Title          *string `json:"title,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewPersistentNotificationDismiss(entities []string, notificationId *string) *PersistentNotificationDismiss {
	serviceDomain := "persistent_notification"
	serviceType := "call_service"
	serviceService := "dismiss"
	p := &PersistentNotificationDismiss{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			NotificationId *string `json:"notification_id,omitempty"`
		}{NotificationId: notificationId},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return p
}

type PersistentNotificationDismiss struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		NotificationId *string `json:"notification_id,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewPersistentNotificationMarkRead(entities []string, notificationId *string) *PersistentNotificationMarkRead {
	serviceDomain := "persistent_notification"
	serviceType := "call_service"
	serviceService := "mark_read"
	p := &PersistentNotificationMarkRead{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			NotificationId *string `json:"notification_id,omitempty"`
		}{NotificationId: notificationId},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return p
}

type PersistentNotificationMarkRead struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		NotificationId *string `json:"notification_id,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (p *PersistentNotificationMarkRead) JSON() string {
	data, _ := json.Marshal(p)
	return string(data)
}
func (p *PersistentNotificationMarkRead) SetID(id *int) {
	p.Id = id
}
