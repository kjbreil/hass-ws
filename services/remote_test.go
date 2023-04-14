package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestRemoteDeleteCommand_JSON(t *testing.T) {
	device := "data"

	tests := []struct {
		name   string
		fields *RemoteDeleteCommand
		want   string
	}{{
		fields: NewRemoteDeleteCommand(Targets("climate.kitchen")).Device(device),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"remote\",\"service\":\"delete_command\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"device\":\"data\"}}",
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
func TestRemoteLearnCommand_JSON(t *testing.T) {
	commandType := CommandTypeir
	device := "data"
	timeout := 1.2

	tests := []struct {
		name   string
		fields *RemoteLearnCommand
		want   string
	}{{
		fields: NewRemoteLearnCommand(Targets("climate.kitchen")).CommandType(commandType).Device(device).Timeout(timeout),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"remote\",\"service\":\"learn_command\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"command_type\":\"ir\",\"device\":\"data\",\"timeout\":1.2}}",
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
func TestRemoteSendCommand_JSON(t *testing.T) {
	delaySecs := 1.2
	device := "data"
	holdSecs := 1.2
	numRepeats := 1.2

	tests := []struct {
		name   string
		fields *RemoteSendCommand
		want   string
	}{{
		fields: NewRemoteSendCommand(Targets("climate.kitchen")).DelaySecs(delaySecs).Device(device).HoldSecs(holdSecs).NumRepeats(numRepeats),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"remote\",\"service\":\"send_command\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"delay_secs\":1.2,\"device\":\"data\",\"hold_secs\":1.2,\"num_repeats\":1.2}}",
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
func TestRemoteToggle_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *RemoteToggle
		want   string
	}{{
		fields: NewRemoteToggle(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"remote\",\"service\":\"toggle\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestRemoteTurnOff_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *RemoteTurnOff
		want   string
	}{{
		fields: NewRemoteTurnOff(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"remote\",\"service\":\"turn_off\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
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
func TestRemoteTurnOn_JSON(t *testing.T) {
	activity := "data"

	tests := []struct {
		name   string
		fields *RemoteTurnOn
		want   string
	}{{
		fields: NewRemoteTurnOn(Targets("climate.kitchen")).Activity(activity),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"remote\",\"service\":\"turn_on\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"activity\":\"data\"}}",
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
