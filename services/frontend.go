package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewFrontendReloadThemes creates the object that can be sent to Home Assistant for domain frontend, service reload_themes
// "Reload themes from YAML configuration."
func NewFrontendReloadThemes(entities []string) *FrontendReloadThemes {
	serviceDomain := "frontend"
	serviceType := "call_service"
	serviceService := "reload_themes"
	f := &FrontendReloadThemes{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FrontendReloadThemes struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (f *FrontendReloadThemes) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FrontendReloadThemes) SetID(id *int) {
	f.Id = id
}

// NewFrontendSetTheme creates the object that can be sent to Home Assistant for domain frontend, service set_theme
// "Set a theme unless the client selected per-device theme."
func NewFrontendSetTheme(entities []string, mode *Mode) *FrontendSetTheme {
	serviceDomain := "frontend"
	serviceType := "call_service"
	serviceService := "set_theme"
	f := &FrontendSetTheme{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Mode *Mode `json:"mode,omitempty"`
		}{Mode: mode},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return f
}

type FrontendSetTheme struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Mode *Mode `json:"mode,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (f *FrontendSetTheme) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FrontendSetTheme) SetID(id *int) {
	f.Id = id
}
