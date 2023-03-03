package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewDeviceTrackerSee creates the object that can be sent to Home Assistant for domain device_tracker, service see
// "Control tracked device."
func NewDeviceTrackerSee(target Target, deviceTrackerSeeParams DeviceTrackerSeeParams) *DeviceTrackerSee {
	serviceDomain := "device_tracker"
	serviceType := "call_service"
	serviceService := "see"
	d := &DeviceTrackerSee{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: deviceTrackerSeeParams,
	}
	return d
}

type DeviceTrackerSee struct {
	ServiceBase
	ServiceData DeviceTrackerSeeParams `json:"service_data,omitempty"`
}
type DeviceTrackerSeeParams struct {
	Battery      *float64 `json:"battery,omitempty"`
	DevId        *string  `json:"dev_id,omitempty"`
	GpsAccuracy  *float64 `json:"gps_accuracy,omitempty"`
	HostName     *string  `json:"host_name,omitempty"`
	LocationName *string  `json:"location_name,omitempty"`
	Mac          *string  `json:"mac,omitempty"`
}

func (d *DeviceTrackerSee) JSON() string {
	data, _ := json.Marshal(d)
	return string(data)
}
func (d *DeviceTrackerSee) SetID(id *int) {
	d.Id = id
}
