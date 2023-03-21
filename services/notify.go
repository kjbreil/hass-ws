package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewNotifyNotify creates the object that can be sent to Home Assistant for domain notify, service notify
// "Sends a notification message using the notify service."
func NewNotifyNotify(target Target, notifyNotifyParams *NotifyNotifyParams) *NotifyNotify {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "notify"
	n := &NotifyNotify{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *notifyNotifyParams,
	}
	return n
}

type NotifyNotify struct {
	ServiceBase
	ServiceData NotifyNotifyParams `json:"service_data,omitempty"`
}
type NotifyNotifyParams struct {
	Message *string `json:"message,omitempty"`
	Title   *string `json:"title,omitempty"`
}

func (n *NotifyNotify) JSON() string {
	data, _ := json.Marshal(n)
	return string(data)
}
func (n *NotifyNotify) Name() string {
	return fmt.Sprintf("%s.%s", *n.Domain, *n.Service)
}
func (n *NotifyNotify) SetID(id *int) {
	n.Id = id
}

// NewNotifyPersistentNotification creates the object that can be sent to Home Assistant for domain notify, service persistent_notification
// "Sends a notification that is visible in the front-end."
func NewNotifyPersistentNotification(target Target, notifyPersistentNotificationParams *NotifyPersistentNotificationParams) *NotifyPersistentNotification {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "persistent_notification"
	n := &NotifyPersistentNotification{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *notifyPersistentNotificationParams,
	}
	return n
}

type NotifyPersistentNotification struct {
	ServiceBase
	ServiceData NotifyPersistentNotificationParams `json:"service_data,omitempty"`
}
type NotifyPersistentNotificationParams struct {
	Message *string `json:"message,omitempty"`
	Title   *string `json:"title,omitempty"`
}

func (n *NotifyPersistentNotification) JSON() string {
	data, _ := json.Marshal(n)
	return string(data)
}
func (n *NotifyPersistentNotification) Name() string {
	return fmt.Sprintf("%s.%s", *n.Domain, *n.Service)
}
func (n *NotifyPersistentNotification) SetID(id *int) {
	n.Id = id
}
