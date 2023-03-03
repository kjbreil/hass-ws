package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewTextSetValue creates the object that can be sent to Home Assistant for domain text, service set_value
// "Set value of a text entity."
func NewTextSetValue(target Target, textSetValueParams TextSetValueParams) *TextSetValue {
	serviceDomain := "text"
	serviceType := "call_service"
	serviceService := "set_value"
	t := &TextSetValue{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: textSetValueParams,
	}
	return t
}

type TextSetValue struct {
	ServiceBase
	ServiceData TextSetValueParams `json:"service_data,omitempty"`
}
type TextSetValueParams struct {
	Value *string `json:"value,omitempty"`
}

func (t *TextSetValue) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TextSetValue) SetID(id *int) {
	t.Id = id
}
