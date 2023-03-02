package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewCameraDisableMotionDetection creates the object that can be sent to Home Assistant for domain camera, service disable_motion_detection
// "Disable the motion detection in a camera."
func NewCameraDisableMotionDetection(entities []string) *CameraDisableMotionDetection {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "disable_motion_detection"
	c := &CameraDisableMotionDetection{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CameraDisableMotionDetection struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CameraDisableMotionDetection) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CameraDisableMotionDetection) SetID(id *int) {
	c.Id = id
}

// NewCameraEnableMotionDetection creates the object that can be sent to Home Assistant for domain camera, service enable_motion_detection
// "Enable the motion detection in a camera."
func NewCameraEnableMotionDetection(entities []string) *CameraEnableMotionDetection {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "enable_motion_detection"
	c := &CameraEnableMotionDetection{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CameraEnableMotionDetection struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CameraEnableMotionDetection) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CameraEnableMotionDetection) SetID(id *int) {
	c.Id = id
}

// NewCameraPlayStream creates the object that can be sent to Home Assistant for domain camera, service play_stream
// "Play camera stream on supported media player."
func NewCameraPlayStream(entities []string, format *Format) *CameraPlayStream {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "play_stream"
	c := &CameraPlayStream{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Format *Format `json:"format,omitempty"`
		}{Format: format},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CameraPlayStream struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Format *Format `json:"format,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CameraPlayStream) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CameraPlayStream) SetID(id *int) {
	c.Id = id
}

// NewCameraRecord creates the object that can be sent to Home Assistant for domain camera, service record
// "Record live camera feed."
func NewCameraRecord(entities []string, duration *int, filename *string, lookback *int) *CameraRecord {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "record"
	c := &CameraRecord{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Duration *int    `json:"duration,omitempty"`
			Filename *string `json:"filename,omitempty"`
			Lookback *int    `json:"lookback,omitempty"`
		}{
			Duration: duration,
			Filename: filename,
			Lookback: lookback,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CameraRecord struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Duration *int    `json:"duration,omitempty"`
		Filename *string `json:"filename,omitempty"`
		Lookback *int    `json:"lookback,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CameraRecord) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CameraRecord) SetID(id *int) {
	c.Id = id
}

// NewCameraSnapshot creates the object that can be sent to Home Assistant for domain camera, service snapshot
// "Take a snapshot from a camera."
func NewCameraSnapshot(entities []string, filename *string) *CameraSnapshot {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "snapshot"
	c := &CameraSnapshot{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Filename *string `json:"filename,omitempty"`
		}{Filename: filename},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CameraSnapshot struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Filename *string `json:"filename,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CameraSnapshot) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CameraSnapshot) SetID(id *int) {
	c.Id = id
}

// NewCameraTurnOff creates the object that can be sent to Home Assistant for domain camera, service turn_off
// "Turn off camera."
func NewCameraTurnOff(entities []string) *CameraTurnOff {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "turn_off"
	c := &CameraTurnOff{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CameraTurnOff struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CameraTurnOff) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CameraTurnOff) SetID(id *int) {
	c.Id = id
}

// NewCameraTurnOn creates the object that can be sent to Home Assistant for domain camera, service turn_on
// "Turn on camera."
func NewCameraTurnOn(entities []string) *CameraTurnOn {
	serviceDomain := "camera"
	serviceType := "call_service"
	serviceService := "turn_on"
	c := &CameraTurnOn{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return c
}

type CameraTurnOn struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (c *CameraTurnOn) JSON() string {
	data, _ := json.Marshal(c)
	return string(data)
}
func (c *CameraTurnOn) SetID(id *int) {
	c.Id = id
}
