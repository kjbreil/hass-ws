package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestHueActivateScene_JSON(t *testing.T) {
	brightness := 1.2
	speed := 1.2
	transition := 1.2

	tests := []struct {
		name   string
		fields *HueActivateScene
		want   string
	}{{
		fields: NewHueActivateScene(Targets("climate.kitchen"), &HueActivateSceneParams{
			Brightness: &brightness,
			Speed:      &speed,
			Transition: &transition,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"hue\",\"service\":\"activate_scene\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"brightness\":1.2,\"speed\":1.2,\"transition\":1.2}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestHueHueActivateScene_JSON(t *testing.T) {
	groupName := "data"
	sceneName := "data"

	tests := []struct {
		name   string
		fields *HueHueActivateScene
		want   string
	}{{
		fields: NewHueHueActivateScene(Targets("climate.kitchen"), &HueHueActivateSceneParams{
			GroupName: &groupName,
			SceneName: &sceneName,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"hue\",\"service\":\"hue_activate_scene\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"group_name\":\"data\",\"scene_name\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
