package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewRemoteDeleteCommand creates the object that can be sent to Home Assistant for domain remote, service delete_command
// "Deletes a command or a list of commands from the database."
func NewRemoteDeleteCommand(target Target, remoteDeleteCommandParams *RemoteDeleteCommandParams) *RemoteDeleteCommand {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "delete_command"
	r := &RemoteDeleteCommand{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *remoteDeleteCommandParams,
	}
	return r
}

type RemoteDeleteCommand struct {
	ServiceBase
	ServiceData RemoteDeleteCommandParams `json:"service_data,omitempty"`
}
type RemoteDeleteCommandParams struct {
	Device *string `json:"device,omitempty"`
}

func (r *RemoteDeleteCommand) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RemoteDeleteCommand) SetID(id *int) {
	r.Id = id
}

// NewRemoteLearnCommand creates the object that can be sent to Home Assistant for domain remote, service learn_command
// "Learns a command or a list of commands from a device."
func NewRemoteLearnCommand(target Target, remoteLearnCommandParams *RemoteLearnCommandParams) *RemoteLearnCommand {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "learn_command"
	r := &RemoteLearnCommand{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *remoteLearnCommandParams,
	}
	return r
}

type RemoteLearnCommand struct {
	ServiceBase
	ServiceData RemoteLearnCommandParams `json:"service_data,omitempty"`
}
type RemoteLearnCommandParams struct {
	CommandType *CommandType `json:"command_type,omitempty"`
	Device      *string      `json:"device,omitempty"`
	Timeout     *float64     `json:"timeout,omitempty"`
}

func (r *RemoteLearnCommand) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RemoteLearnCommand) SetID(id *int) {
	r.Id = id
}

// NewRemoteSendCommand creates the object that can be sent to Home Assistant for domain remote, service send_command
// "Sends a command or a list of commands to a device."
func NewRemoteSendCommand(target Target, remoteSendCommandParams *RemoteSendCommandParams) *RemoteSendCommand {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "send_command"
	r := &RemoteSendCommand{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *remoteSendCommandParams,
	}
	return r
}

type RemoteSendCommand struct {
	ServiceBase
	ServiceData RemoteSendCommandParams `json:"service_data,omitempty"`
}
type RemoteSendCommandParams struct {
	DelaySecs  *float64 `json:"delay_secs,omitempty"`
	Device     *string  `json:"device,omitempty"`
	HoldSecs   *float64 `json:"hold_secs,omitempty"`
	NumRepeats *float64 `json:"num_repeats,omitempty"`
}

func (r *RemoteSendCommand) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RemoteSendCommand) SetID(id *int) {
	r.Id = id
}

// NewRemoteToggle creates the object that can be sent to Home Assistant for domain remote, service toggle
// "Toggles a device."
func NewRemoteToggle(target Target) *RemoteToggle {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "toggle"
	r := &RemoteToggle{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return r
}

type RemoteToggle struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (r *RemoteToggle) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RemoteToggle) SetID(id *int) {
	r.Id = id
}

// NewRemoteTurnOff creates the object that can be sent to Home Assistant for domain remote, service turn_off
// "Sends the Power Off Command."
func NewRemoteTurnOff(target Target) *RemoteTurnOff {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "turn_off"
	r := &RemoteTurnOff{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return r
}

type RemoteTurnOff struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (r *RemoteTurnOff) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RemoteTurnOff) SetID(id *int) {
	r.Id = id
}

// NewRemoteTurnOn creates the object that can be sent to Home Assistant for domain remote, service turn_on
// "Sends the Power On Command."
func NewRemoteTurnOn(target Target, remoteTurnOnParams *RemoteTurnOnParams) *RemoteTurnOn {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "turn_on"
	r := &RemoteTurnOn{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *remoteTurnOnParams,
	}
	return r
}

type RemoteTurnOn struct {
	ServiceBase
	ServiceData RemoteTurnOnParams `json:"service_data,omitempty"`
}
type RemoteTurnOnParams struct {
	Activity *string `json:"activity,omitempty"`
}

func (r *RemoteTurnOn) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RemoteTurnOn) SetID(id *int) {
	r.Id = id
}
