package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCoverCloseCover creates the object that can be sent to Home Assistant for domain cover, service close_cover
// "Close all or specified cover."
func NewCoverCloseCover(entities []string) *CoverCloseCover {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "close_cover"
	c := &CoverCloseCover{
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

type CoverCloseCover struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CoverCloseCover) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverCloseCover) SetID(id *int) {
	c.Id = id
}

// NewCoverCloseCoverTilt creates the object that can be sent to Home Assistant for domain cover, service close_cover_tilt
// "Close all or specified cover tilt."
func NewCoverCloseCoverTilt(entities []string) *CoverCloseCoverTilt {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "close_cover_tilt"
	c := &CoverCloseCoverTilt{
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

type CoverCloseCoverTilt struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CoverCloseCoverTilt) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverCloseCoverTilt) SetID(id *int) {
	c.Id = id
}

// NewCoverOpenCover creates the object that can be sent to Home Assistant for domain cover, service open_cover
// "Open all or specified cover."
func NewCoverOpenCover(entities []string) *CoverOpenCover {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "open_cover"
	c := &CoverOpenCover{
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

type CoverOpenCover struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CoverOpenCover) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverOpenCover) SetID(id *int) {
	c.Id = id
}

// NewCoverOpenCoverTilt creates the object that can be sent to Home Assistant for domain cover, service open_cover_tilt
// "Open all or specified cover tilt."
func NewCoverOpenCoverTilt(entities []string) *CoverOpenCoverTilt {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "open_cover_tilt"
	c := &CoverOpenCoverTilt{
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

type CoverOpenCoverTilt struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CoverOpenCoverTilt) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverOpenCoverTilt) SetID(id *int) {
	c.Id = id
}

// NewCoverSetCoverPosition creates the object that can be sent to Home Assistant for domain cover, service set_cover_position
// "Move to specific position all or specified cover."
func NewCoverSetCoverPosition(entities []string, position *float64) *CoverSetCoverPosition {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "set_cover_position"
	c := &CoverSetCoverPosition{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Position *float64 `json:"position,omitempty"`
		}{Position: position},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CoverSetCoverPosition struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Position *float64 `json:"position,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewCoverSetCoverTiltPosition(entities []string, tiltPosition *float64) *CoverSetCoverTiltPosition {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "set_cover_tilt_position"
	c := &CoverSetCoverTiltPosition{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			TiltPosition *float64 `json:"tilt_position,omitempty"`
		}{TiltPosition: tiltPosition},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CoverSetCoverTiltPosition struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		TiltPosition *float64 `json:"tilt_position,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewCoverStopCover(entities []string) *CoverStopCover {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "stop_cover"
	c := &CoverStopCover{
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

type CoverStopCover struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CoverStopCover) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverStopCover) SetID(id *int) {
	c.Id = id
}

// NewCoverStopCoverTilt creates the object that can be sent to Home Assistant for domain cover, service stop_cover_tilt
// "Stop all or specified cover."
func NewCoverStopCoverTilt(entities []string) *CoverStopCoverTilt {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "stop_cover_tilt"
	c := &CoverStopCoverTilt{
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

type CoverStopCoverTilt struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CoverStopCoverTilt) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverStopCoverTilt) SetID(id *int) {
	c.Id = id
}

// NewCoverToggle creates the object that can be sent to Home Assistant for domain cover, service toggle
// "Toggle a cover open/closed."
func NewCoverToggle(entities []string) *CoverToggle {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "toggle"
	c := &CoverToggle{
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

type CoverToggle struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CoverToggle) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverToggle) SetID(id *int) {
	c.Id = id
}

// NewCoverToggleCoverTilt creates the object that can be sent to Home Assistant for domain cover, service toggle_cover_tilt
// "Toggle a cover tilt open/closed."
func NewCoverToggleCoverTilt(entities []string) *CoverToggleCoverTilt {
	serviceDomain := "cover"
	serviceType := "call_service"
	serviceService := "toggle_cover_tilt"
	c := &CoverToggleCoverTilt{
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

type CoverToggleCoverTilt struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CoverToggleCoverTilt) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CoverToggleCoverTilt) SetID(id *int) {
	c.Id = id
}
