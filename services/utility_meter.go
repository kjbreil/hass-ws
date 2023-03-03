package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewUtilityMeterCalibrate creates the object that can be sent to Home Assistant for domain utility_meter, service calibrate
// "Calibrates a utility meter sensor."
func NewUtilityMeterCalibrate(target Target, utilityMeterCalibrateParams UtilityMeterCalibrateParams) *UtilityMeterCalibrate {
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
		ServiceData: utilityMeterCalibrateParams,
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

func (u *UtilityMeterCalibrate) JSON() string {
	data, _ := json.Marshal(u)
	return string(data)
}
func (u *UtilityMeterCalibrate) SetID(id *int) {
	u.Id = id
}

// NewUtilityMeterReset creates the object that can be sent to Home Assistant for domain utility_meter, service reset
// "Resets all counters of a utility meter."
func NewUtilityMeterReset(target Target, utilityMeterResetParams UtilityMeterResetParams) *UtilityMeterReset {
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
		ServiceData: utilityMeterResetParams,
	}
	return u
}

type UtilityMeterReset struct {
	ServiceBase
	ServiceData UtilityMeterResetParams `json:"service_data,omitempty"`
}
type UtilityMeterResetParams struct{}

func (u *UtilityMeterReset) JSON() string {
	data, _ := json.Marshal(u)
	return string(data)
}
func (u *UtilityMeterReset) SetID(id *int) {
	u.Id = id
}
