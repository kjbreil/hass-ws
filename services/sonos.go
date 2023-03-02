package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewSonosClearSleepTimer creates the object that can be sent to Home Assistant for domain sonos, service clear_sleep_timer
// "Clear a Sonos timer."
func NewSonosClearSleepTimer(entities []string) *SonosClearSleepTimer {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "clear_sleep_timer"
	s := &SonosClearSleepTimer{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SonosClearSleepTimer struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SonosClearSleepTimer) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SonosClearSleepTimer) SetID(id *int) {
	s.Id = id
}

// NewSonosPlayQueue creates the object that can be sent to Home Assistant for domain sonos, service play_queue
// "Start playing the queue from the first item."
func NewSonosPlayQueue(entities []string, queuePosition *int) *SonosPlayQueue {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "play_queue"
	s := &SonosPlayQueue{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			QueuePosition *int `json:"queue_position,omitempty"`
		}{QueuePosition: queuePosition},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SonosPlayQueue struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		QueuePosition *int `json:"queue_position,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SonosPlayQueue) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SonosPlayQueue) SetID(id *int) {
	s.Id = id
}

// NewSonosRemoveFromQueue creates the object that can be sent to Home Assistant for domain sonos, service remove_from_queue
// "Removes an item from the queue."
func NewSonosRemoveFromQueue(entities []string, queuePosition *int) *SonosRemoveFromQueue {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "remove_from_queue"
	s := &SonosRemoveFromQueue{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			QueuePosition *int `json:"queue_position,omitempty"`
		}{QueuePosition: queuePosition},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SonosRemoveFromQueue struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		QueuePosition *int `json:"queue_position,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SonosRemoveFromQueue) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SonosRemoveFromQueue) SetID(id *int) {
	s.Id = id
}

// NewSonosRestore creates the object that can be sent to Home Assistant for domain sonos, service restore
// "Restore a snapshot of the media player."
func NewSonosRestore(entities []string) *SonosRestore {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "restore"
	s := &SonosRestore{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SonosRestore struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SonosRestore) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SonosRestore) SetID(id *int) {
	s.Id = id
}

// NewSonosSetSleepTimer creates the object that can be sent to Home Assistant for domain sonos, service set_sleep_timer
// "Set a Sonos timer."
func NewSonosSetSleepTimer(entities []string, sleepTime *int) *SonosSetSleepTimer {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "set_sleep_timer"
	s := &SonosSetSleepTimer{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			SleepTime *int `json:"sleep_time,omitempty"`
		}{SleepTime: sleepTime},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SonosSetSleepTimer struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		SleepTime *int `json:"sleep_time,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SonosSetSleepTimer) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SonosSetSleepTimer) SetID(id *int) {
	s.Id = id
}

// NewSonosSnapshot creates the object that can be sent to Home Assistant for domain sonos, service snapshot
// "Take a snapshot of the media player."
func NewSonosSnapshot(entities []string) *SonosSnapshot {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "snapshot"
	s := &SonosSnapshot{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SonosSnapshot struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SonosSnapshot) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SonosSnapshot) SetID(id *int) {
	s.Id = id
}

// NewSonosUpdateAlarm creates the object that can be sent to Home Assistant for domain sonos, service update_alarm
// "Updates an alarm with new time and volume settings."
func NewSonosUpdateAlarm(entities []string, alarmId *int, volume *int) *SonosUpdateAlarm {
	serviceDomain := "sonos"
	serviceType := "call_service"
	serviceService := "update_alarm"
	s := &SonosUpdateAlarm{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			AlarmId *int `json:"alarm_id,omitempty"`
			Volume  *int `json:"volume,omitempty"`
		}{
			AlarmId: alarmId,
			Volume:  volume,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return s
}

type SonosUpdateAlarm struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		AlarmId *int `json:"alarm_id,omitempty"`
		Volume  *int `json:"volume,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (s *SonosUpdateAlarm) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
func (s *SonosUpdateAlarm) SetID(id *int) {
	s.Id = id
}
