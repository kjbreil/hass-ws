package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewFrontendReloadThemes creates the object that can be sent to Home Assistant for domain frontend, service reload_themes
// "Reload themes from YAML configuration."
func NewFrontendReloadThemes(target Target) *FrontendReloadThemes {
	serviceDomain := "frontend"
	serviceType := "call_service"
	serviceService := "reload_themes"
	f := &FrontendReloadThemes{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return f
}

type FrontendReloadThemes struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (f *FrontendReloadThemes) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FrontendReloadThemes) Name() string {
	return fmt.Sprintf("%s.%s", *f.Domain, *f.Service)
}
func (f *FrontendReloadThemes) SetID(id *int) {
	f.Id = id
}

// NewFrontendSetTheme creates the object that can be sent to Home Assistant for domain frontend, service set_theme
// "Set a theme unless the client selected per-device theme."
func NewFrontendSetTheme(target Target, frontendSetThemeParams *FrontendSetThemeParams) *FrontendSetTheme {
	serviceDomain := "frontend"
	serviceType := "call_service"
	serviceService := "set_theme"
	f := &FrontendSetTheme{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *frontendSetThemeParams,
	}
	return f
}

type FrontendSetTheme struct {
	ServiceBase
	ServiceData FrontendSetThemeParams `json:"service_data,omitempty"`
}
type FrontendSetThemeParams struct {
	Mode *Mode `json:"mode,omitempty"`
}

func (f *FrontendSetTheme) JSON() string {
	data, _ := json.Marshal(f)
	return string(data)
}
func (f *FrontendSetTheme) Name() string {
	return fmt.Sprintf("%s.%s", *f.Domain, *f.Service)
}
func (f *FrontendSetTheme) SetID(id *int) {
	f.Id = id
}
