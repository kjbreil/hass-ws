package entities

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type MediaPlayer struct {
	Additional   Additional `json:"additional,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
}

func GetMediaPlayer(attributes map[string]interface{}) *MediaPlayer {
	var m MediaPlayer
	fillFields(&m, attributes)
	m.Additional = attributes
	return &m
}
