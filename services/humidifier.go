package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewHumidifierSetHumidity creates the object that can be sent to Home Assistant for domain humidifier, service set_humidity
// "Set target humidity of humidifier device."
func NewHumidifierSetHumidity(entities []string, humidity *float64) *HumidifierSetHumidity {
	serviceDomain := "humidifier"
	serviceType := "call_service"
	serviceService := "set_humidity"
	h := &HumidifierSetHumidity{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Humidity *float64 `json:"humidity,omitempty"`
		}{Humidity: humidity},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return h
}

type HumidifierSetHumidity struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Humidity *float64 `json:"humidity,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewHumidifierSetMode(entities []string, mode *string) *HumidifierSetMode {
	serviceDomain := "humidifier"
	serviceType := "call_service"
	serviceService := "set_mode"
	h := &HumidifierSetMode{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Mode *string `json:"mode,omitempty"`
		}{Mode: mode},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return h
}

type HumidifierSetMode struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Mode *string `json:"mode,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewHumidifierToggle(entities []string) *HumidifierToggle {
	serviceDomain := "humidifier"
	serviceType := "call_service"
	serviceService := "toggle"
	h := &HumidifierToggle{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return h
}

type HumidifierToggle struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewHumidifierTurnOff(entities []string) *HumidifierTurnOff {
	serviceDomain := "humidifier"
	serviceType := "call_service"
	serviceService := "turn_off"
	h := &HumidifierTurnOff{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return h
}

type HumidifierTurnOff struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
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
func NewHumidifierTurnOn(entities []string) *HumidifierTurnOn {
	serviceDomain := "humidifier"
	serviceType := "call_service"
	serviceService := "turn_on"
	h := &HumidifierTurnOn{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return h
}

type HumidifierTurnOn struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (h *HumidifierTurnOn) JSON() string {
	data, _ := json.Marshal(h)
	return string(data)
}
func (h *HumidifierTurnOn) SetID(id *int) {
	h.Id = id
}
