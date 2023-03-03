package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCoverCloseCover creates the object that can be sent to Home Assistant for domain cover, service close_cover
// "Close all or specified cover."
func NewCoverCloseCover(target Target, coverCloseCoverParams CoverCloseCoverParams) *CoverCloseCover {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "close_cover"
	c := &CoverCloseCover{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: coverCloseCoverParams,
	}
	return c
}

type CoverCloseCover struct {
	ServiceBase
	ServiceData CoverCloseCoverParams `json:"service_data,omitempty"`
}
type CoverCloseCoverParams struct{}

func (c *CoverCloseCover) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverCloseCover) SetID(id *int) {
	c.Id = id
}

// NewCoverCloseCoverTilt creates the object that can be sent to Home Assistant for domain cover, service close_cover_tilt
// "Close all or specified cover tilt."
func NewCoverCloseCoverTilt(target Target, coverCloseCoverTiltParams CoverCloseCoverTiltParams) *CoverCloseCoverTilt {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "close_cover_tilt"
	c := &CoverCloseCoverTilt{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: coverCloseCoverTiltParams,
	}
	return c
}

type CoverCloseCoverTilt struct {
	ServiceBase
	ServiceData CoverCloseCoverTiltParams `json:"service_data,omitempty"`
}
type CoverCloseCoverTiltParams struct{}

func (c *CoverCloseCoverTilt) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverCloseCoverTilt) SetID(id *int) {
	c.Id = id
}

// NewCoverOpenCover creates the object that can be sent to Home Assistant for domain cover, service open_cover
// "Open all or specified cover."
func NewCoverOpenCover(target Target, coverOpenCoverParams CoverOpenCoverParams) *CoverOpenCover {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "open_cover"
	c := &CoverOpenCover{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: coverOpenCoverParams,
	}
	return c
}

type CoverOpenCover struct {
	ServiceBase
	ServiceData CoverOpenCoverParams `json:"service_data,omitempty"`
}
type CoverOpenCoverParams struct{}

func (c *CoverOpenCover) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverOpenCover) SetID(id *int) {
	c.Id = id
}

// NewCoverOpenCoverTilt creates the object that can be sent to Home Assistant for domain cover, service open_cover_tilt
// "Open all or specified cover tilt."
func NewCoverOpenCoverTilt(target Target, coverOpenCoverTiltParams CoverOpenCoverTiltParams) *CoverOpenCoverTilt {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "open_cover_tilt"
	c := &CoverOpenCoverTilt{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: coverOpenCoverTiltParams,
	}
	return c
}

type CoverOpenCoverTilt struct {
	ServiceBase
	ServiceData CoverOpenCoverTiltParams `json:"service_data,omitempty"`
}
type CoverOpenCoverTiltParams struct{}

func (c *CoverOpenCoverTilt) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverOpenCoverTilt) SetID(id *int) {
	c.Id = id
}

// NewCoverSetCoverPosition creates the object that can be sent to Home Assistant for domain cover, service set_cover_position
// "Move to specific position all or specified cover."
func NewCoverSetCoverPosition(target Target, coverSetCoverPositionParams CoverSetCoverPositionParams) *CoverSetCoverPosition {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "set_cover_position"
	c := &CoverSetCoverPosition{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: coverSetCoverPositionParams,
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

func (c *CoverSetCoverPosition) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverSetCoverPosition) SetID(id *int) {
	c.Id = id
}

// NewCoverSetCoverTiltPosition creates the object that can be sent to Home Assistant for domain cover, service set_cover_tilt_position
// "Move to specific position all or specified cover tilt."
func NewCoverSetCoverTiltPosition(target Target, coverSetCoverTiltPositionParams CoverSetCoverTiltPositionParams) *CoverSetCoverTiltPosition {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "set_cover_tilt_position"
	c := &CoverSetCoverTiltPosition{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: coverSetCoverTiltPositionParams,
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

func (c *CoverSetCoverTiltPosition) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverSetCoverTiltPosition) SetID(id *int) {
	c.Id = id
}

// NewCoverStopCover creates the object that can be sent to Home Assistant for domain cover, service stop_cover
// "Stop all or specified cover."
func NewCoverStopCover(target Target, coverStopCoverParams CoverStopCoverParams) *CoverStopCover {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "stop_cover"
	c := &CoverStopCover{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: coverStopCoverParams,
	}
	return c
}

type CoverStopCover struct {
	ServiceBase
	ServiceData CoverStopCoverParams `json:"service_data,omitempty"`
}
type CoverStopCoverParams struct{}

func (c *CoverStopCover) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverStopCover) SetID(id *int) {
	c.Id = id
}

// NewCoverStopCoverTilt creates the object that can be sent to Home Assistant for domain cover, service stop_cover_tilt
// "Stop all or specified cover."
func NewCoverStopCoverTilt(target Target, coverStopCoverTiltParams CoverStopCoverTiltParams) *CoverStopCoverTilt {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "stop_cover_tilt"
	c := &CoverStopCoverTilt{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: coverStopCoverTiltParams,
	}
	return c
}

type CoverStopCoverTilt struct {
	ServiceBase
	ServiceData CoverStopCoverTiltParams `json:"service_data,omitempty"`
}
type CoverStopCoverTiltParams struct{}

func (c *CoverStopCoverTilt) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverStopCoverTilt) SetID(id *int) {
	c.Id = id
}

// NewCoverToggle creates the object that can be sent to Home Assistant for domain cover, service toggle
// "Toggle a cover open/closed."
func NewCoverToggle(target Target, coverToggleParams CoverToggleParams) *CoverToggle {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "toggle"
	c := &CoverToggle{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: coverToggleParams,
	}
	return c
}

type CoverToggle struct {
	ServiceBase
	ServiceData CoverToggleParams `json:"service_data,omitempty"`
}
type CoverToggleParams struct{}

func (c *CoverToggle) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverToggle) SetID(id *int) {
	c.Id = id
}

// NewCoverToggleCoverTilt creates the object that can be sent to Home Assistant for domain cover, service toggle_cover_tilt
// "Toggle a cover tilt open/closed."
func NewCoverToggleCoverTilt(target Target, coverToggleCoverTiltParams CoverToggleCoverTiltParams) *CoverToggleCoverTilt {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "toggle_cover_tilt"
	c := &CoverToggleCoverTilt{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: coverToggleCoverTiltParams,
	}
	return c
}

type CoverToggleCoverTilt struct {
	ServiceBase
	ServiceData CoverToggleCoverTiltParams `json:"service_data,omitempty"`
}
type CoverToggleCoverTiltParams struct{}

func (c *CoverToggleCoverTilt) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverToggleCoverTilt) SetID(id *int) {
	c.Id = id
}
