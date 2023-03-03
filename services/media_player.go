package services

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewMediaPlayerClearPlaylist creates the object that can be sent to Home Assistant for domain media_player, service clear_playlist
// "Send the media player the command to clear players playlist."
func NewMediaPlayerClearPlaylist(target Target, mediaPlayerClearPlaylistParams MediaPlayerClearPlaylistParams) *MediaPlayerClearPlaylist {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "clear_playlist"
	m := &MediaPlayerClearPlaylist{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerClearPlaylistParams,
	}
	return m
}

type MediaPlayerClearPlaylist struct {
	ServiceBase
	ServiceData MediaPlayerClearPlaylistParams `json:"service_data,omitempty"`
}
type MediaPlayerClearPlaylistParams struct{}

func (m *MediaPlayerClearPlaylist) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerClearPlaylist) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerJoin creates the object that can be sent to Home Assistant for domain media_player, service join
// "Group players together. Only works on platforms with support for player groups."
func NewMediaPlayerJoin(target Target, mediaPlayerJoinParams MediaPlayerJoinParams) *MediaPlayerJoin {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "join"
	m := &MediaPlayerJoin{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerJoinParams,
	}
	return m
}

type MediaPlayerJoin struct {
	ServiceBase
	ServiceData MediaPlayerJoinParams `json:"service_data,omitempty"`
}
type MediaPlayerJoinParams struct{}

func (m *MediaPlayerJoin) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerJoin) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaNextTrack creates the object that can be sent to Home Assistant for domain media_player, service media_next_track
// "Send the media player the command for next track."
func NewMediaPlayerMediaNextTrack(target Target, mediaPlayerMediaNextTrackParams MediaPlayerMediaNextTrackParams) *MediaPlayerMediaNextTrack {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_next_track"
	m := &MediaPlayerMediaNextTrack{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerMediaNextTrackParams,
	}
	return m
}

type MediaPlayerMediaNextTrack struct {
	ServiceBase
	ServiceData MediaPlayerMediaNextTrackParams `json:"service_data,omitempty"`
}
type MediaPlayerMediaNextTrackParams struct{}

func (m *MediaPlayerMediaNextTrack) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaNextTrack) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaPause creates the object that can be sent to Home Assistant for domain media_player, service media_pause
// "Send the media player the command for pause."
func NewMediaPlayerMediaPause(target Target, mediaPlayerMediaPauseParams MediaPlayerMediaPauseParams) *MediaPlayerMediaPause {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_pause"
	m := &MediaPlayerMediaPause{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerMediaPauseParams,
	}
	return m
}

type MediaPlayerMediaPause struct {
	ServiceBase
	ServiceData MediaPlayerMediaPauseParams `json:"service_data,omitempty"`
}
type MediaPlayerMediaPauseParams struct{}

func (m *MediaPlayerMediaPause) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaPause) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaPlay creates the object that can be sent to Home Assistant for domain media_player, service media_play
// "Send the media player the command for play."
func NewMediaPlayerMediaPlay(target Target, mediaPlayerMediaPlayParams MediaPlayerMediaPlayParams) *MediaPlayerMediaPlay {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_play"
	m := &MediaPlayerMediaPlay{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerMediaPlayParams,
	}
	return m
}

type MediaPlayerMediaPlay struct {
	ServiceBase
	ServiceData MediaPlayerMediaPlayParams `json:"service_data,omitempty"`
}
type MediaPlayerMediaPlayParams struct{}

func (m *MediaPlayerMediaPlay) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaPlay) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaPlayPause creates the object that can be sent to Home Assistant for domain media_player, service media_play_pause
// "Toggle media player play/pause state."
func NewMediaPlayerMediaPlayPause(target Target, mediaPlayerMediaPlayPauseParams MediaPlayerMediaPlayPauseParams) *MediaPlayerMediaPlayPause {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_play_pause"
	m := &MediaPlayerMediaPlayPause{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerMediaPlayPauseParams,
	}
	return m
}

type MediaPlayerMediaPlayPause struct {
	ServiceBase
	ServiceData MediaPlayerMediaPlayPauseParams `json:"service_data,omitempty"`
}
type MediaPlayerMediaPlayPauseParams struct{}

func (m *MediaPlayerMediaPlayPause) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaPlayPause) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaPreviousTrack creates the object that can be sent to Home Assistant for domain media_player, service media_previous_track
// "Send the media player the command for previous track."
func NewMediaPlayerMediaPreviousTrack(target Target, mediaPlayerMediaPreviousTrackParams MediaPlayerMediaPreviousTrackParams) *MediaPlayerMediaPreviousTrack {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_previous_track"
	m := &MediaPlayerMediaPreviousTrack{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerMediaPreviousTrackParams,
	}
	return m
}

