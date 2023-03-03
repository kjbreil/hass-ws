package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestNotifyMobileAppAsk_JSON(t *testing.T) {
	message := "data"
	title := "data"

	tests := []struct {
		name   string
		fields *NotifyMobileAppAsk
		want   string
	}{{
		fields: NewNotifyMobileAppAsk(Targets("climate.kitchen"), &NotifyMobileAppAskParams{
			Message: &message,
			Title:   &title,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"notify\",\"service\":\"mobile_app_ask\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"message\":\"data\",\"title\":\"data\"}}",
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
func TestNotifyMobileAppFranphone_JSON(t *testing.T) {
	message := "data"
	title := "data"

	tests := []struct {
		name   string
		fields *NotifyMobileAppFranphone
		want   string
	}{{
		fields: NewNotifyMobileAppFranphone(Targets("climate.kitchen"), &NotifyMobileAppFranphoneParams{
			Message: &message,
			Title:   &title,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"notify\",\"service\":\"mobile_app_franphone\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"message\":\"data\",\"title\":\"data\"}}",
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
func TestNotifyMobileAppIpad_JSON(t *testing.T) {
	message := "data"
	title := "data"

	tests := []struct {
		name   string
		fields *NotifyMobileAppIpad
		want   string
	}{{
		fields: NewNotifyMobileAppIpad(Targets("climate.kitchen"), &NotifyMobileAppIpadParams{
			Message: &message,
			Title:   &title,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"notify\",\"service\":\"mobile_app_ipad\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"message\":\"data\",\"title\":\"data\"}}",
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
func TestNotifyMobileAppKioskair_JSON(t *testing.T) {
	message := "data"
	title := "data"

	tests := []struct {
		name   string
		fields *NotifyMobileAppKioskair
		want   string
	}{{
		fields: NewNotifyMobileAppKioskair(Targets("climate.kitchen"), &NotifyMobileAppKioskairParams{
			Message: &message,
			Title:   &title,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"notify\",\"service\":\"mobile_app_kioskair\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"message\":\"data\",\"title\":\"data\"}}",
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
func TestNotifyMobileAppLoki_JSON(t *testing.T) {
	message := "data"
	title := "data"

	tests := []struct {
		name   string
		fields *NotifyMobileAppLoki
		want   string
	}{{
		fields: NewNotifyMobileAppLoki(Targets("climate.kitchen"), &NotifyMobileAppLokiParams{
			Message: &message,
			Title:   &title,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"notify\",\"service\":\"mobile_app_loki\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"message\":\"data\",\"title\":\"data\"}}",
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
func TestNotifyMobileAppSkadi_JSON(t *testing.T) {
	message := "data"
	title := "data"

	tests := []struct {
		name   string
		fields *NotifyMobileAppSkadi
		want   string
	}{{
		fields: NewNotifyMobileAppSkadi(Targets("climate.kitchen"), &NotifyMobileAppSkadiParams{
			Message: &message,
			Title:   &title,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"notify\",\"service\":\"mobile_app_skadi\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"message\":\"data\",\"title\":\"data\"}}",
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
func TestNotifyNotify_JSON(t *testing.T) {
	message := "data"
	title := "data"

	tests := []struct {
		name   string
		fields *NotifyNotify
		want   string
	}{{
		fields: NewNotifyNotify(Targets("climate.kitchen"), &NotifyNotifyParams{
			Message: &message,
			Title:   &title,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"notify\",\"service\":\"notify\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"message\":\"data\",\"title\":\"data\"}}",
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
		fields: NewNotifyPersistentNotification(Targets("climate.kitchen"), &NotifyPersistentNotificationParams{
			Message: &message,
			Title:   &title,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"notify\",\"service\":\"persistent_notification\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"message\":\"data\",\"title\":\"data\"}}",
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
