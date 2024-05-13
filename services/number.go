package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewNumberSetValue creates the object that can be sent to Home Assistant for domain number, service set_value
// "Set the value of a Number entity."
func NewNumberSetValue(target Target) *NumberSetValue {
	serviceDomain := "number"
	serviceType := "call_service"
	serviceService := "set_value"
	n := &NumberSetValue{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: NumberSetValueParams{},
	}
	return n
}

type NumberSetValue struct {
	ServiceBase
	ServiceData NumberSetValueParams `json:"service_data,omitempty"`
}
type NumberSetValueParams struct {
	Value *string `json:"value,omitempty"`
}

func (n *NumberSetValue) Value(value string) *NumberSetValue {
	n.ServiceData.Value = &value
	return n
}
func (n *NumberSetValue) JSON() string {
	data, _ := gojson.Marshal(n)
	return string(data)
}
func (n *NumberSetValue) Targets() []string {
	return n.Target.EntityId
}
func (n *NumberSetValue) Name() string {
	return fmt.Sprintf("%s.%s", *n.Domain, *n.Service)
}
