package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCastShowLovelaceView creates the object that can be sent to Home Assistant for domain cast, service show_lovelace_view
// "Show a Lovelace view on a Chromecast."
func NewCastShowLovelaceView(target Target) *CastShowLovelaceView {
	serviceDomain := "cast"
	serviceType := "call_service"
	serviceService := "show_lovelace_view"
	c := &CastShowLovelaceView{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: CastShowLovelaceViewParams{},
	}
	return c
}

type CastShowLovelaceView struct {
	ServiceBase
	ServiceData CastShowLovelaceViewParams `json:"service_data,omitempty"`
}
type CastShowLovelaceViewParams struct {
	DashboardPath *string `json:"dashboard_path,omitempty"`
	ViewPath      *string `json:"view_path,omitempty"`
}

func (c *CastShowLovelaceView) DashboardPath(dashboardPath string) *CastShowLovelaceView {
	c.ServiceData.DashboardPath = &dashboardPath
	return c
}
func (c *CastShowLovelaceView) ViewPath(viewPath string) *CastShowLovelaceView {
	c.ServiceData.ViewPath = &viewPath
	return c
}
func (c *CastShowLovelaceView) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CastShowLovelaceView) Targets() []string {
	return c.Target.EntityId
}
func (c *CastShowLovelaceView) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}
func (c *CastShowLovelaceView) SetID(id *int) {
	c.Id = id
}
