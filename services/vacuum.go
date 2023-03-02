package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewVacuumCleanSpot creates the object that can be sent to Home Assistant for domain vacuum, service clean_spot
// "Tell the vacuum cleaner to do a spot clean-up."
func NewVacuumCleanSpot(entities []string) *VacuumCleanSpot {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "clean_spot"
	v := &VacuumCleanSpot{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return v
}

type VacuumCleanSpot struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (v *VacuumCleanSpot) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumCleanSpot) SetID(id *int) {
	v.Id = id
}

// NewVacuumLocate creates the object that can be sent to Home Assistant for domain vacuum, service locate
// "Locate the vacuum cleaner robot."
func NewVacuumLocate(entities []string) *VacuumLocate {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "locate"
	v := &VacuumLocate{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return v
}

type VacuumLocate struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (v *VacuumLocate) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumLocate) SetID(id *int) {
	v.Id = id
}

// NewVacuumPause creates the object that can be sent to Home Assistant for domain vacuum, service pause
// "Pause the cleaning task."
func NewVacuumPause(entities []string) *VacuumPause {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "pause"
	v := &VacuumPause{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return v
}

type VacuumPause struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (v *VacuumPause) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumPause) SetID(id *int) {
	v.Id = id
}

// NewVacuumReturnToBase creates the object that can be sent to Home Assistant for domain vacuum, service return_to_base
// "Tell the vacuum cleaner to return to its dock."
func NewVacuumReturnToBase(entities []string) *VacuumReturnToBase {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "return_to_base"
	v := &VacuumReturnToBase{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return v
}

type VacuumReturnToBase struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (v *VacuumReturnToBase) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumReturnToBase) SetID(id *int) {
	v.Id = id
}

// NewVacuumSendCommand creates the object that can be sent to Home Assistant for domain vacuum, service send_command
// "Send a raw command to the vacuum cleaner."
func NewVacuumSendCommand(entities []string, command *string) *VacuumSendCommand {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "send_command"
	v := &VacuumSendCommand{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Command *string `json:"command,omitempty"`
		}{Command: command},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return v
}

type VacuumSendCommand struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Command *string `json:"command,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (v *VacuumSendCommand) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumSendCommand) SetID(id *int) {
	v.Id = id
}

// NewVacuumSetFanSpeed creates the object that can be sent to Home Assistant for domain vacuum, service set_fan_speed
// "Set the fan speed of the vacuum cleaner."
func NewVacuumSetFanSpeed(entities []string, fanSpeed *string) *VacuumSetFanSpeed {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "set_fan_speed"
	v := &VacuumSetFanSpeed{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			FanSpeed *string `json:"fan_speed,omitempty"`
		}{FanSpeed: fanSpeed},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return v
}

type VacuumSetFanSpeed struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		FanSpeed *string `json:"fan_speed,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (v *VacuumSetFanSpeed) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumSetFanSpeed) SetID(id *int) {
	v.Id = id
}

// NewVacuumStart creates the object that can be sent to Home Assistant for domain vacuum, service start
// "Start or resume the cleaning task."
func NewVacuumStart(entities []string) *VacuumStart {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "start"
	v := &VacuumStart{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return v
}

type VacuumStart struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (v *VacuumStart) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumStart) SetID(id *int) {
	v.Id = id
}

// NewVacuumStartPause creates the object that can be sent to Home Assistant for domain vacuum, service start_pause
// "Start, pause, or resume the cleaning task."
func NewVacuumStartPause(entities []string) *VacuumStartPause {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "start_pause"
	v := &VacuumStartPause{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return v
}

type VacuumStartPause struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (v *VacuumStartPause) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumStartPause) SetID(id *int) {
	v.Id = id
}

// NewVacuumStop creates the object that can be sent to Home Assistant for domain vacuum, service stop
// "Stop the current cleaning task."
func NewVacuumStop(entities []string) *VacuumStop {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "stop"
	v := &VacuumStop{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return v
}

type VacuumStop struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (v *VacuumStop) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumStop) SetID(id *int) {
	v.Id = id
}

// NewVacuumToggle creates the object that can be sent to Home Assistant for domain vacuum, service toggle
// ""
func NewVacuumToggle(entities []string) *VacuumToggle {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "toggle"
	v := &VacuumToggle{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return v
}

type VacuumToggle struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (v *VacuumToggle) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumToggle) SetID(id *int) {
	v.Id = id
}

// NewVacuumTurnOff creates the object that can be sent to Home Assistant for domain vacuum, service turn_off
// "Stop the current cleaning task and return to home."
func NewVacuumTurnOff(entities []string) *VacuumTurnOff {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "turn_off"
	v := &VacuumTurnOff{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return v
}

type VacuumTurnOff struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (v *VacuumTurnOff) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumTurnOff) SetID(id *int) {
	v.Id = id
}

// NewVacuumTurnOn creates the object that can be sent to Home Assistant for domain vacuum, service turn_on
// "Start a new cleaning task."
func NewVacuumTurnOn(entities []string) *VacuumTurnOn {
	serviceDomain := "vacuum"
	serviceType := "call_service"
	serviceService := "turn_on"
	v := &VacuumTurnOn{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return v
}

type VacuumTurnOn struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (v *VacuumTurnOn) JSON() string {
	data, _ := json.Marshal(v)
	return string(data)
}
func (v *VacuumTurnOn) SetID(id *int) {
	v.Id = id
}
