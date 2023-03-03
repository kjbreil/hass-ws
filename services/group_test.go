package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestGroupReload_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *GroupReload
		want   string
	}{{
		fields: NewGroupReload(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"group\",\"service\":\"reload\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestGroupRemove_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *GroupRemove
		want   string
	}{{
		fields: NewGroupRemove(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"group\",\"service\":\"remove\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestGroupSet_JSON(t *testing.T) {
	name := "data"
	objectId := "data"

	tests := []struct {
		name   string
		fields *GroupSet
		want   string
	}{{
		fields: NewGroupSet(Targets("climate.kitchen"), &GroupSetParams{
			Name:     &name,
			ObjectId: &objectId,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"group\",\"service\":\"set\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"name\":\"data\",\"object_id\":\"data\"}}",
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
