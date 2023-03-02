package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewHueActivateScene creates the object that can be sent to Home Assistant for domain hue, service activate_scene
// "Activate a Hue scene with more control over the options."
func NewHueActivateScene(entities []string, brightness *int, speed *int, transition *int) *HueActivateScene {
	serviceDomain := "hue"
	serviceType := "call_service"
	serviceService := "activate_scene"
	h := &HueActivateScene{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Brightness *int `json:"brightness,omitempty"`
			Speed      *int `json:"speed,omitempty"`
			Transition *int `json:"transition,omitempty"`
		}{
			Brightness: brightness,
			Speed:      speed,
			Transition: transition,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return h
}

type HueActivateScene struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Brightness *int `json:"brightness,omitempty"`
		Speed      *int `json:"speed,omitempty"`
		Transition *int `json:"transition,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewHueHueActivateScene(entities []string, groupName *string, sceneName *string) *HueHueActivateScene {
	serviceDomain := "hue"
	serviceType := "call_service"
	serviceService := "hue_activate_scene"
	h := &HueHueActivateScene{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			GroupName *string `json:"group_name,omitempty"`
			SceneName *string `json:"scene_name,omitempty"`
		}{
			GroupName: groupName,
			SceneName: sceneName,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return h
}

type HueHueActivateScene struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		GroupName *string `json:"group_name,omitempty"`
		SceneName *string `json:"scene_name,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HueHueActivateScene) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HueHueActivateScene) SetID(id *int) {
	h.Id = id
}
