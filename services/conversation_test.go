package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestConversationProcess_JSON(t *testing.T) {
	text := "data"

	tests := []struct {
		name   string
		fields *ConversationProcess
		want   string
	}{{
		fields: NewConversationProcess(Targets("climate.kitchen"), &ConversationProcessParams{Text: &text}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"conversation\",\"service\":\"process\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"text\":\"data\"}}",
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
func TestConversationReload_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *ConversationReload
		want   string
	}{{
		fields: NewConversationReload(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"conversation\",\"service\":\"reload\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
