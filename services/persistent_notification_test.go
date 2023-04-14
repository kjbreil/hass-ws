package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestPersistentNotificationCreate_JSON(t *testing.T) {
	message := "data"
	notificationId := "data"
	title := "data"

	tests := []struct {
		name   string
		fields *PersistentNotificationCreate
		want   string
	}{{
		fields: NewPersistentNotificationCreate(Targets("climate.kitchen")).Message(message).NotificationId(notificationId).Title(title),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"persistent_notification\",\"service\":\"create\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"message\":\"data\",\"notification_id\":\"data\",\"title\":\"data\"}}",
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
func TestPersistentNotificationDismiss_JSON(t *testing.T) {
	notificationId := "data"

	tests := []struct {
		name   string
		fields *PersistentNotificationDismiss
		want   string
	}{{
		fields: NewPersistentNotificationDismiss(Targets("climate.kitchen")).NotificationId(notificationId),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"persistent_notification\",\"service\":\"dismiss\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"notification_id\":\"data\"}}",
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
func TestPersistentNotificationMarkRead_JSON(t *testing.T) {
	notificationId := "data"

	tests := []struct {
		name   string
		fields *PersistentNotificationMarkRead
		want   string
	}{{
		fields: NewPersistentNotificationMarkRead(Targets("climate.kitchen")).NotificationId(notificationId),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"persistent_notification\",\"service\":\"mark_read\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"notification_id\":\"data\"}}",
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
