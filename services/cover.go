package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCoverCloseCover creates the object that can be sent to Home Assistant for domain cover, service close_cover
// "Close all or specified cover."
func NewCoverCloseCover(target Target) *CoverCloseCover {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "close_cover"
	c := &CoverCloseCover{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type CoverCloseCover struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CoverCloseCover) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CoverCloseCover) Targets() []string {
	return c.Target.EntityId
}
func (c *CoverCloseCover) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCoverCloseCoverTilt creates the object that can be sent to Home Assistant for domain cover, service close_cover_tilt
// "Close all or specified cover tilt."
func NewCoverCloseCoverTilt(target Target) *CoverCloseCoverTilt {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "close_cover_tilt"
	c := &CoverCloseCoverTilt{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type CoverCloseCoverTilt struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CoverCloseCoverTilt) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CoverCloseCoverTilt) Targets() []string {
	return c.Target.EntityId
}
func (c *CoverCloseCoverTilt) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCoverOpenCover creates the object that can be sent to Home Assistant for domain cover, service open_cover
// "Open all or specified cover."
func NewCoverOpenCover(target Target) *CoverOpenCover {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "open_cover"
	c := &CoverOpenCover{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type CoverOpenCover struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CoverOpenCover) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CoverOpenCover) Targets() []string {
	return c.Target.EntityId
}
func (c *CoverOpenCover) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCoverOpenCoverTilt creates the object that can be sent to Home Assistant for domain cover, service open_cover_tilt
// "Open all or specified cover tilt."
func NewCoverOpenCoverTilt(target Target) *CoverOpenCoverTilt {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "open_cover_tilt"
	c := &CoverOpenCoverTilt{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type CoverOpenCoverTilt struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CoverOpenCoverTilt) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CoverOpenCoverTilt) Targets() []string {
	return c.Target.EntityId
}
func (c *CoverOpenCoverTilt) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCoverSetCoverPosition creates the object that can be sent to Home Assistant for domain cover, service set_cover_position
// "Move to specific position all or specified cover."
func NewCoverSetCoverPosition(target Target) *CoverSetCoverPosition {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "set_cover_position"
	c := &CoverSetCoverPosition{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: CoverSetCoverPositionParams{},
	}
	return c
}

type CoverSetCoverPosition struct {
	ServiceBase
	ServiceData CoverSetCoverPositionParams `json:"service_data,omitempty"`
}
type CoverSetCoverPositionParams struct {
	Position *float64 `json:"position,omitempty"`
}

func (c *CoverSetCoverPosition) Position(position float64) *CoverSetCoverPosition {
	c.ServiceData.Position = &position
	return c
}
func (c *CoverSetCoverPosition) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CoverSetCoverPosition) Targets() []string {
	return c.Target.EntityId
}
func (c *CoverSetCoverPosition) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCoverSetCoverTiltPosition creates the object that can be sent to Home Assistant for domain cover, service set_cover_tilt_position
// "Move to specific position all or specified cover tilt."
func NewCoverSetCoverTiltPosition(target Target) *CoverSetCoverTiltPosition {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "set_cover_tilt_position"
	c := &CoverSetCoverTiltPosition{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: CoverSetCoverTiltPositionParams{},
	}
	return c
}

type CoverSetCoverTiltPosition struct {
	ServiceBase
	ServiceData CoverSetCoverTiltPositionParams `json:"service_data,omitempty"`
}
type CoverSetCoverTiltPositionParams struct {
	TiltPosition *float64 `json:"tilt_position,omitempty"`
}

func (c *CoverSetCoverTiltPosition) TiltPosition(tiltPosition float64) *CoverSetCoverTiltPosition {
	c.ServiceData.TiltPosition = &tiltPosition
	return c
}
func (c *CoverSetCoverTiltPosition) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CoverSetCoverTiltPosition) Targets() []string {
	return c.Target.EntityId
}
func (c *CoverSetCoverTiltPosition) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCoverStopCover creates the object that can be sent to Home Assistant for domain cover, service stop_cover
// "Stop all or specified cover."
func NewCoverStopCover(target Target) *CoverStopCover {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "stop_cover"
	c := &CoverStopCover{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type CoverStopCover struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CoverStopCover) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CoverStopCover) Targets() []string {
	return c.Target.EntityId
}
func (c *CoverStopCover) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCoverStopCoverTilt creates the object that can be sent to Home Assistant for domain cover, service stop_cover_tilt
// "Stop all or specified cover."
func NewCoverStopCoverTilt(target Target) *CoverStopCoverTilt {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "stop_cover_tilt"
	c := &CoverStopCoverTilt{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type CoverStopCoverTilt struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CoverStopCoverTilt) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CoverStopCoverTilt) Targets() []string {
	return c.Target.EntityId
}
func (c *CoverStopCoverTilt) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCoverToggle creates the object that can be sent to Home Assistant for domain cover, service toggle
// "Toggle a cover open/closed."
func NewCoverToggle(target Target) *CoverToggle {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "toggle"
	c := &CoverToggle{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type CoverToggle struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CoverToggle) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CoverToggle) Targets() []string {
	return c.Target.EntityId
}
func (c *CoverToggle) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCoverToggleCoverTilt creates the object that can be sent to Home Assistant for domain cover, service toggle_cover_tilt
// "Toggle a cover tilt open/closed."
func NewCoverToggleCoverTilt(target Target) *CoverToggleCoverTilt {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "toggle_cover_tilt"
	c := &CoverToggleCoverTilt{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: nil,
	}
	return c
}

type CoverToggleCoverTilt struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CoverToggleCoverTilt) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CoverToggleCoverTilt) Targets() []string {
	return c.Target.EntityId
}
func (c *CoverToggleCoverTilt) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}
