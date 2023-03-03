package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewMediaPlayerClearPlaylist creates the object that can be sent to Home Assistant for domain media_player, service clear_playlist
// "Send the media player the command to clear players playlist."
func NewMediaPlayerClearPlaylist(entities []string) *MediaPlayerClearPlaylist {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "clear_playlist"
	m := &MediaPlayerClearPlaylist{
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

type MediaPlayerClearPlaylist struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerClearPlaylist) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerClearPlaylist) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerJoin creates the object that can be sent to Home Assistant for domain media_player, service join
// "Group players together. Only works on platforms with support for player groups."
func NewMediaPlayerJoin(entities []string) *MediaPlayerJoin {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "join"
	m := &MediaPlayerJoin{
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

type MediaPlayerJoin struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerJoin) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerJoin) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaNextTrack creates the object that can be sent to Home Assistant for domain media_player, service media_next_track
// "Send the media player the command for next track."
func NewMediaPlayerMediaNextTrack(entities []string) *MediaPlayerMediaNextTrack {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_next_track"
	m := &MediaPlayerMediaNextTrack{
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

type MediaPlayerMediaNextTrack struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerMediaNextTrack) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaNextTrack) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaPause creates the object that can be sent to Home Assistant for domain media_player, service media_pause
// "Send the media player the command for pause."
func NewMediaPlayerMediaPause(entities []string) *MediaPlayerMediaPause {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_pause"
	m := &MediaPlayerMediaPause{
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

type MediaPlayerMediaPause struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerMediaPause) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaPause) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaPlay creates the object that can be sent to Home Assistant for domain media_player, service media_play
// "Send the media player the command for play."
func NewMediaPlayerMediaPlay(entities []string) *MediaPlayerMediaPlay {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_play"
	m := &MediaPlayerMediaPlay{
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

type MediaPlayerMediaPlay struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerMediaPlay) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaPlay) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaPlayPause creates the object that can be sent to Home Assistant for domain media_player, service media_play_pause
// "Toggle media player play/pause state."
func NewMediaPlayerMediaPlayPause(entities []string) *MediaPlayerMediaPlayPause {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_play_pause"
	m := &MediaPlayerMediaPlayPause{
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

type MediaPlayerMediaPlayPause struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerMediaPlayPause) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaPlayPause) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaPreviousTrack creates the object that can be sent to Home Assistant for domain media_player, service media_previous_track
// "Send the media player the command for previous track."
func NewMediaPlayerMediaPreviousTrack(entities []string) *MediaPlayerMediaPreviousTrack {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_previous_track"
	m := &MediaPlayerMediaPreviousTrack{
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

type MediaPlayerMediaPreviousTrack struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerMediaPreviousTrack) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaPreviousTrack) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaSeek creates the object that can be sent to Home Assistant for domain media_player, service media_seek
// "Send the media player the command to seek in current playing media."
func NewMediaPlayerMediaSeek(entities []string, seekPosition *float64) *MediaPlayerMediaSeek {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_seek"
	m := &MediaPlayerMediaSeek{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			SeekPosition *float64 `json:"seek_position,omitempty"`
		}{SeekPosition: seekPosition},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return m
}

type MediaPlayerMediaSeek struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		SeekPosition *float64 `json:"seek_position,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerMediaSeek) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaSeek) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaStop creates the object that can be sent to Home Assistant for domain media_player, service media_stop
// "Send the media player the stop command."
func NewMediaPlayerMediaStop(entities []string) *MediaPlayerMediaStop {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_stop"
	m := &MediaPlayerMediaStop{
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

type MediaPlayerMediaStop struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerMediaStop) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaStop) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerPlayMedia creates the object that can be sent to Home Assistant for domain media_player, service play_media
// "Send the media player the command for playing media."
func NewMediaPlayerPlayMedia(entities []string, enqueue *Enqueue, mediaContentId *string, mediaContentType *string) *MediaPlayerPlayMedia {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "play_media"
	m := &MediaPlayerPlayMedia{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Enqueue          *Enqueue `json:"enqueue,omitempty"`
			MediaContentId   *string  `json:"media_content_id,omitempty"`
			MediaContentType *string  `json:"media_content_type,omitempty"`
		}{
			Enqueue:          enqueue,
			MediaContentId:   mediaContentId,
			MediaContentType: mediaContentType,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return m
}

type MediaPlayerPlayMedia struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Enqueue          *Enqueue `json:"enqueue,omitempty"`
		MediaContentId   *string  `json:"media_content_id,omitempty"`
		MediaContentType *string  `json:"media_content_type,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerPlayMedia) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerPlayMedia) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerRepeatSet creates the object that can be sent to Home Assistant for domain media_player, service repeat_set
// "Set repeat mode"
func NewMediaPlayerRepeatSet(entities []string, repeat *Repeat) *MediaPlayerRepeatSet {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "repeat_set"
	m := &MediaPlayerRepeatSet{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Repeat *Repeat `json:"repeat,omitempty"`
		}{Repeat: repeat},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return m
}

