package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewMqttDump creates the object that can be sent to Home Assistant for domain mqtt, service dump
// "Dump messages on a topic selector to the 'mqtt_dump.txt' file in your configuration folder."
func NewMqttDump(entities []string, duration *float64, topic *string) *MqttDump {
	serviceDomain := "mqtt"
	serviceType := "call_service"
	serviceService := "dump"
	m := &MqttDump{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Duration *float64 `json:"duration,omitempty"`
			Topic    *string  `json:"topic,omitempty"`
		}{
			Duration: duration,
			Topic:    topic,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return m
}

type MqttDump struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Duration *float64 `json:"duration,omitempty"`
		Topic    *string  `json:"topic,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MqttDump) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MqttDump) SetID(id *int) {
	m.Id = id
}

// NewMqttPublish creates the object that can be sent to Home Assistant for domain mqtt, service publish
// "Publish a message to an MQTT topic."
func NewMqttPublish(entities []string, payload *string, qos *Qos, topic *string) *MqttPublish {
	serviceDomain := "mqtt"
	serviceType := "call_service"
	serviceService := "publish"
	m := &MqttPublish{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Payload *string `json:"payload,omitempty"`
			Qos     *Qos    `json:"qos,omitempty"`
			Topic   *string `json:"topic,omitempty"`
		}{
			Payload: payload,
			Qos:     qos,
			Topic:   topic,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return m
}

type MqttPublish struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Payload *string `json:"payload,omitempty"`
		Qos     *Qos    `json:"qos,omitempty"`
		Topic   *string `json:"topic,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MqttPublish) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MqttPublish) SetID(id *int) {
	m.Id = id
}

// NewMqttReload creates the object that can be sent to Home Assistant for domain mqtt, service reload
// "Reload all MQTT entities from YAML."
func NewMqttReload(entities []string) *MqttReload {
	serviceDomain := "mqtt"
	serviceType := "call_service"
	serviceService := "reload"
	m := &MqttReload{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return m
}

type MqttReload struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MqttReload) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MqttReload) SetID(id *int) {
	m.Id = id
}
