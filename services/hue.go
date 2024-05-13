package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewHueActivateScene creates the object that can be sent to Home Assistant for domain hue, service activate_scene
// "Activate a Hue scene with more control over the options."
func NewHueActivateScene(target Target) *HueActivateScene {
	serviceDomain := "hue"
	serviceType := "call_service"
	serviceService := "activate_scene"
	h := &HueActivateScene{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: HueActivateSceneParams{},
	}
	return h
}

type HueActivateScene struct {
	ServiceBase
	ServiceData HueActivateSceneParams `json:"service_data,omitempty"`
}
type HueActivateSceneParams struct {
	Brightness *float64 `json:"brightness,omitempty"`
	Speed      *float64 `json:"speed,omitempty"`
	Transition *float64 `json:"transition,omitempty"`
}

func (h *HueActivateScene) Brightness(brightness float64) *HueActivateScene {
	h.ServiceData.Brightness = &brightness
	return h
}
func (h *HueActivateScene) Speed(speed float64) *HueActivateScene {
	h.ServiceData.Speed = &speed
	return h
}
func (h *HueActivateScene) Transition(transition float64) *HueActivateScene {
	h.ServiceData.Transition = &transition
	return h
}
func (h *HueActivateScene) JSON() string {
	data, _ := gojson.Marshal(h)
	return string(data)
}
func (h *HueActivateScene) Targets() []string {
	return h.Target.EntityId
}
func (h *HueActivateScene) Name() string {
	return fmt.Sprintf("%s.%s", *h.Domain, *h.Service)
}

// NewHueHueActivateScene creates the object that can be sent to Home Assistant for domain hue, service hue_activate_scene
// "Activate a hue scene stored in the hue hub."
func NewHueHueActivateScene(target Target) *HueHueActivateScene {
	serviceDomain := "hue"
	serviceType := "call_service"
	serviceService := "hue_activate_scene"
	h := &HueHueActivateScene{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: HueHueActivateSceneParams{},
	}
	return h
}

type HueHueActivateScene struct {
	ServiceBase
	ServiceData HueHueActivateSceneParams `json:"service_data,omitempty"`
}
type HueHueActivateSceneParams struct {
	GroupName *string `json:"group_name,omitempty"`
	SceneName *string `json:"scene_name,omitempty"`
}

func (h *HueHueActivateScene) GroupName(groupName string) *HueHueActivateScene {
	h.ServiceData.GroupName = &groupName
	return h
}
func (h *HueHueActivateScene) SceneName(sceneName string) *HueHueActivateScene {
	h.ServiceData.SceneName = &sceneName
	return h
}
func (h *HueHueActivateScene) JSON() string {
	data, _ := gojson.Marshal(h)
	return string(data)
}
func (h *HueHueActivateScene) Targets() []string {
	return h.Target.EntityId
}
func (h *HueHueActivateScene) Name() string {
	return fmt.Sprintf("%s.%s", *h.Domain, *h.Service)
}
