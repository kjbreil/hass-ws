package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestCastShowLovelaceView_JSON(t *testing.T) {
	dashboardPath := "data"
	viewPath := "data"

	tests := []struct {
		name   string
		fields *CastShowLovelaceView
		want   string
	}{{
		fields: NewCastShowLovelaceView(Targets("climate.kitchen")).DashboardPath(dashboardPath).ViewPath(viewPath),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"cast\",\"service\":\"show_lovelace_view\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"dashboard_path\":\"data\",\"view_path\":\"data\"}}",
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
