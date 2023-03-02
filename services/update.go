package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewUpdateClearSkipped creates the object that can be sent to Home Assistant for domain update, service clear_skipped
// "Removes the skipped version marker from an update."
func NewUpdateClearSkipped(entities []string) *UpdateClearSkipped {
	serviceDomain := "update"
	serviceType := "call_service"
	serviceService := "clear_skipped"
	u := &UpdateClearSkipped{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return u
}

type UpdateClearSkipped struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (u *UpdateClearSkipped) JSON() string {
	data, _ := json.Marshal(u)
	return string(data)
}
func (u *UpdateClearSkipped) SetID(id *int) {
	u.Id = id
}

// NewUpdateInstall creates the object that can be sent to Home Assistant for domain update, service install
// "Install an update for this device or service"
func NewUpdateInstall(entities []string, version *string) *UpdateInstall {
	serviceDomain := "update"
	serviceType := "call_service"
	serviceService := "install"
	u := &UpdateInstall{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Version *string `json:"version,omitempty"`
		}{Version: version},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return u
}

type UpdateInstall struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Version *string `json:"version,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewUpdateSkip(entities []string) *UpdateSkip {
	serviceDomain := "update"
	serviceType := "call_service"
	serviceService := "skip"
	u := &UpdateSkip{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return u
}

type UpdateSkip struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (u *UpdateSkip) JSON() string {
	data, _ := json.Marshal(u)
	return string(data)
}
func (u *UpdateSkip) SetID(id *int) {
	u.Id = id
}
