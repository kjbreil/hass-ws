package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewRemoteDeleteCommand creates the object that can be sent to Home Assistant for domain remote, service delete_command
// "Deletes a command or a list of commands from the database."
func NewRemoteDeleteCommand(entities []string, device *string) *RemoteDeleteCommand {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "delete_command"
	r := &RemoteDeleteCommand{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Device *string `json:"device,omitempty"`
		}{Device: device},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return r
}

type RemoteDeleteCommand struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Device *string `json:"device,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewRemoteLearnCommand(entities []string, commandType *CommandType, device *string, timeout *int) *RemoteLearnCommand {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "learn_command"
	r := &RemoteLearnCommand{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			CommandType *CommandType `json:"command_type,omitempty"`
			Device      *string      `json:"device,omitempty"`
			Timeout     *int         `json:"timeout,omitempty"`
		}{
			CommandType: commandType,
			Device:      device,
			Timeout:     timeout,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return r
}

type RemoteLearnCommand struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		CommandType *CommandType `json:"command_type,omitempty"`
		Device      *string      `json:"device,omitempty"`
		Timeout     *int         `json:"timeout,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewRemoteSendCommand(entities []string, delaySecs *int, device *string, holdSecs *int, numRepeats *int) *RemoteSendCommand {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "send_command"
	r := &RemoteSendCommand{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			DelaySecs  *int    `json:"delay_secs,omitempty"`
			Device     *string `json:"device,omitempty"`
			HoldSecs   *int    `json:"hold_secs,omitempty"`
			NumRepeats *int    `json:"num_repeats,omitempty"`
		}{
			DelaySecs:  delaySecs,
			Device:     device,
			HoldSecs:   holdSecs,
			NumRepeats: numRepeats,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return r
}

type RemoteSendCommand struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		DelaySecs  *int    `json:"delay_secs,omitempty"`
		Device     *string `json:"device,omitempty"`
		HoldSecs   *int    `json:"hold_secs,omitempty"`
		NumRepeats *int    `json:"num_repeats,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewRemoteToggle(entities []string) *RemoteToggle {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "toggle"
	r := &RemoteToggle{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return r
}

type RemoteToggle struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewRemoteTurnOff(entities []string) *RemoteTurnOff {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "turn_off"
	r := &RemoteTurnOff{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return r
}

type RemoteTurnOff struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewRemoteTurnOn(entities []string, activity *string) *RemoteTurnOn {
	serviceDomain := "remote"
	serviceType := "call_service"
	serviceService := "turn_on"
	r := &RemoteTurnOn{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Activity *string `json:"activity,omitempty"`
		}{Activity: activity},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return r
}

type RemoteTurnOn struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Activity *string `json:"activity,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (r *RemoteTurnOn) JSON() string {
	data, _ := json.Marshal(r)
	return string(data)
}
func (r *RemoteTurnOn) SetID(id *int) {
	r.Id = id
}
