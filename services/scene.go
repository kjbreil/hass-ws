package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSceneApply creates the object that can be sent to Home Assistant for domain scene, service apply
// "Activate a scene with configuration."
func NewSceneApply(target Target, sceneApplyParams *SceneApplyParams) *SceneApply {
	serviceDomain := "scene"
	serviceType := "call_service"
	serviceService := "apply"
	s := &SceneApply{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *sceneApplyParams,
	}
	return s
}

type SceneApply struct {
	ServiceBase
	ServiceData SceneApplyParams `json:"service_data,omitempty"`
}
type SceneApplyParams struct {
	Transition *float64 `json:"transition,omitempty"`
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
func NewSceneCreate(target Target, sceneCreateParams *SceneCreateParams) *SceneCreate {
	serviceDomain := "scene"
	serviceType := "call_service"
	serviceService := "create"
	s := &SceneCreate{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *sceneCreateParams,
	}
	return s
}

type SceneCreate struct {
	ServiceBase
	ServiceData SceneCreateParams `json:"service_data,omitempty"`
}
type SceneCreateParams struct {
	SceneId *string `json:"scene_id,omitempty"`
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
func NewSceneReload(target Target) *SceneReload {
	serviceDomain := "scene"
	serviceType := "call_service"
	serviceService := "reload"
	s := &SceneReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return s
}

type SceneReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
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
func NewSceneTurnOn(target Target, sceneTurnOnParams *SceneTurnOnParams) *SceneTurnOn {
	serviceDomain := "scene"
	serviceType := "call_service"
	serviceService := "turn_on"
	s := &SceneTurnOn{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *sceneTurnOnParams,
	}
	return s
}

type SceneTurnOn struct {
	ServiceBase
	ServiceData SceneTurnOnParams `json:"service_data,omitempty"`
}
type SceneTurnOnParams struct {
	Transition *float64 `json:"transition,omitempty"`
}

func (s *SceneTurnOn) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SceneTurnOn) SetID(id *int) {
	s.Id = id
}