type MediaPlayerRepeatSet struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Repeat *Repeat `json:"repeat,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerRepeatSet) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerRepeatSet) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerSelectSoundMode creates the object that can be sent to Home Assistant for domain media_player, service select_sound_mode
// "Send the media player the command to change sound mode."
func NewMediaPlayerSelectSoundMode(entities []string, soundMode *string) *MediaPlayerSelectSoundMode {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "select_sound_mode"
	m := &MediaPlayerSelectSoundMode{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			SoundMode *string `json:"sound_mode,omitempty"`
		}{SoundMode: soundMode},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return m
}

type MediaPlayerSelectSoundMode struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		SoundMode *string `json:"sound_mode,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerSelectSoundMode) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerSelectSoundMode) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerSelectSource creates the object that can be sent to Home Assistant for domain media_player, service select_source
// "Send the media player the command to change input source."
func NewMediaPlayerSelectSource(entities []string, source *string) *MediaPlayerSelectSource {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "select_source"
	m := &MediaPlayerSelectSource{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Source *string `json:"source,omitempty"`
		}{Source: source},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return m
}

type MediaPlayerSelectSource struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Source *string `json:"source,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerSelectSource) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerSelectSource) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerShuffleSet creates the object that can be sent to Home Assistant for domain media_player, service shuffle_set
// "Set shuffling state."
func NewMediaPlayerShuffleSet(entities []string) *MediaPlayerShuffleSet {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "shuffle_set"
	m := &MediaPlayerShuffleSet{
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

type MediaPlayerShuffleSet struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerShuffleSet) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerShuffleSet) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerToggle creates the object that can be sent to Home Assistant for domain media_player, service toggle
// "Toggles a media player power state."
func NewMediaPlayerToggle(entities []string) *MediaPlayerToggle {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "toggle"
	m := &MediaPlayerToggle{
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

type MediaPlayerToggle struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerToggle) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerToggle) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerTurnOff creates the object that can be sent to Home Assistant for domain media_player, service turn_off
// "Turn a media player power off."
func NewMediaPlayerTurnOff(entities []string) *MediaPlayerTurnOff {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "turn_off"
	m := &MediaPlayerTurnOff{
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

type MediaPlayerTurnOff struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerTurnOff) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerTurnOff) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerTurnOn creates the object that can be sent to Home Assistant for domain media_player, service turn_on
// "Turn a media player power on."
func NewMediaPlayerTurnOn(entities []string) *MediaPlayerTurnOn {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "turn_on"
	m := &MediaPlayerTurnOn{
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

type MediaPlayerTurnOn struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerTurnOn) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerTurnOn) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerUnjoin creates the object that can be sent to Home Assistant for domain media_player, service unjoin
// "Unjoin the player from a group. Only works on platforms with support for player groups."
func NewMediaPlayerUnjoin(entities []string) *MediaPlayerUnjoin {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "unjoin"
	m := &MediaPlayerUnjoin{
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

type MediaPlayerUnjoin struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerUnjoin) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerUnjoin) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerVolumeDown creates the object that can be sent to Home Assistant for domain media_player, service volume_down
// "Turn a media player volume down."
func NewMediaPlayerVolumeDown(entities []string) *MediaPlayerVolumeDown {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "volume_down"
	m := &MediaPlayerVolumeDown{
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

type MediaPlayerVolumeDown struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerVolumeDown) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerVolumeDown) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerVolumeMute creates the object that can be sent to Home Assistant for domain media_player, service volume_mute
// "Mute a media player's volume."
func NewMediaPlayerVolumeMute(entities []string) *MediaPlayerVolumeMute {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "volume_mute"
	m := &MediaPlayerVolumeMute{
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

type MediaPlayerVolumeMute struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerVolumeMute) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerVolumeMute) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerVolumeSet creates the object that can be sent to Home Assistant for domain media_player, service volume_set
// "Set a media player's volume level."
func NewMediaPlayerVolumeSet(entities []string, volumeLevel *float64) *MediaPlayerVolumeSet {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "volume_set"
	m := &MediaPlayerVolumeSet{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			VolumeLevel *float64 `json:"volume_level,omitempty"`
		}{VolumeLevel: volumeLevel},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return m
}

type MediaPlayerVolumeSet struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		VolumeLevel *float64 `json:"volume_level,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerVolumeSet) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerVolumeSet) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerVolumeUp creates the object that can be sent to Home Assistant for domain media_player, service volume_up
// "Turn a media player volume up."
func NewMediaPlayerVolumeUp(entities []string) *MediaPlayerVolumeUp {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "volume_up"
	m := &MediaPlayerVolumeUp{
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

type MediaPlayerVolumeUp struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (m *MediaPlayerVolumeUp) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerVolumeUp) SetID(id *int) {
	m.Id = id
}
