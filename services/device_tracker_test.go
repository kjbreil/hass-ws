package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestDeviceTrackerSee_JSON(t *testing.T) {
	battery := 1.2
	devId := "data"
	gpsAccuracy := 1.2
	hostName := "data"
	locationName := "data"
	mac := "data"

	tests := []struct {
		name   string
		fields *DeviceTrackerSee
		want   string
	}{{
		fields: NewDeviceTrackerSee(Targets("climate.kitchen"), &DeviceTrackerSeeParams{
			Battery:      &battery,
			DevId:        &devId,
			GpsAccuracy:  &gpsAccuracy,
			HostName:     &hostName,
			LocationName: &locationName,
			Mac:          &mac,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"device_tracker\",\"service\":\"see\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"battery\":1.2,\"dev_id\":\"data\",\"gps_accuracy\":1.2,\"host_name\":\"data\",\"location_name\":\"data\",\"mac\":\"data\"}}",
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
