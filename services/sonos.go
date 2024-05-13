package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSonosClearSleepTimer creates the object that can be sent to Home Assistant for domain sonos, service clear_sleep_timer
// "Clear a Sonos timer."
func NewSonosClearSleepTimer(target Target) *SonosClearSleepTimer {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "clear_sleep_timer"
	s := &SonosClearSleepTimer{
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
	return s
}

type SonosClearSleepTimer struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *SonosClearSleepTimer) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SonosClearSleepTimer) Targets() []string {
	return s.Target.EntityId
}
func (s *SonosClearSleepTimer) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}

// NewSonosPlayQueue creates the object that can be sent to Home Assistant for domain sonos, service play_queue
// "Start playing the queue from the first item."
func NewSonosPlayQueue(target Target) *SonosPlayQueue {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "play_queue"
	s := &SonosPlayQueue{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: SonosPlayQueueParams{},
	}
	return s
}

type SonosPlayQueue struct {
	ServiceBase
	ServiceData SonosPlayQueueParams `json:"service_data,omitempty"`
}
type SonosPlayQueueParams struct {
	QueuePosition *float64 `json:"queue_position,omitempty"`
}

func (s *SonosPlayQueue) QueuePosition(queuePosition float64) *SonosPlayQueue {
	s.ServiceData.QueuePosition = &queuePosition
	return s
}
func (s *SonosPlayQueue) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SonosPlayQueue) Targets() []string {
	return s.Target.EntityId
}
func (s *SonosPlayQueue) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}

// NewSonosRemoveFromQueue creates the object that can be sent to Home Assistant for domain sonos, service remove_from_queue
// "Removes an item from the queue."
func NewSonosRemoveFromQueue(target Target) *SonosRemoveFromQueue {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "remove_from_queue"
	s := &SonosRemoveFromQueue{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: SonosRemoveFromQueueParams{},
	}
	return s
}

type SonosRemoveFromQueue struct {
	ServiceBase
	ServiceData SonosRemoveFromQueueParams `json:"service_data,omitempty"`
}
type SonosRemoveFromQueueParams struct {
	QueuePosition *float64 `json:"queue_position,omitempty"`
}

func (s *SonosRemoveFromQueue) QueuePosition(queuePosition float64) *SonosRemoveFromQueue {
	s.ServiceData.QueuePosition = &queuePosition
	return s
}
func (s *SonosRemoveFromQueue) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SonosRemoveFromQueue) Targets() []string {
	return s.Target.EntityId
}
func (s *SonosRemoveFromQueue) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}

// NewSonosRestore creates the object that can be sent to Home Assistant for domain sonos, service restore
// "Restore a snapshot of the media player."
func NewSonosRestore(target Target) *SonosRestore {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "restore"
	s := &SonosRestore{
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
	return s
}

type SonosRestore struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *SonosRestore) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SonosRestore) Targets() []string {
	return s.Target.EntityId
}
func (s *SonosRestore) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}

// NewSonosSetSleepTimer creates the object that can be sent to Home Assistant for domain sonos, service set_sleep_timer
// "Set a Sonos timer."
func NewSonosSetSleepTimer(target Target) *SonosSetSleepTimer {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "set_sleep_timer"
	s := &SonosSetSleepTimer{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: SonosSetSleepTimerParams{},
	}
	return s
}

type SonosSetSleepTimer struct {
	ServiceBase
	ServiceData SonosSetSleepTimerParams `json:"service_data,omitempty"`
}
type SonosSetSleepTimerParams struct {
	SleepTime *float64 `json:"sleep_time,omitempty"`
}

func (s *SonosSetSleepTimer) SleepTime(sleepTime float64) *SonosSetSleepTimer {
	s.ServiceData.SleepTime = &sleepTime
	return s
}
func (s *SonosSetSleepTimer) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SonosSetSleepTimer) Targets() []string {
	return s.Target.EntityId
}
func (s *SonosSetSleepTimer) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}

// NewSonosSnapshot creates the object that can be sent to Home Assistant for domain sonos, service snapshot
// "Take a snapshot of the media player."
func NewSonosSnapshot(target Target) *SonosSnapshot {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "snapshot"
	s := &SonosSnapshot{
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
	return s
}

type SonosSnapshot struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (s *SonosSnapshot) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SonosSnapshot) Targets() []string {
	return s.Target.EntityId
}
func (s *SonosSnapshot) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}

// NewSonosUpdateAlarm creates the object that can be sent to Home Assistant for domain sonos, service update_alarm
// "Updates an alarm with new time and volume settings."
func NewSonosUpdateAlarm(target Target) *SonosUpdateAlarm {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "update_alarm"
	s := &SonosUpdateAlarm{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: SonosUpdateAlarmParams{},
	}
	return s
}

type SonosUpdateAlarm struct {
	ServiceBase
	ServiceData SonosUpdateAlarmParams `json:"service_data,omitempty"`
}
type SonosUpdateAlarmParams struct {
	AlarmId *float64 `json:"alarm_id,omitempty"`
	Volume  *float64 `json:"volume,omitempty"`
}

func (s *SonosUpdateAlarm) AlarmId(alarmId float64) *SonosUpdateAlarm {
	s.ServiceData.AlarmId = &alarmId
	return s
}
func (s *SonosUpdateAlarm) Volume(volume float64) *SonosUpdateAlarm {
	s.ServiceData.Volume = &volume
	return s
}
func (s *SonosUpdateAlarm) JSON() string {
	data, _ := gojson.Marshal(s)
	return string(data)
}
func (s *SonosUpdateAlarm) Targets() []string {
	return s.Target.EntityId
}
func (s *SonosUpdateAlarm) Name() string {
	return fmt.Sprintf("%s.%s", *s.Domain, *s.Service)
}
