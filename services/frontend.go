package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
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
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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
	data, _ := gojson.Marshal(f)
	return string(data)
}
func (f *FrontendReloadThemes) Targets() []string {
	return f.Target.EntityId
}
func (f *FrontendReloadThemes) Name() string {
	return fmt.Sprintf("%s.%s", *f.Domain, *f.Service)
}

// NewFrontendSetTheme creates the object that can be sent to Home Assistant for domain frontend, service set_theme
// "Set a theme unless the client selected per-device theme."
func NewFrontendSetTheme(target Target) *FrontendSetTheme {
	serviceDomain := "frontend"
	serviceType := "call_service"
	serviceService := "set_theme"
	f := &FrontendSetTheme{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: FrontendSetThemeParams{},
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

func (f *FrontendSetTheme) Mode(mode Mode) *FrontendSetTheme {
	f.ServiceData.Mode = &mode
	return f
}
func (f *FrontendSetTheme) JSON() string {
	data, _ := gojson.Marshal(f)
	return string(data)
}
func (f *FrontendSetTheme) Targets() []string {
	return f.Target.EntityId
}
func (f *FrontendSetTheme) Name() string {
	return fmt.Sprintf("%s.%s", *f.Domain, *f.Service)
}
