package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCameraDisableMotionDetection creates the object that can be sent to Home Assistant for domain camera, service disable_motion_detection
// "Disable the motion detection in a camera."
func NewCameraDisableMotionDetection(target Target) *CameraDisableMotionDetection {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "disable_motion_detection"
	c := &CameraDisableMotionDetection{
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
	return c
}

type CameraDisableMotionDetection struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CameraDisableMotionDetection) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CameraDisableMotionDetection) Targets() []string {
	return c.Target.EntityId
}
func (c *CameraDisableMotionDetection) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCameraEnableMotionDetection creates the object that can be sent to Home Assistant for domain camera, service enable_motion_detection
// "Enable the motion detection in a camera."
func NewCameraEnableMotionDetection(target Target) *CameraEnableMotionDetection {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "enable_motion_detection"
	c := &CameraEnableMotionDetection{
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
	return c
}

type CameraEnableMotionDetection struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CameraEnableMotionDetection) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CameraEnableMotionDetection) Targets() []string {
	return c.Target.EntityId
}
func (c *CameraEnableMotionDetection) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCameraPlayStream creates the object that can be sent to Home Assistant for domain camera, service play_stream
// "Play camera stream on supported media player."
func NewCameraPlayStream(target Target) *CameraPlayStream {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "play_stream"
	c := &CameraPlayStream{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: CameraPlayStreamParams{},
	}
	return c
}

type CameraPlayStream struct {
	ServiceBase
	ServiceData CameraPlayStreamParams `json:"service_data,omitempty"`
}
type CameraPlayStreamParams struct {
	Format *Format `json:"format,omitempty"`
}

func (c *CameraPlayStream) Format(format Format) *CameraPlayStream {
	c.ServiceData.Format = &format
	return c
}
func (c *CameraPlayStream) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CameraPlayStream) Targets() []string {
	return c.Target.EntityId
}
func (c *CameraPlayStream) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCameraRecord creates the object that can be sent to Home Assistant for domain camera, service record
// "Record live camera feed."
func NewCameraRecord(target Target) *CameraRecord {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "record"
	c := &CameraRecord{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: CameraRecordParams{},
	}
	return c
}

type CameraRecord struct {
	ServiceBase
	ServiceData CameraRecordParams `json:"service_data,omitempty"`
}
type CameraRecordParams struct {
	Duration *float64 `json:"duration,omitempty"`
	Filename *string  `json:"filename,omitempty"`
	Lookback *float64 `json:"lookback,omitempty"`
}

func (c *CameraRecord) Duration(duration float64) *CameraRecord {
	c.ServiceData.Duration = &duration
	return c
}
func (c *CameraRecord) Filename(filename string) *CameraRecord {
	c.ServiceData.Filename = &filename
	return c
}
func (c *CameraRecord) Lookback(lookback float64) *CameraRecord {
	c.ServiceData.Lookback = &lookback
	return c
}
func (c *CameraRecord) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CameraRecord) Targets() []string {
	return c.Target.EntityId
}
func (c *CameraRecord) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCameraSnapshot creates the object that can be sent to Home Assistant for domain camera, service snapshot
// "Take a snapshot from a camera."
func NewCameraSnapshot(target Target) *CameraSnapshot {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "snapshot"
	c := &CameraSnapshot{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: CameraSnapshotParams{},
	}
	return c
}

type CameraSnapshot struct {
	ServiceBase
	ServiceData CameraSnapshotParams `json:"service_data,omitempty"`
}
type CameraSnapshotParams struct {
	Filename *string `json:"filename,omitempty"`
}

func (c *CameraSnapshot) Filename(filename string) *CameraSnapshot {
	c.ServiceData.Filename = &filename
	return c
}
func (c *CameraSnapshot) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CameraSnapshot) Targets() []string {
	return c.Target.EntityId
}
func (c *CameraSnapshot) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCameraTurnOff creates the object that can be sent to Home Assistant for domain camera, service turn_off
// "Turn off camera."
func NewCameraTurnOff(target Target) *CameraTurnOff {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "turn_off"
	c := &CameraTurnOff{
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
	return c
}

type CameraTurnOff struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CameraTurnOff) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CameraTurnOff) Targets() []string {
	return c.Target.EntityId
}
func (c *CameraTurnOff) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}

// NewCameraTurnOn creates the object that can be sent to Home Assistant for domain camera, service turn_on
// "Turn on camera."
func NewCameraTurnOn(target Target) *CameraTurnOn {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "turn_on"
	c := &CameraTurnOn{
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
	return c
}

type CameraTurnOn struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (c *CameraTurnOn) JSON() string {
	data, _ := gojson.Marshal(c)
	return string(data)
}
func (c *CameraTurnOn) Targets() []string {
	return c.Target.EntityId
}
func (c *CameraTurnOn) Name() string {
	return fmt.Sprintf("%s.%s", *c.Domain, *c.Service)
}
