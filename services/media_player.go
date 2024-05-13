package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewMediaPlayerClearPlaylist creates the object that can be sent to Home Assistant for domain media_player, service clear_playlist
// "Send the media player the command to clear players playlist."
func NewMediaPlayerClearPlaylist(target Target) *MediaPlayerClearPlaylist {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "clear_playlist"
	m := &MediaPlayerClearPlaylist{
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
	return m
}

type MediaPlayerClearPlaylist struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerClearPlaylist) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerClearPlaylist) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerClearPlaylist) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerJoin creates the object that can be sent to Home Assistant for domain media_player, service join
// "Group players together. Only works on platforms with support for player groups."
func NewMediaPlayerJoin(target Target) *MediaPlayerJoin {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "join"
	m := &MediaPlayerJoin{
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
	return m
}

type MediaPlayerJoin struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerJoin) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerJoin) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerJoin) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerMediaNextTrack creates the object that can be sent to Home Assistant for domain media_player, service media_next_track
// "Send the media player the command for next track."
func NewMediaPlayerMediaNextTrack(target Target) *MediaPlayerMediaNextTrack {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_next_track"
	m := &MediaPlayerMediaNextTrack{
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
	return m
}

type MediaPlayerMediaNextTrack struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerMediaNextTrack) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaNextTrack) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerMediaNextTrack) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerMediaPause creates the object that can be sent to Home Assistant for domain media_player, service media_pause
// "Send the media player the command for pause."
func NewMediaPlayerMediaPause(target Target) *MediaPlayerMediaPause {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_pause"
	m := &MediaPlayerMediaPause{
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
	return m
}

type MediaPlayerMediaPause struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerMediaPause) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaPause) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerMediaPause) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerMediaPlay creates the object that can be sent to Home Assistant for domain media_player, service media_play
// "Send the media player the command for play."
func NewMediaPlayerMediaPlay(target Target) *MediaPlayerMediaPlay {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_play"
	m := &MediaPlayerMediaPlay{
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
	return m
}

type MediaPlayerMediaPlay struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerMediaPlay) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaPlay) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerMediaPlay) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerMediaPlayPause creates the object that can be sent to Home Assistant for domain media_player, service media_play_pause
// "Toggle media player play/pause state."
func NewMediaPlayerMediaPlayPause(target Target) *MediaPlayerMediaPlayPause {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_play_pause"
	m := &MediaPlayerMediaPlayPause{
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
	return m
}

type MediaPlayerMediaPlayPause struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerMediaPlayPause) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaPlayPause) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerMediaPlayPause) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerMediaPreviousTrack creates the object that can be sent to Home Assistant for domain media_player, service media_previous_track
// "Send the media player the command for previous track."
func NewMediaPlayerMediaPreviousTrack(target Target) *MediaPlayerMediaPreviousTrack {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_previous_track"
	m := &MediaPlayerMediaPreviousTrack{
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
	return m
}

type MediaPlayerMediaPreviousTrack struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerMediaPreviousTrack) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaPreviousTrack) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerMediaPreviousTrack) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerMediaSeek creates the object that can be sent to Home Assistant for domain media_player, service media_seek
// "Send the media player the command to seek in current playing media."
func NewMediaPlayerMediaSeek(target Target) *MediaPlayerMediaSeek {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_seek"
	m := &MediaPlayerMediaSeek{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: MediaPlayerMediaSeekParams{},
	}
	return m
}

type MediaPlayerMediaSeek struct {
	ServiceBase
	ServiceData MediaPlayerMediaSeekParams `json:"service_data,omitempty"`
}
type MediaPlayerMediaSeekParams struct {
	SeekPosition *float64 `json:"seek_position,omitempty"`
}

func (m *MediaPlayerMediaSeek) SeekPosition(seekPosition float64) *MediaPlayerMediaSeek {
	m.ServiceData.SeekPosition = &seekPosition
	return m
}
func (m *MediaPlayerMediaSeek) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaSeek) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerMediaSeek) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerMediaStop creates the object that can be sent to Home Assistant for domain media_player, service media_stop
// "Send the media player the stop command."
func NewMediaPlayerMediaStop(target Target) *MediaPlayerMediaStop {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_stop"
	m := &MediaPlayerMediaStop{
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
	return m
}

type MediaPlayerMediaStop struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerMediaStop) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaStop) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerMediaStop) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerPlayMedia creates the object that can be sent to Home Assistant for domain media_player, service play_media
// "Send the media player the command for playing media."
func NewMediaPlayerPlayMedia(target Target) *MediaPlayerPlayMedia {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "play_media"
	m := &MediaPlayerPlayMedia{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: MediaPlayerPlayMediaParams{},
	}
	return m
}

