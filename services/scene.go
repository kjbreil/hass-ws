package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSceneApply creates the object that can be sent to Home Assistant for domain scene, service apply
// "Activate a scene with configuration."
func NewSceneApply(entities []string, transition *float64) *SceneApply {
	serviceDomain := "scene"
	serviceType := "call_service"
	serviceService := "apply"
	s := &SceneApply{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Transition *float64 `json:"transition,omitempty"`
		}{Transition: transition},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SceneApply struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Transition *float64 `json:"transition,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SceneApply) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SceneApply) SetID(id *int) {
	s.Id = id
}

// NewSceneCreate creates the object that can be sent to Home Assistant for domain scene, service create
// "Creates a new scene."
func NewSceneCreate(entities []string, sceneId *string) *SceneCreate {
	serviceDomain := "scene"
	serviceType := "call_service"
	serviceService := "create"
	s := &SceneCreate{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			SceneId *string `json:"scene_id,omitempty"`
		}{SceneId: sceneId},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SceneCreate struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		SceneId *string `json:"scene_id,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SceneCreate) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SceneCreate) SetID(id *int) {
	s.Id = id
}

// NewSceneReload creates the object that can be sent to Home Assistant for domain scene, service reload
// "Reload the scene configuration."
func NewSceneReload(entities []string) *SceneReload {
	serviceDomain := "scene"
	serviceType := "call_service"
	serviceService := "reload"
	s := &SceneReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SceneReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SceneReload) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SceneReload) SetID(id *int) {
	s.Id = id
}

// NewSceneTurnOn creates the object that can be sent to Home Assistant for domain scene, service turn_on
// "Activate a scene."
func NewSceneTurnOn(entities []string, transition *float64) *SceneTurnOn {
	serviceDomain := "scene"
	serviceType := "call_service"
	serviceService := "turn_on"
	s := &SceneTurnOn{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Transition *float64 `json:"transition,omitempty"`
		}{Transition: transition},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SceneTurnOn struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Transition *float64 `json:"transition,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SceneTurnOn) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SceneTurnOn) SetID(id *int) {
	s.Id = id
}
