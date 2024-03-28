package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestCalendarCreateEvent_JSON(t *testing.T) {
	description := "data"
	location := "data"
	summary := "data"

	tests := []struct {
		name   string
		fields *CalendarCreateEvent
		want   string
	}{{
		fields: NewCalendarCreateEvent(Targets("climate.kitchen")).Description(description).Location(location).Summary(summary),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"calendar\",\"service\":\"create_event\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"description\":\"data\",\"location\":\"data\",\"summary\":\"data\"}}",
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
func TestCalendarListEvents_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CalendarListEvents
		want   string
	}{{
		fields: NewCalendarListEvents(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"calendar\",\"service\":\"list_events\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