type MediaPlayerPlayMedia struct {
	ServiceBase
	ServiceData MediaPlayerPlayMediaParams `json:"service_data,omitempty"`
}
type MediaPlayerPlayMediaParams struct {
	Enqueue          *Enqueue `json:"enqueue,omitempty"`
	MediaContentId   *string  `json:"media_content_id,omitempty"`
	MediaContentType *string  `json:"media_content_type,omitempty"`
}

func (m *MediaPlayerPlayMedia) Enqueue(enqueue Enqueue) *MediaPlayerPlayMedia {
	m.ServiceData.Enqueue = &enqueue
	return m
}
func (m *MediaPlayerPlayMedia) MediaContentId(mediaContentId string) *MediaPlayerPlayMedia {
	m.ServiceData.MediaContentId = &mediaContentId
	return m
}
func (m *MediaPlayerPlayMedia) MediaContentType(mediaContentType string) *MediaPlayerPlayMedia {
	m.ServiceData.MediaContentType = &mediaContentType
	return m
}
func (m *MediaPlayerPlayMedia) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerPlayMedia) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerPlayMedia) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerRepeatSet creates the object that can be sent to Home Assistant for domain media_player, service repeat_set
// "Set repeat mode"
func NewMediaPlayerRepeatSet(target Target) *MediaPlayerRepeatSet {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "repeat_set"
	m := &MediaPlayerRepeatSet{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: MediaPlayerRepeatSetParams{},
	}
	return m
}

type MediaPlayerRepeatSet struct {
	ServiceBase
	ServiceData MediaPlayerRepeatSetParams `json:"service_data,omitempty"`
}
type MediaPlayerRepeatSetParams struct {
	Repeat *Repeat `json:"repeat,omitempty"`
}

func (m *MediaPlayerRepeatSet) Repeat(repeat Repeat) *MediaPlayerRepeatSet {
	m.ServiceData.Repeat = &repeat
	return m
}
func (m *MediaPlayerRepeatSet) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerRepeatSet) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerRepeatSet) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerSelectSoundMode creates the object that can be sent to Home Assistant for domain media_player, service select_sound_mode
// "Send the media player the command to change sound mode."
func NewMediaPlayerSelectSoundMode(target Target) *MediaPlayerSelectSoundMode {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "select_sound_mode"
	m := &MediaPlayerSelectSoundMode{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: MediaPlayerSelectSoundModeParams{},
	}
	return m
}

type MediaPlayerSelectSoundMode struct {
	ServiceBase
	ServiceData MediaPlayerSelectSoundModeParams `json:"service_data,omitempty"`
}
type MediaPlayerSelectSoundModeParams struct {
	SoundMode *string `json:"sound_mode,omitempty"`
}

func (m *MediaPlayerSelectSoundMode) SoundMode(soundMode string) *MediaPlayerSelectSoundMode {
	m.ServiceData.SoundMode = &soundMode
	return m
}
func (m *MediaPlayerSelectSoundMode) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerSelectSoundMode) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerSelectSoundMode) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerSelectSource creates the object that can be sent to Home Assistant for domain media_player, service select_source
// "Send the media player the command to change input source."
func NewMediaPlayerSelectSource(target Target) *MediaPlayerSelectSource {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "select_source"
	m := &MediaPlayerSelectSource{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: MediaPlayerSelectSourceParams{},
	}
	return m
}

type MediaPlayerSelectSource struct {
	ServiceBase
	ServiceData MediaPlayerSelectSourceParams `json:"service_data,omitempty"`
}
type MediaPlayerSelectSourceParams struct {
	Source *string `json:"source,omitempty"`
}

