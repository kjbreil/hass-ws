package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewUtilityMeterCalibrate creates the object that can be sent to Home Assistant for domain utility_meter, service calibrate
// "Calibrates a utility meter sensor."
func NewUtilityMeterCalibrate(entities []string, value *string) *UtilityMeterCalibrate {
	serviceDomain := "utility_meter"
	serviceType := "call_service"
	serviceService := "calibrate"
	u := &UtilityMeterCalibrate{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Value *string `json:"value,omitempty"`
		}{Value: value},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return u
}

type UtilityMeterCalibrate struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Value *string `json:"value,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewUtilityMeterReset(entities []string) *UtilityMeterReset {
	serviceDomain := "utility_meter"
	serviceType := "call_service"
	serviceService := "reset"
	u := &UtilityMeterReset{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return u
}

type UtilityMeterReset struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (u *UtilityMeterReset) JSON() string {
	data, _ := json.Marshal(u)
	return string(data)
}
func (u *UtilityMeterReset) SetID(id *int) {
	u.Id = id
}
