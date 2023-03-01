package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewNotifyMobileAppAsk creates the object that can be sent to Home Assistant for domain notify, service mobile_app_ask
// "Sends a notification message using the mobile_app_ask integration."
func NewNotifyMobileAppAsk(entities []string, message *string, title *string) *NotifyMobileAppAsk {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "mobile_app_ask"
	n := &NotifyMobileAppAsk{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Message *string `json:"message,omitempty"`
			Title   *string `json:"title,omitempty"`
		}{
			Message: message,
			Title:   title,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return n
}

type NotifyMobileAppAsk struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Message *string `json:"message,omitempty"`
		Title   *string `json:"title,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewNotifyMobileAppFranphone(entities []string, message *string, title *string) *NotifyMobileAppFranphone {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "mobile_app_franphone"
	n := &NotifyMobileAppFranphone{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Message *string `json:"message,omitempty"`
			Title   *string `json:"title,omitempty"`
		}{
			Message: message,
			Title:   title,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return n
}

type NotifyMobileAppFranphone struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Message *string `json:"message,omitempty"`
		Title   *string `json:"title,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewNotifyMobileAppIpad(entities []string, message *string, title *string) *NotifyMobileAppIpad {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "mobile_app_ipad"
	n := &NotifyMobileAppIpad{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Message *string `json:"message,omitempty"`
			Title   *string `json:"title,omitempty"`
		}{
			Message: message,
			Title:   title,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return n
}

type NotifyMobileAppIpad struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Message *string `json:"message,omitempty"`
		Title   *string `json:"title,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewNotifyMobileAppKioskair(entities []string, message *string, title *string) *NotifyMobileAppKioskair {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "mobile_app_kioskair"
	n := &NotifyMobileAppKioskair{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Message *string `json:"message,omitempty"`
			Title   *string `json:"title,omitempty"`
		}{
			Message: message,
			Title:   title,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return n
}

type NotifyMobileAppKioskair struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Message *string `json:"message,omitempty"`
		Title   *string `json:"title,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewNotifyMobileAppLoki(entities []string, message *string, title *string) *NotifyMobileAppLoki {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "mobile_app_loki"
	n := &NotifyMobileAppLoki{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Message *string `json:"message,omitempty"`
			Title   *string `json:"title,omitempty"`
		}{
			Message: message,
			Title:   title,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return n
}

type NotifyMobileAppLoki struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Message *string `json:"message,omitempty"`
		Title   *string `json:"title,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewNotifyMobileAppSkadi(entities []string, message *string, title *string) *NotifyMobileAppSkadi {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "mobile_app_skadi"
	n := &NotifyMobileAppSkadi{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Message *string `json:"message,omitempty"`
			Title   *string `json:"title,omitempty"`
		}{
			Message: message,
			Title:   title,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return n
}

type NotifyMobileAppSkadi struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Message *string `json:"message,omitempty"`
		Title   *string `json:"title,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewNotifyNotify(entities []string, message *string, title *string) *NotifyNotify {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "notify"
	n := &NotifyNotify{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Message *string `json:"message,omitempty"`
			Title   *string `json:"title,omitempty"`
		}{
			Message: message,
			Title:   title,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return n
}

type NotifyNotify struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Message *string `json:"message,omitempty"`
		Title   *string `json:"title,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewNotifyPersistentNotification(entities []string, message *string, title *string) *NotifyPersistentNotification {
	serviceDomain := "notify"
	serviceType := "call_service"
	serviceService := "persistent_notification"
	n := &NotifyPersistentNotification{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Message *string `json:"message,omitempty"`
			Title   *string `json:"title,omitempty"`
		}{
			Message: message,
			Title:   title,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return n
}

type NotifyPersistentNotification struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Message *string `json:"message,omitempty"`
		Title   *string `json:"title,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (n *NotifyPersistentNotification) JSON() string {
	data, _ := json.Marshal(n)
	return string(data)
}
func (n *NotifyPersistentNotification) SetID(id *int) {
	n.Id = id
}
