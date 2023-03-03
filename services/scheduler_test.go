package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestSchedulerAdd_JSON(t *testing.T) {
	name := "data"
	repeatType := RepeatTypepause

	tests := []struct {
		name   string
		fields *SchedulerAdd
		want   string
	}{{
		fields: NewSchedulerAdd(Targets("climate.kitchen"), &SchedulerAddParams{
			Name:       &name,
			RepeatType: &repeatType,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"scheduler\",\"service\":\"add\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"name\":\"data\",\"repeat_type\":\"pause\"}}",
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
func TestSchedulerCopy_JSON(t *testing.T) {
	name := "data"

	tests := []struct {
		name   string
		fields *SchedulerCopy
		want   string
	}{{
		fields: NewSchedulerCopy(Targets("climate.kitchen"), &SchedulerCopyParams{Name: &name}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"scheduler\",\"service\":\"copy\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"name\":\"data\"}}",
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
func TestSchedulerEdit_JSON(t *testing.T) {
	name := "data"
	repeatType := RepeatTypepause

	tests := []struct {
		name   string
		fields *SchedulerEdit
		want   string
	}{{
		fields: NewSchedulerEdit(Targets("climate.kitchen"), &SchedulerEditParams{
			Name:       &name,
			RepeatType: &repeatType,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"scheduler\",\"service\":\"edit\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"name\":\"data\",\"repeat_type\":\"pause\"}}",
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
func TestSchedulerRemove_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *SchedulerRemove
		want   string
	}{{
		fields: NewSchedulerRemove(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"scheduler\",\"service\":\"remove\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestSchedulerRunAction_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *SchedulerRunAction
		want   string
	}{{
		fields: NewSchedulerRunAction(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"scheduler\",\"service\":\"run_action\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
