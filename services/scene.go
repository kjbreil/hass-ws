package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSceneApply creates the object that can be sent to Home Assistant for domain scene, service apply
// "Activate a scene with configuration."
func NewSceneApply(target Target) *SceneApply {
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
		ServiceData: SceneApplyParams{},
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

func (s *SceneApply) Transition(transition float64) *SceneApply {
	s.ServiceData.Transition = &transition
	return s
}
func (s *SceneApply) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SceneApply) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *SceneApply) SetID(id *int) {
	s.Id = id
}

// NewSceneCreate creates the object that can be sent to Home Assistant for domain scene, service create
// "Creates a new scene."
func NewSceneCreate(target Target) *SceneCreate {
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
		ServiceData: SceneCreateParams{},
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

func (s *SceneCreate) SceneId(sceneId string) *SceneCreate {
	s.ServiceData.SceneId = &sceneId
	return s
}
func (s *SceneCreate) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SceneCreate) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
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
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SceneReload) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *SceneReload) SetID(id *int) {
	s.Id = id
}

// NewSceneTurnOn creates the object that can be sent to Home Assistant for domain scene, service turn_on
// "Activate a scene."
func NewSceneTurnOn(target Target) *SceneTurnOn {
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
		ServiceData: SceneTurnOnParams{},
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

func (s *SceneTurnOn) Transition(transition float64) *SceneTurnOn {
	s.ServiceData.Transition = &transition
	return s
}
func (s *SceneTurnOn) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SceneTurnOn) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
func (s *SceneTurnOn) SetID(id *int) {
	s.Id = id
}
