package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestMqttDump_JSON(t *testing.T) {
	duration := 1.2
	topic := "data"

	tests := []struct {
		name   string
		fields *MqttDump
		want   string
	}{{
		fields: NewMqttDump(Targets("climate.kitchen")).Duration(duration).Topic(topic),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"mqtt\",\"service\":\"dump\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"duration\":1.2,\"topic\":\"data\"}}",
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
func TestMqttPublish_JSON(t *testing.T) {
	payload := "data"
	qos := Qos0
	topic := "data"

	tests := []struct {
		name   string
		fields *MqttPublish
		want   string
	}{{
		fields: NewMqttPublish(Targets("climate.kitchen")).Payload(payload).Qos(qos).Topic(topic),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"mqtt\",\"service\":\"publish\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"payload\":\"data\",\"qos\":\"0\",\"topic\":\"data\"}}",
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
func TestMqttReload_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MqttReload
		want   string
	}{{
		fields: NewMqttReload(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"mqtt\",\"service\":\"reload\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
