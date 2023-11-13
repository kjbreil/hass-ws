package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestNotifyNotify_JSON(t *testing.T) {
	data := "data"
	message := "data"
	title := "data"

	tests := []struct {
		name   string
		fields *NotifyNotify
		want   string
	}{{
		fields: NewNotifyNotify(Targets("climate.kitchen")).Data(data).Message(message).Title(title),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"notify\",\"service\":\"notify\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"data\":\"data\",\"message\":\"data\",\"title\":\"data\"}}",
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
func TestNotifyPersistentNotification_JSON(t *testing.T) {
	message := "data"
	title := "data"

	tests := []struct {
		name   string
		fields *NotifyPersistentNotification
		want   string
	}{{
		fields: NewNotifyPersistentNotification(Targets("climate.kitchen")).Message(message).Title(title),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"notify\",\"service\":\"persistent_notification\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"message\":\"data\",\"title\":\"data\"}}",
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
