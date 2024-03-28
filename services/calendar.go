package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCalendarCreateEvent creates the object that can be sent to Home Assistant for domain calendar, service create_event
// "Adds a new calendar event."
func NewCalendarCreateEvent(target Target) *CalendarCreateEvent {
	serviceDomain := "calendar"
	serviceType := "call_service"
	serviceService := "create_event"
	c := &CalendarCreateEvent{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: CalendarCreateEventParams{},
	}
	return c
}

type CalendarCreateEvent struct {
	ServiceBase
	ServiceData CalendarCreateEventParams `json:"service_data,omitempty"`
}
type CalendarCreateEventParams struct {
	Description *string `json:"description,omitempty"`
	Location    *string `json:"location,omitempty"`
	Summary     *string `json:"summary,omitempty"`
}

func (c *CalendarCreateEvent) Description(description string) *CalendarCreateEvent {
	c.ServiceData.Description = &description
	return c
}
func (c *CalendarCreateEvent) Location(location string) *CalendarCreateEvent {
	c.ServiceData.Location = &location
	return c
}
func (c *CalendarCreateEvent) Summary(summary string) *CalendarCreateEvent {
	c.ServiceData.Summary = &summary
	return c
}
func (c *CalendarCreateEvent) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CalendarCreateEvent) Targets() []string {
	return c.Target.EntityId
}
func (c *CalendarCreateEvent) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}
func (c *CalendarCreateEvent) SetID(id *int) {
	c.Id = id
}

// NewCalendarListEvents creates the object that can be sent to Home Assistant for domain calendar, service list_events
// "Lists events on a calendar within a time range."
func NewCalendarListEvents(target Target) *CalendarListEvents {
	serviceDomain := "calendar"
	serviceType := "call_service"
	serviceService := "list_events"
	c := &CalendarListEvents{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type CalendarListEvents struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CalendarListEvents) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CalendarListEvents) Targets() []string {
	return c.Target.EntityId
}
func (c *CalendarListEvents) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}
func (c *CalendarListEvents) SetID(id *int) {
	c.Id = id
}