type MediaPlayerMediaPreviousTrack struct {
	ServiceBase
	ServiceData MediaPlayerMediaPreviousTrackParams `json:"service_data,omitempty"`
}
type MediaPlayerMediaPreviousTrackParams struct{}

func (m *MediaPlayerMediaPreviousTrack) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaPreviousTrack) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaSeek creates the object that can be sent to Home Assistant for domain media_player, service media_seek
// "Send the media player the command to seek in current playing media."
func NewMediaPlayerMediaSeek(target Target, mediaPlayerMediaSeekParams MediaPlayerMediaSeekParams) *MediaPlayerMediaSeek {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_seek"
	m := &MediaPlayerMediaSeek{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerMediaSeekParams,
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

func (m *MediaPlayerMediaSeek) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaSeek) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerMediaStop creates the object that can be sent to Home Assistant for domain media_player, service media_stop
// "Send the media player the stop command."
func NewMediaPlayerMediaStop(target Target, mediaPlayerMediaStopParams MediaPlayerMediaStopParams) *MediaPlayerMediaStop {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "media_stop"
	m := &MediaPlayerMediaStop{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerMediaStopParams,
	}
	return m
}

type MediaPlayerMediaStop struct {
	ServiceBase
	ServiceData MediaPlayerMediaStopParams `json:"service_data,omitempty"`
}
type MediaPlayerMediaStopParams struct{}

func (m *MediaPlayerMediaStop) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerMediaStop) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerPlayMedia creates the object that can be sent to Home Assistant for domain media_player, service play_media
// "Send the media player the command for playing media."
func NewMediaPlayerPlayMedia(target Target, mediaPlayerPlayMediaParams MediaPlayerPlayMediaParams) *MediaPlayerPlayMedia {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "play_media"
	m := &MediaPlayerPlayMedia{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerPlayMediaParams,
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

func (m *MediaPlayerPlayMedia) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerPlayMedia) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerRepeatSet creates the object that can be sent to Home Assistant for domain media_player, service repeat_set
// "Set repeat mode"
func NewMediaPlayerRepeatSet(target Target, mediaPlayerRepeatSetParams MediaPlayerRepeatSetParams) *MediaPlayerRepeatSet {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "repeat_set"
	m := &MediaPlayerRepeatSet{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerRepeatSetParams,
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

func (m *MediaPlayerRepeatSet) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerRepeatSet) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerSelectSoundMode creates the object that can be sent to Home Assistant for domain media_player, service select_sound_mode
// "Send the media player the command to change sound mode."
func NewMediaPlayerSelectSoundMode(target Target, mediaPlayerSelectSoundModeParams MediaPlayerSelectSoundModeParams) *MediaPlayerSelectSoundMode {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "select_sound_mode"
	m := &MediaPlayerSelectSoundMode{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerSelectSoundModeParams,
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

func (m *MediaPlayerSelectSoundMode) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerSelectSoundMode) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerSelectSource creates the object that can be sent to Home Assistant for domain media_player, service select_source
// "Send the media player the command to change input source."
func NewMediaPlayerSelectSource(target Target, mediaPlayerSelectSourceParams MediaPlayerSelectSourceParams) *MediaPlayerSelectSource {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "select_source"
	m := &MediaPlayerSelectSource{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerSelectSourceParams,
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

func (m *MediaPlayerSelectSource) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerSelectSource) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerShuffleSet creates the object that can be sent to Home Assistant for domain media_player, service shuffle_set
// "Set shuffling state."
func NewMediaPlayerShuffleSet(target Target, mediaPlayerShuffleSetParams MediaPlayerShuffleSetParams) *MediaPlayerShuffleSet {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "shuffle_set"
	m := &MediaPlayerShuffleSet{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerShuffleSetParams,
	}
	return m
}

type MediaPlayerShuffleSet struct {
	ServiceBase
	ServiceData MediaPlayerShuffleSetParams `json:"service_data,omitempty"`
}
type MediaPlayerShuffleSetParams struct{}

func (m *MediaPlayerShuffleSet) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerShuffleSet) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerToggle creates the object that can be sent to Home Assistant for domain media_player, service toggle
// "Toggles a media player power state."
func NewMediaPlayerToggle(target Target, mediaPlayerToggleParams MediaPlayerToggleParams) *MediaPlayerToggle {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "toggle"
	m := &MediaPlayerToggle{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerToggleParams,
	}
	return m
}

type MediaPlayerToggle struct {
	ServiceBase
	ServiceData MediaPlayerToggleParams `json:"service_data,omitempty"`
}
type MediaPlayerToggleParams struct{}

func (m *MediaPlayerToggle) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerToggle) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerTurnOff creates the object that can be sent to Home Assistant for domain media_player, service turn_off
// "Turn a media player power off."
func NewMediaPlayerTurnOff(target Target, mediaPlayerTurnOffParams MediaPlayerTurnOffParams) *MediaPlayerTurnOff {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "turn_off"
	m := &MediaPlayerTurnOff{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerTurnOffParams,
	}
	return m
}

type MediaPlayerTurnOff struct {
	ServiceBase
	ServiceData MediaPlayerTurnOffParams `json:"service_data,omitempty"`
}
type MediaPlayerTurnOffParams struct{}

func (m *MediaPlayerTurnOff) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerTurnOff) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerTurnOn creates the object that can be sent to Home Assistant for domain media_player, service turn_on
// "Turn a media player power on."
func NewMediaPlayerTurnOn(target Target, mediaPlayerTurnOnParams MediaPlayerTurnOnParams) *MediaPlayerTurnOn {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "turn_on"
	m := &MediaPlayerTurnOn{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerTurnOnParams,
	}
	return m
}

type MediaPlayerTurnOn struct {
	ServiceBase
	ServiceData MediaPlayerTurnOnParams `json:"service_data,omitempty"`
}
type MediaPlayerTurnOnParams struct{}

func (m *MediaPlayerTurnOn) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerTurnOn) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerUnjoin creates the object that can be sent to Home Assistant for domain media_player, service unjoin
// "Unjoin the player from a group. Only works on platforms with support for player groups."
func NewMediaPlayerUnjoin(target Target, mediaPlayerUnjoinParams MediaPlayerUnjoinParams) *MediaPlayerUnjoin {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "unjoin"
	m := &MediaPlayerUnjoin{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerUnjoinParams,
	}
	return m
}

