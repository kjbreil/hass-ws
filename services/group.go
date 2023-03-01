package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewGroupReload creates the object that can be sent to Home Assistant for domain group, service reload
// "Reload group configuration, entities, and notify services."
func NewGroupReload(entities []string) *GroupReload {
	serviceDomain := "group"
	serviceType := "call_service"
	serviceService := "reload"
	g := &GroupReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return g
}

type GroupReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewGroupRemove(entities []string) *GroupRemove {
	serviceDomain := "group"
	serviceType := "call_service"
	serviceService := "remove"
	g := &GroupRemove{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return g
}

type GroupRemove struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewGroupSet(entities []string, name *string, objectId *string) *GroupSet {
	serviceDomain := "group"
	serviceType := "call_service"
	serviceService := "set"
	g := &GroupSet{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Name     *string `json:"name,omitempty"`
			ObjectId *string `json:"object_id,omitempty"`
		}{
			Name:     name,
			ObjectId: objectId,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return g
}

type GroupSet struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Name     *string `json:"name,omitempty"`
		ObjectId *string `json:"object_id,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (g *GroupSet) JSON() string {
	data, _ := json.Marshal(g)
	return string(data)
}
func (g *GroupSet) SetID(id *int) {
	g.Id = id
}
