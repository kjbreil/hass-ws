package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestCloudRemoteConnect_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CloudRemoteConnect
		want   string
	}{{
		fields: NewCloudRemoteConnect(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"cloud\",\"service\":\"remote_connect\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCloudRemoteDisconnect_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CloudRemoteDisconnect
		want   string
	}{{
		fields: NewCloudRemoteDisconnect(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"cloud\",\"service\":\"remote_disconnect\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
