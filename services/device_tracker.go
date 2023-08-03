package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewDeviceTrackerSee creates the object that can be sent to Home Assistant for domain device_tracker, service see
// "Control tracked device."
func NewDeviceTrackerSee(target Target) *DeviceTrackerSee {
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
		ServiceData: DeviceTrackerSeeParams{},
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

func (d *DeviceTrackerSee) Battery(battery float64) *DeviceTrackerSee {
	d.ServiceData.Battery = &battery
	return d
}
func (d *DeviceTrackerSee) DevId(devId string) *DeviceTrackerSee {
	d.ServiceData.DevId = &devId
	return d
}
func (d *DeviceTrackerSee) GpsAccuracy(gpsAccuracy float64) *DeviceTrackerSee {
	d.ServiceData.GpsAccuracy = &gpsAccuracy
	return d
}
func (d *DeviceTrackerSee) HostName(hostName string) *DeviceTrackerSee {
	d.ServiceData.HostName = &hostName
	return d
}
func (d *DeviceTrackerSee) LocationName(locationName string) *DeviceTrackerSee {
	d.ServiceData.LocationName = &locationName
	return d
}
func (d *DeviceTrackerSee) Mac(mac string) *DeviceTrackerSee {
	d.ServiceData.Mac = &mac
	return d
}
func (d *DeviceTrackerSee) JSON() string {
	data, _ := gojson.Marshal(d)
	return string(data)
}
func (d *DeviceTrackerSee) Targets() []string {
	return d.Target.EntityId
}
func (d *DeviceTrackerSee) Name() string {
	return fmt.Sprintf("%s.%s", *d.Domain, *d.Service)
}
func (d *DeviceTrackerSee) SetID(id *int) {
	d.Id = id
}
