package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestLogbookLog_JSON(t *testing.T) {
	logbookLogDomain := "data"
	message := "data"
	logbookLogName := "data"

	tests := []struct {
		name   string
		fields *LogbookLog
		want   string
	}{{
		fields: NewLogbookLog(Targets("climate.kitchen")).LogbookLogDomain(logbookLogDomain).Message(message).LogbookLogName(logbookLogName),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"logbook\",\"service\":\"log\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"domain\":\"data\",\"message\":\"data\",\"name\":\"data\"}}",
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
