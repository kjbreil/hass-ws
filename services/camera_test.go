package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestCameraDisableMotionDetection_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CameraDisableMotionDetection
		want   string
	}{{
		fields: NewCameraDisableMotionDetection(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"camera\",\"service\":\"disable_motion_detection\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCameraEnableMotionDetection_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CameraEnableMotionDetection
		want   string
	}{{
		fields: NewCameraEnableMotionDetection(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"camera\",\"service\":\"enable_motion_detection\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCameraPlayStream_JSON(t *testing.T) {
	format := Formathls

	tests := []struct {
		name   string
		fields *CameraPlayStream
		want   string
	}{{
		fields: NewCameraPlayStream(Targets("climate.kitchen"), &CameraPlayStreamParams{Format: &format}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"camera\",\"service\":\"play_stream\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"format\":\"hls\"}}",
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
func TestCameraRecord_JSON(t *testing.T) {
	duration := 1.2
	filename := "data"
	lookback := 1.2

	tests := []struct {
		name   string
		fields *CameraRecord
		want   string
	}{{
		fields: NewCameraRecord(Targets("climate.kitchen"), &CameraRecordParams{
			Duration: &duration,
			Filename: &filename,
			Lookback: &lookback,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"camera\",\"service\":\"record\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"duration\":1.2,\"filename\":\"data\",\"lookback\":1.2}}",
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
func TestCameraSnapshot_JSON(t *testing.T) {
	filename := "data"

	tests := []struct {
		name   string
		fields *CameraSnapshot
		want   string
	}{{
		fields: NewCameraSnapshot(Targets("climate.kitchen"), &CameraSnapshotParams{Filename: &filename}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"camera\",\"service\":\"snapshot\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"filename\":\"data\"}}",
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
func TestCameraTurnOff_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CameraTurnOff
		want   string
	}{{
		fields: NewCameraTurnOff(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"camera\",\"service\":\"turn_off\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestCameraTurnOn_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *CameraTurnOn
		want   string
	}{{
		fields: NewCameraTurnOn(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"camera\",\"service\":\"turn_on\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
