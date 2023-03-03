package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewNotifyMobileAppAsk creates the object that can be sent to Home Assistant for domain notify, service mobile_app_ask
// "Sends a notification message using the mobile_app_ask integration."
func NewNotifyMobileAppAsk(target Target, notifyMobileAppAskParams NotifyMobileAppAskParams) *NotifyMobileAppAsk {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "mobile_app_ask"
	n := &NotifyMobileAppAsk{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: notifyMobileAppAskParams,
	}
	return n
}

type NotifyMobileAppAsk struct {
	ServiceBase
	ServiceData NotifyMobileAppAskParams `json:"service_data,omitempty"`
}
type NotifyMobileAppAskParams struct {
	Message *string `json:"message,omitempty"`
	Title   *string `json:"title,omitempty"`
}

func (n *NotifyMobileAppAsk) JSON() string {
	data, _ := json.Marshal(n)
	return string(data)
}
func (n *NotifyMobileAppAsk) SetID(id *int) {
	n.Id = id
}

// NewNotifyMobileAppFranphone creates the object that can be sent to Home Assistant for domain notify, service mobile_app_franphone
// "Sends a notification message using the mobile_app_franphone integration."
func NewNotifyMobileAppFranphone(target Target, notifyMobileAppFranphoneParams NotifyMobileAppFranphoneParams) *NotifyMobileAppFranphone {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "mobile_app_franphone"
	n := &NotifyMobileAppFranphone{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: notifyMobileAppFranphoneParams,
	}
	return n
}

type NotifyMobileAppFranphone struct {
	ServiceBase
	ServiceData NotifyMobileAppFranphoneParams `json:"service_data,omitempty"`
}
type NotifyMobileAppFranphoneParams struct {
	Message *string `json:"message,omitempty"`
	Title   *string `json:"title,omitempty"`
}

func (n *NotifyMobileAppFranphone) JSON() string {
	data, _ := json.Marshal(n)
	return string(data)
}
func (n *NotifyMobileAppFranphone) SetID(id *int) {
	n.Id = id
}

// NewNotifyMobileAppIpad creates the object that can be sent to Home Assistant for domain notify, service mobile_app_ipad
// "Sends a notification message using the mobile_app_ipad integration."
func NewNotifyMobileAppIpad(target Target, notifyMobileAppIpadParams NotifyMobileAppIpadParams) *NotifyMobileAppIpad {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "mobile_app_ipad"
	n := &NotifyMobileAppIpad{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: notifyMobileAppIpadParams,
	}
	return n
}

type NotifyMobileAppIpad struct {
	ServiceBase
	ServiceData NotifyMobileAppIpadParams `json:"service_data,omitempty"`
}
type NotifyMobileAppIpadParams struct {
	Message *string `json:"message,omitempty"`
	Title   *string `json:"title,omitempty"`
}

func (n *NotifyMobileAppIpad) JSON() string {
	data, _ := json.Marshal(n)
	return string(data)
}
func (n *NotifyMobileAppIpad) SetID(id *int) {
	n.Id = id
}

// NewNotifyMobileAppKioskair creates the object that can be sent to Home Assistant for domain notify, service mobile_app_kioskair
// "Sends a notification message using the mobile_app_kioskair integration."
func NewNotifyMobileAppKioskair(target Target, notifyMobileAppKioskairParams NotifyMobileAppKioskairParams) *NotifyMobileAppKioskair {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "mobile_app_kioskair"
	n := &NotifyMobileAppKioskair{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: notifyMobileAppKioskairParams,
	}
	return n
}

type NotifyMobileAppKioskair struct {
	ServiceBase
	ServiceData NotifyMobileAppKioskairParams `json:"service_data,omitempty"`
}
type NotifyMobileAppKioskairParams struct {
	Message *string `json:"message,omitempty"`
	Title   *string `json:"title,omitempty"`
}

func (n *NotifyMobileAppKioskair) JSON() string {
	data, _ := json.Marshal(n)
	return string(data)
}
func (n *NotifyMobileAppKioskair) SetID(id *int) {
	n.Id = id
}

// NewNotifyMobileAppLoki creates the object that can be sent to Home Assistant for domain notify, service mobile_app_loki
// "Sends a notification message using the mobile_app_loki integration."
func NewNotifyMobileAppLoki(target Target, notifyMobileAppLokiParams NotifyMobileAppLokiParams) *NotifyMobileAppLoki {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "mobile_app_loki"
	n := &NotifyMobileAppLoki{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: notifyMobileAppLokiParams,
	}
	return n
}

type NotifyMobileAppLoki struct {
	ServiceBase
	ServiceData NotifyMobileAppLokiParams `json:"service_data,omitempty"`
}
type NotifyMobileAppLokiParams struct {
	Message *string `json:"message,omitempty"`
	Title   *string `json:"title,omitempty"`
}

func (n *NotifyMobileAppLoki) JSON() string {
	data, _ := json.Marshal(n)
	return string(data)
}
func (n *NotifyMobileAppLoki) SetID(id *int) {
	n.Id = id
}

// NewNotifyMobileAppSkadi creates the object that can be sent to Home Assistant for domain notify, service mobile_app_skadi
// "Sends a notification message using the mobile_app_skadi integration."
func NewNotifyMobileAppSkadi(target Target, notifyMobileAppSkadiParams NotifyMobileAppSkadiParams) *NotifyMobileAppSkadi {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "mobile_app_skadi"
	n := &NotifyMobileAppSkadi{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: notifyMobileAppSkadiParams,
	}
	return n
}

type NotifyMobileAppSkadi struct {
	ServiceBase
	ServiceData NotifyMobileAppSkadiParams `json:"service_data,omitempty"`
}
type NotifyMobileAppSkadiParams struct {
	Message *string `json:"message,omitempty"`
	Title   *string `json:"title,omitempty"`
}

func (n *NotifyMobileAppSkadi) JSON() string {
	data, _ := json.Marshal(n)
	return string(data)
}
func (n *NotifyMobileAppSkadi) SetID(id *int) {
	n.Id = id
}

// NewNotifyNotify creates the object that can be sent to Home Assistant for domain notify, service notify
// "Sends a notification message using the notify service."
func NewNotifyNotify(target Target, notifyNotifyParams NotifyNotifyParams) *NotifyNotify {
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
		ServiceData: notifyNotifyParams,
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
func (n *NotifyNotify) SetID(id *int) {
	n.Id = id
}

// NewNotifyPersistentNotification creates the object that can be sent to Home Assistant for domain notify, service persistent_notification
// "Sends a notification that is visible in the front-end."
func NewNotifyPersistentNotification(target Target, notifyPersistentNotificationParams NotifyPersistentNotificationParams) *NotifyPersistentNotification {
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
		ServiceData: notifyPersistentNotificationParams,
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
func (n *NotifyPersistentNotification) SetID(id *int) {
	n.Id = id
}
