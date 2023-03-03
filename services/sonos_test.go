package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestSonosClearSleepTimer_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *SonosClearSleepTimer
		want   string
	}{{
		fields: NewSonosClearSleepTimer(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"sonos\",\"service\":\"clear_sleep_timer\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestSonosPlayQueue_JSON(t *testing.T) {
	queuePosition := 1.2

	tests := []struct {
		name   string
		fields *SonosPlayQueue
		want   string
	}{{
		fields: NewSonosPlayQueue(Targets("climate.kitchen"), &SonosPlayQueueParams{QueuePosition: &queuePosition}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"sonos\",\"service\":\"play_queue\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"queue_position\":1.2}}",
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
func TestSonosRemoveFromQueue_JSON(t *testing.T) {
	queuePosition := 1.2

	tests := []struct {
		name   string
		fields *SonosRemoveFromQueue
		want   string
	}{{
		fields: NewSonosRemoveFromQueue(Targets("climate.kitchen"), &SonosRemoveFromQueueParams{QueuePosition: &queuePosition}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"sonos\",\"service\":\"remove_from_queue\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"queue_position\":1.2}}",
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
func TestSonosRestore_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *SonosRestore
		want   string
	}{{
		fields: NewSonosRestore(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"sonos\",\"service\":\"restore\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestSonosSetSleepTimer_JSON(t *testing.T) {
	sleepTime := 1.2

	tests := []struct {
		name   string
		fields *SonosSetSleepTimer
		want   string
	}{{
		fields: NewSonosSetSleepTimer(Targets("climate.kitchen"), &SonosSetSleepTimerParams{SleepTime: &sleepTime}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"sonos\",\"service\":\"set_sleep_timer\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"sleep_time\":1.2}}",
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
func TestSonosSnapshot_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *SonosSnapshot
		want   string
	}{{
		fields: NewSonosSnapshot(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"sonos\",\"service\":\"snapshot\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestSonosUpdateAlarm_JSON(t *testing.T) {
	alarmId := 1.2
	volume := 1.2

	tests := []struct {
		name   string
		fields *SonosUpdateAlarm
		want   string
	}{{
		fields: NewSonosUpdateAlarm(Targets("climate.kitchen"), &SonosUpdateAlarmParams{
			AlarmId: &alarmId,
			Volume:  &volume,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"sonos\",\"service\":\"update_alarm\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"alarm_id\":1.2,\"volume\":1.2}}",
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
