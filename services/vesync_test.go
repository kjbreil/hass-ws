package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestVesyncUpdateDevices_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *VesyncUpdateDevices
		want   string
	}{{
		fields: NewVesyncUpdateDevices(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"vesync\",\"service\":\"update_devices\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
