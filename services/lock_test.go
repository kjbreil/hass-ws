package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestLockLock_JSON(t *testing.T) {
	code := "data"

	tests := []struct {
		name   string
		fields *LockLock
		want   string
	}{{
		fields: NewLockLock(Targets("climate.kitchen")).Code(code),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"lock\",\"service\":\"lock\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"code\":\"data\"}}",
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
func TestLockOpen_JSON(t *testing.T) {
	code := "data"

	tests := []struct {
		name   string
		fields *LockOpen
		want   string
	}{{
		fields: NewLockOpen(Targets("climate.kitchen")).Code(code),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"lock\",\"service\":\"open\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"code\":\"data\"}}",
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
func TestLockUnlock_JSON(t *testing.T) {
	code := "data"

	tests := []struct {
		name   string
		fields *LockUnlock
		want   string
	}{{
		fields: NewLockUnlock(Targets("climate.kitchen")).Code(code),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"lock\",\"service\":\"unlock\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"code\":\"data\"}}",
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
