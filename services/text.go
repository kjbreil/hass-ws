package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewTextSetValue creates the object that can be sent to Home Assistant for domain text, service set_value
// "Set value of a text entity."
func NewTextSetValue(target Target) *TextSetValue {
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
		ServiceData: TextSetValueParams{},
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

func (t *TextSetValue) Value(value string) *TextSetValue {
	t.ServiceData.Value = &value
	return t
}
func (t *TextSetValue) JSON() string {
	data, _ := gojson.Marshal(t)
	return string(data)
}
func (t *TextSetValue) Targets() []string {
	return t.Target.EntityId
}
func (t *TextSetValue) Name() string {
	return fmt.Sprintf("%s.%s", *t.Domain, *t.Service)
}
func (t *TextSetValue) SetID(id *int) {
	t.Id = id
}
