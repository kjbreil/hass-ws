package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewNumberSetValue creates the object that can be sent to Home Assistant for domain number, service set_value
// "Set the value of a Number entity."
func NewNumberSetValue(target Target, numberSetValueParams *NumberSetValueParams) *NumberSetValue {
	serviceDomain := "number"
	serviceType := "call_service"
	serviceService := "set_value"
	n := &NumberSetValue{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *numberSetValueParams,
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

func (n *NumberSetValue) JSON() string {
	data, _ := json.Marshal(n)
	return string(data)
}
func (n *NumberSetValue) Name() string {
	return fmt.Sprintf("%s.%s", *n.Domain, *n.Service)
}
func (n *NumberSetValue) SetID(id *int) {
	n.Id = id
}
