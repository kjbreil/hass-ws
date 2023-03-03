package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewInputButtonPress creates the object that can be sent to Home Assistant for domain input_button, service press
// "Press the input button entity."
func NewInputButtonPress(target Target, inputButtonPressParams InputButtonPressParams) *InputButtonPress {
	serviceDomain := "input_button"
	serviceType := "call_service"
	serviceService := "press"
	i := &InputButtonPress{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: inputButtonPressParams,
	}
	return i
}

type InputButtonPress struct {
	ServiceBase
	ServiceData InputButtonPressParams `json:"service_data,omitempty"`
}
type InputButtonPressParams struct{}

func (i *InputButtonPress) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputButtonPress) SetID(id *int) {
	i.Id = id
}

// NewInputButtonReload creates the object that can be sent to Home Assistant for domain input_button, service reload
// ""
func NewInputButtonReload(target Target, inputButtonReloadParams InputButtonReloadParams) *InputButtonReload {
	serviceDomain := "input_button"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputButtonReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: inputButtonReloadParams,
	}
	return i
}

type InputButtonReload struct {
	ServiceBase
	ServiceData InputButtonReloadParams `json:"service_data,omitempty"`
}
type InputButtonReloadParams struct{}

func (i *InputButtonReload) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputButtonReload) SetID(id *int) {
	i.Id = id
}
