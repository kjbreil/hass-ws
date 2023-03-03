package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewHumidifierSetHumidity creates the object that can be sent to Home Assistant for domain humidifier, service set_humidity
// "Set target humidity of humidifier device."
func NewHumidifierSetHumidity(target Target, humidifierSetHumidityParams *HumidifierSetHumidityParams) *HumidifierSetHumidity {
	serviceDomain := "humidifier"
	serviceType := "call_service"
	serviceService := "set_humidity"
	h := &HumidifierSetHumidity{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *humidifierSetHumidityParams,
	}
	return h
}

type HumidifierSetHumidity struct {
	ServiceBase
	ServiceData HumidifierSetHumidityParams `json:"service_data,omitempty"`
}
type HumidifierSetHumidityParams struct {
	Humidity *float64 `json:"humidity,omitempty"`
}

func (h *HumidifierSetHumidity) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HumidifierSetHumidity) SetID(id *int) {
	h.Id = id
}

// NewHumidifierSetMode creates the object that can be sent to Home Assistant for domain humidifier, service set_mode
// "Set mode for humidifier device."
func NewHumidifierSetMode(target Target, humidifierSetModeParams *HumidifierSetModeParams) *HumidifierSetMode {
	serviceDomain := "humidifier"
	serviceType := "call_service"
	serviceService := "set_mode"
	h := &HumidifierSetMode{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *humidifierSetModeParams,
	}
	return h
}

type HumidifierSetMode struct {
	ServiceBase
	ServiceData HumidifierSetModeParams `json:"service_data,omitempty"`
}
type HumidifierSetModeParams struct {
	Mode *string `json:"mode,omitempty"`
}

func (h *HumidifierSetMode) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HumidifierSetMode) SetID(id *int) {
	h.Id = id
}

// NewHumidifierToggle creates the object that can be sent to Home Assistant for domain humidifier, service toggle
// "Toggles a humidifier device."
func NewHumidifierToggle(target Target) *HumidifierToggle {
	serviceDomain := "humidifier"
	serviceType := "call_service"
	serviceService := "toggle"
	h := &HumidifierToggle{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return h
}

type HumidifierToggle struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (h *HumidifierToggle) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HumidifierToggle) SetID(id *int) {
	h.Id = id
}

// NewHumidifierTurnOff creates the object that can be sent to Home Assistant for domain humidifier, service turn_off
// "Turn humidifier device off."
func NewHumidifierTurnOff(target Target) *HumidifierTurnOff {
	serviceDomain := "humidifier"
	serviceType := "call_service"
	serviceService := "turn_off"
	h := &HumidifierTurnOff{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return h
}

type HumidifierTurnOff struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (h *HumidifierTurnOff) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HumidifierTurnOff) SetID(id *int) {
	h.Id = id
}

// NewHumidifierTurnOn creates the object that can be sent to Home Assistant for domain humidifier, service turn_on
// "Turn humidifier device on."
func NewHumidifierTurnOn(target Target) *HumidifierTurnOn {
	serviceDomain := "humidifier"
	serviceType := "call_service"
	serviceService := "turn_on"
	h := &HumidifierTurnOn{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return h
}

type HumidifierTurnOn struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (h *HumidifierTurnOn) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HumidifierTurnOn) SetID(id *int) {
	h.Id = id
}
