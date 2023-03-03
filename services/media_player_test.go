package services

import "testing"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

func TestMediaPlayerClearPlaylist_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerClearPlaylist
		want   string
	}{{
		fields: NewMediaPlayerClearPlaylist(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"clear_playlist\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerJoin_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerJoin
		want   string
	}{{
		fields: NewMediaPlayerJoin(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"join\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerMediaNextTrack_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerMediaNextTrack
		want   string
	}{{
		fields: NewMediaPlayerMediaNextTrack(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"media_next_track\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerMediaPause_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerMediaPause
		want   string
	}{{
		fields: NewMediaPlayerMediaPause(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"media_pause\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerMediaPlay_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerMediaPlay
		want   string
	}{{
		fields: NewMediaPlayerMediaPlay(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"media_play\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerMediaPlayPause_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerMediaPlayPause
		want   string
	}{{
		fields: NewMediaPlayerMediaPlayPause(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"media_play_pause\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerMediaPreviousTrack_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerMediaPreviousTrack
		want   string
	}{{
		fields: NewMediaPlayerMediaPreviousTrack(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"media_previous_track\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerMediaSeek_JSON(t *testing.T) {
	seekPosition := 1.2

	tests := []struct {
		name   string
		fields *MediaPlayerMediaSeek
		want   string
	}{{
		fields: NewMediaPlayerMediaSeek(Targets("climate.kitchen"), &MediaPlayerMediaSeekParams{SeekPosition: &seekPosition}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"media_seek\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"seek_position\":1.2}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerMediaStop_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerMediaStop
		want   string
	}{{
		fields: NewMediaPlayerMediaStop(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"media_stop\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerPlayMedia_JSON(t *testing.T) {
	enqueue := Enqueueadd
	mediaContentId := "data"
	mediaContentType := "data"

	tests := []struct {
		name   string
		fields *MediaPlayerPlayMedia
		want   string
	}{{
		fields: NewMediaPlayerPlayMedia(Targets("climate.kitchen"), &MediaPlayerPlayMediaParams{
			Enqueue:          &enqueue,
			MediaContentId:   &mediaContentId,
			MediaContentType: &mediaContentType,
		}),
		name: "base",
		want: "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"play_media\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"enqueue\":\"add\",\"media_content_id\":\"data\",\"media_content_type\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerRepeatSet_JSON(t *testing.T) {
	repeat := Repeatall

	tests := []struct {
		name   string
		fields *MediaPlayerRepeatSet
		want   string
	}{{
		fields: NewMediaPlayerRepeatSet(Targets("climate.kitchen"), &MediaPlayerRepeatSetParams{Repeat: &repeat}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"repeat_set\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"repeat\":\"all\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerSelectSoundMode_JSON(t *testing.T) {
	soundMode := "data"

	tests := []struct {
		name   string
		fields *MediaPlayerSelectSoundMode
		want   string
	}{{
		fields: NewMediaPlayerSelectSoundMode(Targets("climate.kitchen"), &MediaPlayerSelectSoundModeParams{SoundMode: &soundMode}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"select_sound_mode\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"sound_mode\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerSelectSource_JSON(t *testing.T) {
	source := "data"

	tests := []struct {
		name   string
		fields *MediaPlayerSelectSource
		want   string
	}{{
		fields: NewMediaPlayerSelectSource(Targets("climate.kitchen"), &MediaPlayerSelectSourceParams{Source: &source}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"select_source\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"source\":\"data\"}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerShuffleSet_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerShuffleSet
		want   string
	}{{
		fields: NewMediaPlayerShuffleSet(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"shuffle_set\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerToggle_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerToggle
		want   string
	}{{
		fields: NewMediaPlayerToggle(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"toggle\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerTurnOff_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerTurnOff
		want   string
	}{{
		fields: NewMediaPlayerTurnOff(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"turn_off\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerTurnOn_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerTurnOn
		want   string
	}{{
		fields: NewMediaPlayerTurnOn(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"turn_on\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerUnjoin_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerUnjoin
		want   string
	}{{
		fields: NewMediaPlayerUnjoin(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"unjoin\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerVolumeDown_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerVolumeDown
		want   string
	}{{
		fields: NewMediaPlayerVolumeDown(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"volume_down\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerVolumeMute_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerVolumeMute
		want   string
	}{{
		fields: NewMediaPlayerVolumeMute(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"volume_mute\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerVolumeSet_JSON(t *testing.T) {
	volumeLevel := 1.2

	tests := []struct {
		name   string
		fields *MediaPlayerVolumeSet
		want   string
	}{{
		fields: NewMediaPlayerVolumeSet(Targets("climate.kitchen"), &MediaPlayerVolumeSetParams{VolumeLevel: &volumeLevel}),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"volume_set\",\"target\":{\"entity_id\":[\"climate.kitchen\"]},\"service_data\":{\"volume_level\":1.2}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMediaPlayerVolumeUp_JSON(t *testing.T) {

	tests := []struct {
		name   string
		fields *MediaPlayerVolumeUp
		want   string
	}{{
		fields: NewMediaPlayerVolumeUp(Targets("climate.kitchen")),
		name:   "base",
		want:   "{\"id\":null,\"type\":\"call_service\",\"domain\":\"media_player\",\"service\":\"volume_up\",\"target\":{\"entity_id\":[\"climate.kitchen\"]}}",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.JSON(); got != tt.want {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
