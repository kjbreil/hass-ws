package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewUpdateClearSkipped creates the object that can be sent to Home Assistant for domain update, service clear_skipped
// "Removes the skipped version marker from an update."
func NewUpdateClearSkipped(target Target) *UpdateClearSkipped {
	serviceDomain := "update"
	serviceType := "call_service"
	serviceService := "clear_skipped"
	u := &UpdateClearSkipped{
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
	return u
}

type UpdateClearSkipped struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (u *UpdateClearSkipped) JSON() string {
	data, _ := gojson.Marshal(u)
	return string(data)
}
func (u *UpdateClearSkipped) Targets() []string {
	return u.Target.EntityId
}
func (u *UpdateClearSkipped) Name() string {
	return fmt.Sprintf("%s.%s", *u.Domain, *u.Service)
}

// NewUpdateInstall creates the object that can be sent to Home Assistant for domain update, service install
// "Install an update for this device or service"
func NewUpdateInstall(target Target) *UpdateInstall {
	serviceDomain := "update"
	serviceType := "call_service"
	serviceService := "install"
	u := &UpdateInstall{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: UpdateInstallParams{},
	}
	return u
}

type UpdateInstall struct {
	ServiceBase
	ServiceData UpdateInstallParams `json:"service_data,omitempty"`
}
type UpdateInstallParams struct {
	Version *string `json:"version,omitempty"`
}

func (u *UpdateInstall) Version(version string) *UpdateInstall {
	u.ServiceData.Version = &version
	return u
}
func (u *UpdateInstall) JSON() string {
	data, _ := gojson.Marshal(u)
	return string(data)
}
func (u *UpdateInstall) Targets() []string {
	return u.Target.EntityId
}
func (u *UpdateInstall) Name() string {
	return fmt.Sprintf("%s.%s", *u.Domain, *u.Service)
}

// NewUpdateSkip creates the object that can be sent to Home Assistant for domain update, service skip
// "Mark currently available update as skipped."
func NewUpdateSkip(target Target) *UpdateSkip {
	serviceDomain := "update"
	serviceType := "call_service"
	serviceService := "skip"
	u := &UpdateSkip{
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
	return u
}

type UpdateSkip struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (u *UpdateSkip) JSON() string {
	data, _ := gojson.Marshal(u)
	return string(data)
}
func (u *UpdateSkip) Targets() []string {
	return u.Target.EntityId
}
func (u *UpdateSkip) Name() string {
	return fmt.Sprintf("%s.%s", *u.Domain, *u.Service)
}
