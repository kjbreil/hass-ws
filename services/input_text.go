package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewInputTextReload creates the object that can be sent to Home Assistant for domain input_text, service reload
// "Reload the input_text configuration."
func NewInputTextReload(target Target, inputTextReloadParams InputTextReloadParams) *InputTextReload {
	serviceDomain := "input_text"
	serviceType := "call_service"
	serviceService := "reload"
	i := &InputTextReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: inputTextReloadParams,
	}
	return i
}

type InputTextReload struct {
	ServiceBase
	ServiceData InputTextReloadParams `json:"service_data,omitempty"`
}
type InputTextReloadParams struct{}

func (i *InputTextReload) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputTextReload) SetID(id *int) {
	i.Id = id
}

// NewInputTextSetValue creates the object that can be sent to Home Assistant for domain input_text, service set_value
// "Set the value of an input text entity."
func NewInputTextSetValue(target Target, inputTextSetValueParams InputTextSetValueParams) *InputTextSetValue {
	serviceDomain := "input_text"
	serviceType := "call_service"
	serviceService := "set_value"
	i := &InputTextSetValue{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: inputTextSetValueParams,
	}
	return i
}

type InputTextSetValue struct {
	ServiceBase
	ServiceData InputTextSetValueParams `json:"service_data,omitempty"`
}
type InputTextSetValueParams struct {
	Value *string `json:"value,omitempty"`
}

func (i *InputTextSetValue) JSON() string {
	data, _ := json.Marshal(i)
	return string(data)
}
func (i *InputTextSetValue) SetID(id *int) {
	i.Id = id
}
