package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewUtilityMeterCalibrate creates the object that can be sent to Home Assistant for domain utility_meter, service calibrate
// "Calibrates a utility meter sensor."
func NewUtilityMeterCalibrate(target Target) *UtilityMeterCalibrate {
	serviceDomain := "utility_meter"
	serviceType := "call_service"
	serviceService := "calibrate"
	u := &UtilityMeterCalibrate{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: UtilityMeterCalibrateParams{},
	}
	return u
}

type UtilityMeterCalibrate struct {
	ServiceBase
	ServiceData UtilityMeterCalibrateParams `json:"service_data,omitempty"`
}
type UtilityMeterCalibrateParams struct {
	Value *string `json:"value,omitempty"`
}

func (u *UtilityMeterCalibrate) Value(value string) *UtilityMeterCalibrate {
	u.ServiceData.Value = &value
	return u
}
func (u *UtilityMeterCalibrate) JSON() string {
	data, _ := json.Marshal(u)
	return string(data)
}
func (u *UtilityMeterCalibrate) Name() string {
	return fmt.Sprintf("%s.%s", *u.Domain, *u.Service)
}
func (u *UtilityMeterCalibrate) SetID(id *int) {
	u.Id = id
}

// NewUtilityMeterReset creates the object that can be sent to Home Assistant for domain utility_meter, service reset
// "Resets all counters of a utility meter."
func NewUtilityMeterReset(target Target) *UtilityMeterReset {
	serviceDomain := "utility_meter"
	serviceType := "call_service"
	serviceService := "reset"
	u := &UtilityMeterReset{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return u
}

type UtilityMeterReset struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (u *UtilityMeterReset) JSON() string {
	data, _ := json.Marshal(u)
	return string(data)
}
func (u *UtilityMeterReset) Name() string {
	return fmt.Sprintf("%s.%s", *u.Domain, *u.Service)
}
func (u *UtilityMeterReset) SetID(id *int) {
	u.Id = id
}
