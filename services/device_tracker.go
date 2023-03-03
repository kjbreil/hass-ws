package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewDeviceTrackerSee creates the object that can be sent to Home Assistant for domain device_tracker, service see
// "Control tracked device."
func NewDeviceTrackerSee(entities []string, battery *float64, devId *string, gpsAccuracy *float64, hostName *string, locationName *string, mac *string) *DeviceTrackerSee {
	serviceDomain := "device_tracker"
	serviceType := "call_service"
	serviceService := "see"
	d := &DeviceTrackerSee{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Battery      *float64 `json:"battery,omitempty"`
			DevId        *string  `json:"dev_id,omitempty"`
			GpsAccuracy  *float64 `json:"gps_accuracy,omitempty"`
			HostName     *string  `json:"host_name,omitempty"`
			LocationName *string  `json:"location_name,omitempty"`
			Mac          *string  `json:"mac,omitempty"`
		}{
			Battery:      battery,
			DevId:        devId,
			GpsAccuracy:  gpsAccuracy,
			HostName:     hostName,
			LocationName: locationName,
			Mac:          mac,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return d
}

type DeviceTrackerSee struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Battery      *float64 `json:"battery,omitempty"`
		DevId        *string  `json:"dev_id,omitempty"`
		GpsAccuracy  *float64 `json:"gps_accuracy,omitempty"`
		HostName     *string  `json:"host_name,omitempty"`
		LocationName *string  `json:"location_name,omitempty"`
		Mac          *string  `json:"mac,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (d *DeviceTrackerSee) JSON() string {
	data, _ := json.Marshal(d)
	return string(data)
}
func (d *DeviceTrackerSee) SetID(id *int) {
	d.Id = id
}
