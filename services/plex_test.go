package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestPlexRefreshLibrary_JSON(t *testing.T) {
	libraryName := "data"
	serverName := "data"

	tests := []struct {
		name   string
		fields *PlexRefreshLibrary
		want   string
	}{{
		fields: NewPlexRefreshLibrary(Targets("climate.kitchen"), &PlexRefreshLibraryParams{
			LibraryName: &libraryName,
			ServerName:  &serverName,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"plex\",\"service\":\"refresh_library\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"library_name\":\"data\",\"server_name\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestPlexScanForClients_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *PlexScanForClients
		want   string
	}{{
		fields: NewPlexScanForClients(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"plex\",\"service\":\"scan_for_clients\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