type MediaPlayerUnjoin struct {
	ServiceBase
	ServiceData MediaPlayerUnjoinParams `json:"service_data,omitempty"`
}
type MediaPlayerUnjoinParams struct{}

func (m *MediaPlayerUnjoin) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerUnjoin) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerVolumeDown creates the object that can be sent to Home Assistant for domain media_player, service volume_down
// "Turn a media player volume down."
func NewMediaPlayerVolumeDown(target Target, mediaPlayerVolumeDownParams MediaPlayerVolumeDownParams) *MediaPlayerVolumeDown {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "volume_down"
	m := &MediaPlayerVolumeDown{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerVolumeDownParams,
	}
	return m
}

type MediaPlayerVolumeDown struct {
	ServiceBase
	ServiceData MediaPlayerVolumeDownParams `json:"service_data,omitempty"`
}
type MediaPlayerVolumeDownParams struct{}

func (m *MediaPlayerVolumeDown) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerVolumeDown) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerVolumeMute creates the object that can be sent to Home Assistant for domain media_player, service volume_mute
// "Mute a media player's volume."
func NewMediaPlayerVolumeMute(target Target, mediaPlayerVolumeMuteParams MediaPlayerVolumeMuteParams) *MediaPlayerVolumeMute {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "volume_mute"
	m := &MediaPlayerVolumeMute{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerVolumeMuteParams,
	}
	return m
}

type MediaPlayerVolumeMute struct {
	ServiceBase
	ServiceData MediaPlayerVolumeMuteParams `json:"service_data,omitempty"`
}
type MediaPlayerVolumeMuteParams struct{}

func (m *MediaPlayerVolumeMute) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerVolumeMute) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerVolumeSet creates the object that can be sent to Home Assistant for domain media_player, service volume_set
// "Set a media player's volume level."
func NewMediaPlayerVolumeSet(target Target, mediaPlayerVolumeSetParams MediaPlayerVolumeSetParams) *MediaPlayerVolumeSet {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "volume_set"
	m := &MediaPlayerVolumeSet{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerVolumeSetParams,
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

func (m *MediaPlayerVolumeSet) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerVolumeSet) SetID(id *int) {
	m.Id = id
}

// NewMediaPlayerVolumeUp creates the object that can be sent to Home Assistant for domain media_player, service volume_up
// "Turn a media player volume up."
func NewMediaPlayerVolumeUp(target Target, mediaPlayerVolumeUpParams MediaPlayerVolumeUpParams) *MediaPlayerVolumeUp {
	serviceDomain := "media_player"
	serviceType := "call_service"
	serviceService := "volume_up"
	m := &MediaPlayerVolumeUp{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: mediaPlayerVolumeUpParams,
	}
	return m
}

type MediaPlayerVolumeUp struct {
	ServiceBase
	ServiceData MediaPlayerVolumeUpParams `json:"service_data,omitempty"`
}
type MediaPlayerVolumeUpParams struct{}

func (m *MediaPlayerVolumeUp) JSON() string {
	data, _ := json.Marshal(m)
	return string(data)
}
func (m *MediaPlayerVolumeUp) SetID(id *int) {
	m.Id = id
}
