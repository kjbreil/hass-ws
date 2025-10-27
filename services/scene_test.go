package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestSceneApply_JSON(t *testing.T) {
	transition := 1.2

	tests := []struct {
		name   string
		fields *SceneApply
		want   string
	}{{
		fields: NewSceneApply(Targets("climate.kitchen")).Transition(transition),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"scene\",\"service\":\"apply\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"transition\":1.2}}",
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
func TestSceneCreate_JSON(t *testing.T) {
	sceneId := "data"

	tests := []struct {
		name   string
		fields *SceneCreate
		want   string
	}{{
		fields: NewSceneCreate(Targets("climate.kitchen")).SceneId(sceneId),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"scene\",\"service\":\"create\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"scene_id\":\"data\"}}",
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
func TestSceneReload_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *SceneReload
		want   string
	}{{
		fields: NewSceneReload(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"scene\",\"service\":\"reload\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestSceneTurnOn_JSON(t *testing.T) {
	transition := 1.2

	tests := []struct {
		name   string
		fields *SceneTurnOn
		want   string
	}{{
		fields: NewSceneTurnOn(Targets("climate.kitchen")).Transition(transition),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"scene\",\"service\":\"turn_on\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"transition\":1.2}}",
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
