package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type MediaPlayer struct {
	Additional                   map[string]interface{} `json:"additional,omitempty"`
	DeviceClass                  *string                `json:"device_class,omitempty"`
	FriendlyName                 *string                `json:"friendly_name,omitempty"`
	GroupMembers                 *[]string              `json:"group_members,omitempty"`
	MediaImageRemotelyAccessible *bool                  `json:"media_image_remotely_accessible,omitempty"`
	MediaImageUrl                *string                `json:"media_image_url,omitempty"`
	SoundMode                    *string                `json:"sound_mode,omitempty"`
	SoundModeList                *[]string              `json:"sound_mode_list,omitempty"`
	Source                       *string                `json:"source,omitempty"`
	SourceList                   *[]string              `json:"source_list,omitempty"`
	SupportedFeatures            *int                   `json:"supported_features,omitempty"`
}

func GetMediaPlayer(attributes map[string]interface{}) *MediaPlayer {
	var m MediaPlayer
	FillFields(&m, attributes)
	m.Additional = attributes
	return &m
}
