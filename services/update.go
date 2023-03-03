package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewUpdateClearSkipped creates the object that can be sent to Home Assistant for domain update, service clear_skipped
// "Removes the skipped version marker from an update."
func NewUpdateClearSkipped(target Target, updateClearSkippedParams UpdateClearSkippedParams) *UpdateClearSkipped {
	serviceDomain := "update"
	serviceType := "call_service"
	serviceService := "clear_skipped"
	u := &UpdateClearSkipped{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: updateClearSkippedParams,
	}
	return u
}

type UpdateClearSkipped struct {
	ServiceBase
	ServiceData UpdateClearSkippedParams `json:"service_data,omitempty"`
}
type UpdateClearSkippedParams struct{}

func (u *UpdateClearSkipped) JSON() string {
	data, _ := json.Marshal(u)
	return string(data)
}
func (u *UpdateClearSkipped) SetID(id *int) {
	u.Id = id
}

// NewUpdateInstall creates the object that can be sent to Home Assistant for domain update, service install
// "Install an update for this device or service"
func NewUpdateInstall(target Target, updateInstallParams UpdateInstallParams) *UpdateInstall {
	serviceDomain := "update"
	serviceType := "call_service"
	serviceService := "install"
	u := &UpdateInstall{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: updateInstallParams,
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

func (u *UpdateInstall) JSON() string {
	data, _ := json.Marshal(u)
	return string(data)
}
func (u *UpdateInstall) SetID(id *int) {
	u.Id = id
}

// NewUpdateSkip creates the object that can be sent to Home Assistant for domain update, service skip
// "Mark currently available update as skipped."
func NewUpdateSkip(target Target, updateSkipParams UpdateSkipParams) *UpdateSkip {
	serviceDomain := "update"
	serviceType := "call_service"
	serviceService := "skip"
	u := &UpdateSkip{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: updateSkipParams,
	}
	return u
}

type UpdateSkip struct {
	ServiceBase
	ServiceData UpdateSkipParams `json:"service_data,omitempty"`
}
type UpdateSkipParams struct{}

func (u *UpdateSkip) JSON() string {
	data, _ := json.Marshal(u)
	return string(data)
}
func (u *UpdateSkip) SetID(id *int) {
	u.Id = id
}
