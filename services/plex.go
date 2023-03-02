package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewPlexRefreshLibrary creates the object that can be sent to Home Assistant for domain plex, service refresh_library
// "Refresh a Plex library to scan for new and updated media."
func NewPlexRefreshLibrary(entities []string, libraryName *string, serverName *string) *PlexRefreshLibrary {
	serviceDomain := "plex"
	serviceType := "call_service"
	serviceService := "refresh_library"
	p := &PlexRefreshLibrary{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			LibraryName *string `json:"library_name,omitempty"`
			ServerName  *string `json:"server_name,omitempty"`
		}{
			LibraryName: libraryName,
			ServerName:  serverName,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return p
}

type PlexRefreshLibrary struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		LibraryName *string `json:"library_name,omitempty"`
		ServerName  *string `json:"server_name,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (p *PlexRefreshLibrary) JSON() string {
	data, _ := json.Marshal(p)
	return string(data)
}
func (p *PlexRefreshLibrary) SetID(id *int) {
	p.Id = id
}

// NewPlexScanForClients creates the object that can be sent to Home Assistant for domain plex, service scan_for_clients
// "Scan for available clients from the Plex server(s), local network, and plex.tv."
func NewPlexScanForClients(entities []string) *PlexScanForClients {
	serviceDomain := "plex"
	serviceType := "call_service"
	serviceService := "scan_for_clients"
	p := &PlexScanForClients{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return p
}

type PlexScanForClients struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (p *PlexScanForClients) JSON() string {
	data, _ := json.Marshal(p)
	return string(data)
}
func (p *PlexScanForClients) SetID(id *int) {
	p.Id = id
}
