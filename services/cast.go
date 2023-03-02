package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCastShowLovelaceView creates the object that can be sent to Home Assistant for domain cast, service show_lovelace_view
// "Show a Lovelace view on a Chromecast."
func NewCastShowLovelaceView(entities []string, dashboardPath *string, viewPath *string) *CastShowLovelaceView {
	serviceDomain := "cast"
	serviceType := "call_service"
	serviceService := "show_lovelace_view"
	c := &CastShowLovelaceView{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			DashboardPath *string `json:"dashboard_path,omitempty"`
			ViewPath      *string `json:"view_path,omitempty"`
		}{
			DashboardPath: dashboardPath,
			ViewPath:      viewPath,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CastShowLovelaceView struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		DashboardPath *string `json:"dashboard_path,omitempty"`
		ViewPath      *string `json:"view_path,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CastShowLovelaceView) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CastShowLovelaceView) SetID(id *int) {
	c.Id = id
}
