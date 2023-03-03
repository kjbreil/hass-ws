package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCastShowLovelaceView creates the object that can be sent to Home Assistant for domain cast, service show_lovelace_view
// "Show a Lovelace view on a Chromecast."
func NewCastShowLovelaceView(target Target, castShowLovelaceViewParams CastShowLovelaceViewParams) *CastShowLovelaceView {
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
		ServiceData: castShowLovelaceViewParams,
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

func (c *CastShowLovelaceView) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CastShowLovelaceView) SetID(id *int) {
	c.Id = id
}
