package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestWakeOnLanSendMagicPacket_JSON(t *testing.T) {
	broadcastAddress := "data"
	broadcastPort := 1.2
	mac := "data"

	tests := []struct {
		name   string
		fields *WakeOnLanSendMagicPacket
		want   string
	}{{
		fields: NewWakeOnLanSendMagicPacket(Targets("climate.kitchen")).BroadcastAddress(broadcastAddress).BroadcastPort(broadcastPort).Mac(mac),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"wake_on_lan\",\"service\":\"send_magic_packet\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"broadcast_address\":\"data\",\"broadcast_port\":1.2,\"mac\":\"data\"}}",
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
