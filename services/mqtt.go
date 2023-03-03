package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewMqttDump creates the object that can be sent to Home Assistant for domain mqtt, service dump
// "Dump messages on a topic selector to the 'mqtt_dump.txt' file in your configuration folder."
func NewMqttDump(target Target, mqttDumpParams MqttDumpParams) *MqttDump {
	serviceDomain := "mqtt"
	serviceType := "call_service"
	serviceService := "dump"
	m := &MqttDump{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mqttDumpParams,
	}
	return m
}

type MqttDump struct {
	ServiceBase
	ServiceData MqttDumpParams `json:"service_data,omitempty"`
}
type MqttDumpParams struct {
	Duration *float64 `json:"duration,omitempty"`
	Topic    *string  `json:"topic,omitempty"`
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
func NewMqttPublish(target Target, mqttPublishParams MqttPublishParams) *MqttPublish {
	serviceDomain := "mqtt"
	serviceType := "call_service"
	serviceService := "publish"
	m := &MqttPublish{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mqttPublishParams,
	}
	return m
}

type MqttPublish struct {
	ServiceBase
	ServiceData MqttPublishParams `json:"service_data,omitempty"`
}
type MqttPublishParams struct {
	Payload *string `json:"payload,omitempty"`
	Qos     *Qos    `json:"qos,omitempty"`
	Topic   *string `json:"topic,omitempty"`
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
func NewMqttReload(target Target, mqttReloadParams MqttReloadParams) *MqttReload {
	serviceDomain := "mqtt"
	serviceType := "call_service"
	serviceService := "reload"
	m := &MqttReload{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mqttReloadParams,
	}
	return m
}

type MqttReload struct {
	ServiceBase
	ServiceData MqttReloadParams `json:"service_data,omitempty"`
}
type MqttReloadParams struct{}

func (m *MqttReload) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MqttReload) SetID(id *int) {
	m.Id = id
}
