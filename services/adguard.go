package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewAdguardAddUrl creates the object that can be sent to Home Assistant for domain adguard, service add_url
// "Add a new filter subscription to AdGuard Home."
func NewAdguardAddUrl(target Target, adguardAddUrlParams *AdguardAddUrlParams) *AdguardAddUrl {
	serviceDomain := "adguard"
	serviceType := "call_service"
	serviceService := "add_url"
	a := &AdguardAddUrl{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *adguardAddUrlParams,
	}
	return a
}

type AdguardAddUrl struct {
	ServiceBase
	ServiceData AdguardAddUrlParams `json:"service_data,omitempty"`
}
type AdguardAddUrlParams struct {
	Name *string `json:"name,omitempty"`
	Url  *string `json:"url,omitempty"`
}

func (a *AdguardAddUrl) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AdguardAddUrl) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}
func (a *AdguardAddUrl) SetID(id *int) {
	a.Id = id
}

// NewAdguardDisableUrl creates the object that can be sent to Home Assistant for domain adguard, service disable_url
// "Disables a filter subscription in AdGuard Home."
func NewAdguardDisableUrl(target Target, adguardDisableUrlParams *AdguardDisableUrlParams) *AdguardDisableUrl {
	serviceDomain := "adguard"
	serviceType := "call_service"
	serviceService := "disable_url"
	a := &AdguardDisableUrl{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *adguardDisableUrlParams,
	}
	return a
}

type AdguardDisableUrl struct {
	ServiceBase
	ServiceData AdguardDisableUrlParams `json:"service_data,omitempty"`
}
type AdguardDisableUrlParams struct {
	Url *string `json:"url,omitempty"`
}

func (a *AdguardDisableUrl) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AdguardDisableUrl) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}
func (a *AdguardDisableUrl) SetID(id *int) {
	a.Id = id
}

// NewAdguardEnableUrl creates the object that can be sent to Home Assistant for domain adguard, service enable_url
// "Enables a filter subscription in AdGuard Home."
func NewAdguardEnableUrl(target Target, adguardEnableUrlParams *AdguardEnableUrlParams) *AdguardEnableUrl {
	serviceDomain := "adguard"
	serviceType := "call_service"
	serviceService := "enable_url"
	a := &AdguardEnableUrl{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *adguardEnableUrlParams,
	}
	return a
}

type AdguardEnableUrl struct {
	ServiceBase
	ServiceData AdguardEnableUrlParams `json:"service_data,omitempty"`
}
type AdguardEnableUrlParams struct {
	Url *string `json:"url,omitempty"`
}

func (a *AdguardEnableUrl) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AdguardEnableUrl) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}
func (a *AdguardEnableUrl) SetID(id *int) {
	a.Id = id
}

// NewAdguardRefresh creates the object that can be sent to Home Assistant for domain adguard, service refresh
// "Refresh all filter subscriptions in AdGuard Home."
func NewAdguardRefresh(target Target) *AdguardRefresh {
	serviceDomain := "adguard"
	serviceType := "call_service"
	serviceService := "refresh"
	a := &AdguardRefresh{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return a
}

type AdguardRefresh struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (a *AdguardRefresh) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AdguardRefresh) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}
func (a *AdguardRefresh) SetID(id *int) {
	a.Id = id
}

// NewAdguardRemoveUrl creates the object that can be sent to Home Assistant for domain adguard, service remove_url
// "Removes a filter subscription from AdGuard Home."
func NewAdguardRemoveUrl(target Target, adguardRemoveUrlParams *AdguardRemoveUrlParams) *AdguardRemoveUrl {
	serviceDomain := "adguard"
	serviceType := "call_service"
	serviceService := "remove_url"
	a := &AdguardRemoveUrl{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *adguardRemoveUrlParams,
	}
	return a
}

type AdguardRemoveUrl struct {
	ServiceBase
	ServiceData AdguardRemoveUrlParams `json:"service_data,omitempty"`
}
type AdguardRemoveUrlParams struct {
	Url *string `json:"url,omitempty"`
}

func (a *AdguardRemoveUrl) JSON() string {
	data, _ := json.Marshal(a)
	return string(data)
}
func (a *AdguardRemoveUrl) Name() string {
	return fmt.Sprintf("%s.%s", *a.Domain, *a.Service)
}
func (a *AdguardRemoveUrl) SetID(id *int) {
	a.Id = id
}