func (m *MediaPlayerSelectSource) Source(source string) *MediaPlayerSelectSource {
	m.ServiceData.Source = &source
	return m
}
func (m *MediaPlayerSelectSource) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerSelectSource) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerSelectSource) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerShuffleSet creates the object that can be sent to Home Assistant for domain media_player, service shuffle_set
// "Set shuffling state."
func NewMediaPlayerShuffleSet(target Target) *MediaPlayerShuffleSet {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "shuffle_set"
	m := &MediaPlayerShuffleSet{
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
	return m
}

type MediaPlayerShuffleSet struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerShuffleSet) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerShuffleSet) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerShuffleSet) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerToggle creates the object that can be sent to Home Assistant for domain media_player, service toggle
// "Toggles a media player power state."
func NewMediaPlayerToggle(target Target) *MediaPlayerToggle {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "toggle"
	m := &MediaPlayerToggle{
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
	return m
}

type MediaPlayerToggle struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerToggle) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerToggle) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerToggle) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerTurnOff creates the object that can be sent to Home Assistant for domain media_player, service turn_off
// "Turn a media player power off."
func NewMediaPlayerTurnOff(target Target) *MediaPlayerTurnOff {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "turn_off"
	m := &MediaPlayerTurnOff{
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
	return m
}

type MediaPlayerTurnOff struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerTurnOff) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerTurnOff) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerTurnOff) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerTurnOn creates the object that can be sent to Home Assistant for domain media_player, service turn_on
// "Turn a media player power on."
func NewMediaPlayerTurnOn(target Target) *MediaPlayerTurnOn {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "turn_on"
	m := &MediaPlayerTurnOn{
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
	return m
}

type MediaPlayerTurnOn struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerTurnOn) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerTurnOn) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerTurnOn) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerUnjoin creates the object that can be sent to Home Assistant for domain media_player, service unjoin
// "Unjoin the player from a group. Only works on platforms with support for player groups."
func NewMediaPlayerUnjoin(target Target) *MediaPlayerUnjoin {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "unjoin"
	m := &MediaPlayerUnjoin{
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
	return m
}

type MediaPlayerUnjoin struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerUnjoin) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerUnjoin) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerUnjoin) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerVolumeDown creates the object that can be sent to Home Assistant for domain media_player, service volume_down
// "Turn a media player volume down."
func NewMediaPlayerVolumeDown(target Target) *MediaPlayerVolumeDown {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "volume_down"
	m := &MediaPlayerVolumeDown{
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
	return m
}

type MediaPlayerVolumeDown struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerVolumeDown) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerVolumeDown) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerVolumeDown) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerVolumeMute creates the object that can be sent to Home Assistant for domain media_player, service volume_mute
// "Mute a media player's volume."
func NewMediaPlayerVolumeMute(target Target) *MediaPlayerVolumeMute {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "volume_mute"
	m := &MediaPlayerVolumeMute{
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
	return m
}

type MediaPlayerVolumeMute struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerVolumeMute) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerVolumeMute) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerVolumeMute) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerVolumeSet creates the object that can be sent to Home Assistant for domain media_player, service volume_set
// "Set a media player's volume level."
func NewMediaPlayerVolumeSet(target Target) *MediaPlayerVolumeSet {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "volume_set"
	m := &MediaPlayerVolumeSet{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: MediaPlayerVolumeSetParams{},
	}
	return m
}

type MediaPlayerVolumeSet struct {
	ServiceBase
	ServiceData MediaPlayerVolumeSetParams `json:"service_data,omitempty"`
}
type MediaPlayerVolumeSetParams struct {
	VolumeLevel *float64 `json:"volume_level,omitempty"`
}

func (m *MediaPlayerVolumeSet) VolumeLevel(volumeLevel float64) *MediaPlayerVolumeSet {
	m.ServiceData.VolumeLevel = &volumeLevel
	return m
}
func (m *MediaPlayerVolumeSet) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerVolumeSet) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerVolumeSet) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}

// NewMediaPlayerVolumeUp creates the object that can be sent to Home Assistant for domain media_player, service volume_up
// "Turn a media player volume up."
func NewMediaPlayerVolumeUp(target Target) *MediaPlayerVolumeUp {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "volume_up"
	m := &MediaPlayerVolumeUp{
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
	return m
}

type MediaPlayerVolumeUp struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (m *MediaPlayerVolumeUp) JSON() string {
	data, _ := gojson.Marshal(m)
	return string(data)
}
func (m *MediaPlayerVolumeUp) Targets() []string {
	return m.Target.EntityId
}
func (m *MediaPlayerVolumeUp) Name() string {
	return fmt.Sprintf("%s.%s", *m.Domain, *m.Service)
}
