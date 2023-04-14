package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestCoverCloseCover_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CoverCloseCover
		want   string
	}{{
		fields: NewCoverCloseCover(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"cover\",\"service\":\"close_cover\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCoverCloseCoverTilt_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CoverCloseCoverTilt
		want   string
	}{{
		fields: NewCoverCloseCoverTilt(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"cover\",\"service\":\"close_cover_tilt\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCoverOpenCover_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CoverOpenCover
		want   string
	}{{
		fields: NewCoverOpenCover(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"cover\",\"service\":\"open_cover\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCoverOpenCoverTilt_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CoverOpenCoverTilt
		want   string
	}{{
		fields: NewCoverOpenCoverTilt(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"cover\",\"service\":\"open_cover_tilt\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCoverSetCoverPosition_JSON(t *testing.T) {
	position := 1.2

	tests := []struct {
		name   string
		fields *CoverSetCoverPosition
		want   string
	}{{
		fields: NewCoverSetCoverPosition(Targets("climate.kitchen")).Position(position),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"cover\",\"service\":\"set_cover_position\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"position\":1.2}}",
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
func TestCoverSetCoverTiltPosition_JSON(t *testing.T) {
	tiltPosition := 1.2

	tests := []struct {
		name   string
		fields *CoverSetCoverTiltPosition
		want   string
	}{{
		fields: NewCoverSetCoverTiltPosition(Targets("climate.kitchen")).TiltPosition(tiltPosition),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"cover\",\"service\":\"set_cover_tilt_position\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"tilt_position\":1.2}}",
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
func TestCoverStopCover_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CoverStopCover
		want   string
	}{{
		fields: NewCoverStopCover(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"cover\",\"service\":\"stop_cover\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCoverStopCoverTilt_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CoverStopCoverTilt
		want   string
	}{{
		fields: NewCoverStopCoverTilt(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"cover\",\"service\":\"stop_cover_tilt\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCoverToggle_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CoverToggle
		want   string
	}{{
		fields: NewCoverToggle(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"cover\",\"service\":\"toggle\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCoverToggleCoverTilt_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CoverToggleCoverTilt
		want   string
	}{{
		fields: NewCoverToggleCoverTilt(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"cover\",\"service\":\"toggle_cover_tilt\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
