package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewHueActivateScene creates the object that can be sent to Home Assistant for domain hue, service activate_scene
// "Activate a Hue scene with more control over the options."
func NewHueActivateScene(target Target, hueActivateSceneParams *HueActivateSceneParams) *HueActivateScene {
	serviceDomain := "hue"
	serviceType := "call_service"
	serviceService := "activate_scene"
	h := &HueActivateScene{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *hueActivateSceneParams,
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

func (h *HueActivateScene) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HueActivateScene) SetID(id *int) {
	h.Id = id
}

// NewHueHueActivateScene creates the object that can be sent to Home Assistant for domain hue, service hue_activate_scene
// "Activate a hue scene stored in the hue hub."
func NewHueHueActivateScene(target Target, hueHueActivateSceneParams *HueHueActivateSceneParams) *HueHueActivateScene {
	serviceDomain := "hue"
	serviceType := "call_service"
	serviceService := "hue_activate_scene"
	h := &HueHueActivateScene{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *hueHueActivateSceneParams,
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

func (h *HueHueActivateScene) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HueHueActivateScene) SetID(id *int) {
	h.Id = id
}
