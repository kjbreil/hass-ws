package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewRemoteDeleteCommand creates the object that can be sent to Home Assistant for domain remote, service delete_command
// "Deletes a command or a list of commands from the database."
func NewRemoteDeleteCommand(target Target) *RemoteDeleteCommand {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "delete_command"
	r := &RemoteDeleteCommand{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: RemoteDeleteCommandParams{},
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

func (r *RemoteDeleteCommand) Device(device string) *RemoteDeleteCommand {
	r.ServiceData.Device = &device
	return r
}
func (r *RemoteDeleteCommand) JSON() string {
	data, _ := gojson.Marshal(r)
	return string(data)
}
func (r *RemoteDeleteCommand) Targets() []string {
	return r.Target.EntityId
}
func (r *RemoteDeleteCommand) Name() string {
	return fmt.Sprintf("%s.%s", *r.Domain, *r.Service)
}

// NewRemoteLearnCommand creates the object that can be sent to Home Assistant for domain remote, service learn_command
// "Learns a command or a list of commands from a device."
func NewRemoteLearnCommand(target Target) *RemoteLearnCommand {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "learn_command"
	r := &RemoteLearnCommand{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: RemoteLearnCommandParams{},
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

func (r *RemoteLearnCommand) CommandType(commandType CommandType) *RemoteLearnCommand {
	r.ServiceData.CommandType = &commandType
	return r
}
func (r *RemoteLearnCommand) Device(device string) *RemoteLearnCommand {
	r.ServiceData.Device = &device
	return r
}
func (r *RemoteLearnCommand) Timeout(timeout float64) *RemoteLearnCommand {
	r.ServiceData.Timeout = &timeout
	return r
}
func (r *RemoteLearnCommand) JSON() string {
	data, _ := gojson.Marshal(r)
	return string(data)
}
func (r *RemoteLearnCommand) Targets() []string {
	return r.Target.EntityId
}
func (r *RemoteLearnCommand) Name() string {
	return fmt.Sprintf("%s.%s", *r.Domain, *r.Service)
}

// NewRemoteSendCommand creates the object that can be sent to Home Assistant for domain remote, service send_command
// "Sends a command or a list of commands to a device."
func NewRemoteSendCommand(target Target) *RemoteSendCommand {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "send_command"
	r := &RemoteSendCommand{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: RemoteSendCommandParams{},
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

func (r *RemoteSendCommand) DelaySecs(delaySecs float64) *RemoteSendCommand {
	r.ServiceData.DelaySecs = &delaySecs
	return r
}
func (r *RemoteSendCommand) Device(device string) *RemoteSendCommand {
	r.ServiceData.Device = &device
	return r
}
func (r *RemoteSendCommand) HoldSecs(holdSecs float64) *RemoteSendCommand {
	r.ServiceData.HoldSecs = &holdSecs
	return r
}
func (r *RemoteSendCommand) NumRepeats(numRepeats float64) *RemoteSendCommand {
	r.ServiceData.NumRepeats = &numRepeats
	return r
}
func (r *RemoteSendCommand) JSON() string {
	data, _ := gojson.Marshal(r)
	return string(data)
}
func (r *RemoteSendCommand) Targets() []string {
	return r.Target.EntityId
}
func (r *RemoteSendCommand) Name() string {
	return fmt.Sprintf("%s.%s", *r.Domain, *r.Service)
}

// NewRemoteToggle creates the object that can be sent to Home Assistant for domain remote, service toggle
// "Toggles a device."
func NewRemoteToggle(target Target) *RemoteToggle {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "toggle"
	r := &RemoteToggle{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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
	data, _ := gojson.Marshal(r)
	return string(data)
}
func (r *RemoteToggle) Targets() []string {
	return r.Target.EntityId
}
func (r *RemoteToggle) Name() string {
	return fmt.Sprintf("%s.%s", *r.Domain, *r.Service)
}

// NewRemoteTurnOff creates the object that can be sent to Home Assistant for domain remote, service turn_off
// "Sends the Power Off Command."
func NewRemoteTurnOff(target Target) *RemoteTurnOff {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "turn_off"
	r := &RemoteTurnOff{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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
	data, _ := gojson.Marshal(r)
	return string(data)
}
func (r *RemoteTurnOff) Targets() []string {
	return r.Target.EntityId
}
func (r *RemoteTurnOff) Name() string {
	return fmt.Sprintf("%s.%s", *r.Domain, *r.Service)
}

// NewRemoteTurnOn creates the object that can be sent to Home Assistant for domain remote, service turn_on
// "Sends the Power On Command."
func NewRemoteTurnOn(target Target) *RemoteTurnOn {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "turn_on"
	r := &RemoteTurnOn{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: RemoteTurnOnParams{},
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

func (r *RemoteTurnOn) Activity(activity string) *RemoteTurnOn {
	r.ServiceData.Activity = &activity
	return r
}
func (r *RemoteTurnOn) JSON() string {
	data, _ := gojson.Marshal(r)
	return string(data)
}
func (r *RemoteTurnOn) Targets() []string {
	return r.Target.EntityId
}
func (r *RemoteTurnOn) Name() string {
	return fmt.Sprintf("%s.%s", *r.Domain, *r.Service)
}
