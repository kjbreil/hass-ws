package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewNumberSetValue creates the object that can be sent to Home Assistant for domain number, service set_value
// "Set the value of a Number entity."
func NewNumberSetValue(entities []string, value *string) *NumberSetValue {
	serviceDomain := "number"
	serviceType := "call_service"
	serviceService := "set_value"
	n := &NumberSetValue{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Value *string `json:"value,omitempty"`
		}{Value: value},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return n
}

type NumberSetValue struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Value *string `json:"value,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (n *NumberSetValue) JSON() string {
	data, _ := json.Marshal(n)
	return string(data)
}
func (n *NumberSetValue) SetID(id *int) {
	n.Id = id
}
