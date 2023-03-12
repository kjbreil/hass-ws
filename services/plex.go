package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewPlexRefreshLibrary creates the object that can be sent to Home Assistant for domain plex, service refresh_library
// "Refresh a Plex library to scan for new and updated media."
func NewPlexRefreshLibrary(target Target, plexRefreshLibraryParams *PlexRefreshLibraryParams) *PlexRefreshLibrary {
	serviceDomain := "plex"
	serviceType := "call_service"
	serviceService := "refresh_library"
	p := &PlexRefreshLibrary{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *plexRefreshLibraryParams,
	}
	return p
}

type PlexRefreshLibrary struct {
	ServiceBase
	ServiceData PlexRefreshLibraryParams `json:"service_data,omitempty"`
}
type PlexRefreshLibraryParams struct {
	LibraryName *string `json:"library_name,omitempty"`
	ServerName  *string `json:"server_name,omitempty"`
}

func (p *PlexRefreshLibrary) JSON() string {
	data, _ := json.Marshal(p)
	return string(data)
}
func (p *PlexRefreshLibrary) Name() string {
	return fmt.Sprintf("%s.%s", *p.Domain, *p.Service)
}
func (p *PlexRefreshLibrary) SetID(id *int) {
	p.Id = id
}

// NewPlexScanForClients creates the object that can be sent to Home Assistant for domain plex, service scan_for_clients
// "Scan for available clients from the Plex server(s), local network, and plex.tv."
func NewPlexScanForClients(target Target) *PlexScanForClients {
	serviceDomain := "plex"
	serviceType := "call_service"
	serviceService := "scan_for_clients"
	p := &PlexScanForClients{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return p
}

type PlexScanForClients struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (p *PlexScanForClients) JSON() string {
	data, _ := json.Marshal(p)
	return string(data)
}
func (p *PlexScanForClients) Name() string {
	return fmt.Sprintf("%s.%s", *p.Domain, *p.Service)
}
func (p *PlexScanForClients) SetID(id *int) {
	p.Id = id
}
