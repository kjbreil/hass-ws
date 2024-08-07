package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
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
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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
	data, _ := gojson.Marshal(u)
	return string(data)
}
func (u *UtilityMeterCalibrate) Targets() []string {
	return u.Target.EntityId
}
func (u *UtilityMeterCalibrate) Name() string {
	return fmt.Sprintf("%s.%s", *u.Domain, *u.Service)
}

// NewUtilityMeterReset creates the object that can be sent to Home Assistant for domain utility_meter, service reset
// "Resets all counters of a utility meter."
func NewUtilityMeterReset(target Target) *UtilityMeterReset {
	serviceDomain := "utility_meter"
	serviceType := "call_service"
	serviceService := "reset"
	u := &UtilityMeterReset{
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
	return u
}

type UtilityMeterReset struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (u *UtilityMeterReset) JSON() string {
	data, _ := gojson.Marshal(u)
	return string(data)
}
func (u *UtilityMeterReset) Targets() []string {
	return u.Target.EntityId
}
func (u *UtilityMeterReset) Name() string {
	return fmt.Sprintf("%s.%s", *u.Domain, *u.Service)
}
