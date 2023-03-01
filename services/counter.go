package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewCounterConfigure creates the object that can be sent to Home Assistant for domain counter, service configure
// "Change counter parameters."
func NewCounterConfigure(entities []string, initial *int, maximum *int, minimum *int, step *int, value *int) *CounterConfigure {
	serviceDomain := "counter"
	serviceType := "call_service"
	serviceService := "configure"
	c := &CounterConfigure{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Initial *int `json:"initial,omitempty"`
			Maximum *int `json:"maximum,omitempty"`
			Minimum *int `json:"minimum,omitempty"`
			Step    *int `json:"step,omitempty"`
			Value   *int `json:"value,omitempty"`
		}{
			Initial: initial,
			Maximum: maximum,
			Minimum: minimum,
			Step:    step,
			Value:   value,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CounterConfigure struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Initial *int `json:"initial,omitempty"`
		Maximum *int `json:"maximum,omitempty"`
		Minimum *int `json:"minimum,omitempty"`
		Step    *int `json:"step,omitempty"`
		Value   *int `json:"value,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CounterConfigure) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CounterConfigure) SetID(id *int) {
	c.Id = id
}

// NewCounterDecrement creates the object that can be sent to Home Assistant for domain counter, service decrement
// "Decrement a counter."
func NewCounterDecrement(entities []string) *CounterDecrement {
	serviceDomain := "counter"
	serviceType := "call_service"
	serviceService := "decrement"
	c := &CounterDecrement{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CounterDecrement struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CounterDecrement) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CounterDecrement) SetID(id *int) {
	c.Id = id
}

// NewCounterIncrement creates the object that can be sent to Home Assistant for domain counter, service increment
// "Increment a counter."
func NewCounterIncrement(entities []string) *CounterIncrement {
	serviceDomain := "counter"
	serviceType := "call_service"
	serviceService := "increment"
	c := &CounterIncrement{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CounterIncrement struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CounterIncrement) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CounterIncrement) SetID(id *int) {
	c.Id = id
}

// NewCounterReset creates the object that can be sent to Home Assistant for domain counter, service reset
// "Reset a counter."
func NewCounterReset(entities []string) *CounterReset {
	serviceDomain := "counter"
	serviceType := "call_service"
	serviceService := "reset"
	c := &CounterReset{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CounterReset struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CounterReset) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CounterReset) SetID(id *int) {
	c.Id = id
}
