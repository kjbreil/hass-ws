package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewAdguardAddUrl creates the object that can be sent to Home Assistant for domain adguard, service add_url
// "Add a new filter subscription to AdGuard Home."
func NewAdguardAddUrl(entities []string, name *string, url *string) *AdguardAddUrl {
	serviceDomain := "adguard"
	serviceType := "call_service"
	serviceService := "add_url"
	a := &AdguardAddUrl{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Name *string `json:"name,omitempty"`
			Url  *string `json:"url,omitempty"`
		}{
			Name: name,
			Url:  url,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AdguardAddUrl struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Name *string `json:"name,omitempty"`
		Url  *string `json:"url,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AdguardAddUrl) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AdguardAddUrl) SetID(id *int) {
	a.Id = id
}

// NewAdguardDisableUrl creates the object that can be sent to Home Assistant for domain adguard, service disable_url
// "Disables a filter subscription in AdGuard Home."
func NewAdguardDisableUrl(entities []string, url *string) *AdguardDisableUrl {
	serviceDomain := "adguard"
	serviceType := "call_service"
	serviceService := "disable_url"
	a := &AdguardDisableUrl{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Url *string `json:"url,omitempty"`
		}{Url: url},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AdguardDisableUrl struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Url *string `json:"url,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AdguardDisableUrl) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AdguardDisableUrl) SetID(id *int) {
	a.Id = id
}

// NewAdguardEnableUrl creates the object that can be sent to Home Assistant for domain adguard, service enable_url
// "Enables a filter subscription in AdGuard Home."
func NewAdguardEnableUrl(entities []string, url *string) *AdguardEnableUrl {
	serviceDomain := "adguard"
	serviceType := "call_service"
	serviceService := "enable_url"
	a := &AdguardEnableUrl{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Url *string `json:"url,omitempty"`
		}{Url: url},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AdguardEnableUrl struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Url *string `json:"url,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AdguardEnableUrl) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AdguardEnableUrl) SetID(id *int) {
	a.Id = id
}

// NewAdguardRefresh creates the object that can be sent to Home Assistant for domain adguard, service refresh
// "Refresh all filter subscriptions in AdGuard Home."
func NewAdguardRefresh(entities []string) *AdguardRefresh {
	serviceDomain := "adguard"
	serviceType := "call_service"
	serviceService := "refresh"
	a := &AdguardRefresh{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AdguardRefresh struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AdguardRefresh) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AdguardRefresh) SetID(id *int) {
	a.Id = id
}

// NewAdguardRemoveUrl creates the object that can be sent to Home Assistant for domain adguard, service remove_url
// "Removes a filter subscription from AdGuard Home."
func NewAdguardRemoveUrl(entities []string, url *string) *AdguardRemoveUrl {
	serviceDomain := "adguard"
	serviceType := "call_service"
	serviceService := "remove_url"
	a := &AdguardRemoveUrl{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Url *string `json:"url,omitempty"`
		}{Url: url},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return a
}

type AdguardRemoveUrl struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Url *string `json:"url,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (a *AdguardRemoveUrl) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AdguardRemoveUrl) SetID(id *int) {
	a.Id = id
}
