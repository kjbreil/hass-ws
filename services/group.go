package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewGroupReload creates the object that can be sent to Home Assistant for domain group, service reload
// "Reload group configuration, entities, and notify services."
func NewGroupReload(target Target) *GroupReload {
	serviceDomain := "group"
	serviceType := "call_service"
	serviceService := "reload"
	g := &GroupReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return g
}

type GroupReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (g *GroupReload) JSON() string {
	data, _ := json.Marshal(g)
	return string(data)
}
func (g *GroupReload) SetID(id *int) {
	g.Id = id
}

// NewGroupRemove creates the object that can be sent to Home Assistant for domain group, service remove
// "Remove a user group."
func NewGroupRemove(target Target) *GroupRemove {
	serviceDomain := "group"
	serviceType := "call_service"
	serviceService := "remove"
	g := &GroupRemove{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return g
}

type GroupRemove struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (g *GroupRemove) JSON() string {
	data, _ := json.Marshal(g)
	return string(data)
}
func (g *GroupRemove) SetID(id *int) {
	g.Id = id
}

// NewGroupSet creates the object that can be sent to Home Assistant for domain group, service set
// "Create/Update a user group."
func NewGroupSet(target Target, groupSetParams *GroupSetParams) *GroupSet {
	serviceDomain := "group"
	serviceType := "call_service"
	serviceService := "set"
	g := &GroupSet{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *groupSetParams,
	}
	return g
}

type GroupSet struct {
	ServiceBase
	ServiceData GroupSetParams `json:"service_data,omitempty"`
}
type GroupSetParams struct {
	Name     *string `json:"name,omitempty"`
	ObjectId *string `json:"object_id,omitempty"`
}

func (g *GroupSet) JSON() string {
	data, _ := json.Marshal(g)
	return string(data)
}
func (g *GroupSet) SetID(id *int) {
	g.Id = id
}
