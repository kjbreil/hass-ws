package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewMqttDump creates the object that can be sent to Home Assistant for domain mqtt, service dump
// "Dump messages on a topic selector to the 'mqtt_dump.txt' file in your configuration folder."
func NewMqttDump(target Target) *MqttDump {
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
		ServiceData: MqttDumpParams{},
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

func (m *MqttDump) Duration(duration float64) *MqttDump {
	m.ServiceData.Duration = &duration
	return m
}
func (m *MqttDump) Topic(topic string) *MqttDump {
	m.ServiceData.Topic = &topic
	return m
}
func (m *MqttDump) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MqttDump) Targets() []string {
	return m.Target.EntityId
}
func (m *MqttDump) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}
func (m *MqttDump) SetID(id *int) {
	m.Id = id
}

// NewMqttPublish creates the object that can be sent to Home Assistant for domain mqtt, service publish
// "Publish a message to an MQTT topic."
func NewMqttPublish(target Target) *MqttPublish {
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
		ServiceData: MqttPublishParams{},
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

func (m *MqttPublish) Payload(payload string) *MqttPublish {
	m.ServiceData.Payload = &payload
	return m
}
func (m *MqttPublish) Qos(qos Qos) *MqttPublish {
	m.ServiceData.Qos = &qos
	return m
}
func (m *MqttPublish) Topic(topic string) *MqttPublish {
	m.ServiceData.Topic = &topic
	return m
}
func (m *MqttPublish) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MqttPublish) Targets() []string {
	return m.Target.EntityId
}
func (m *MqttPublish) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}
func (m *MqttPublish) SetID(id *int) {
	m.Id = id
}

// NewMqttReload creates the object that can be sent to Home Assistant for domain mqtt, service reload
// "Reload all MQTT entities from YAML."
func NewMqttReload(target Target) *MqttReload {
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
		ServiceData: nil,
	}
	return m
}

type MqttReload struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MqttReload) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MqttReload) Targets() []string {
	return m.Target.EntityId
}
func (m *MqttReload) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}
func (m *MqttReload) SetID(id *int) {
	m.Id = id
}
