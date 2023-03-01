package services

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
//
// NewTtsClearCache creates the object that can be sent to Home Assistant for domain tts, service clear_cache
// "Remove all text-to-speech cache files and RAM cache."
func NewTtsClearCache(entities []string) *TtsClearCache {
	serviceDomain := "tts"
	serviceType := "call_service"
	serviceService := "clear_cache"
	t := &TtsClearCache{
		Domain:      &serviceDomain,
		Id:          nil,
		Service:     &serviceService,
		ServiceData: struct{}{},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return t
}

type TtsClearCache struct {
	Id          *int     `json:"id"`
	Type        *string  `json:"type"`
	Domain      *string  `json:"domain"`
	Service     *string  `json:"service"`
	ServiceData struct{} `json:"service_data,omitempty"`
	Target      struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (t *TtsClearCache) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TtsClearCache) SetID(id *int) {
	t.Id = id
}

// NewTtsCloudSay creates the object that can be sent to Home Assistant for domain tts, service cloud_say
// "Say something using text-to-speech on a media player with cloud."
func NewTtsCloudSay(entities []string, language *string, message *string) *TtsCloudSay {
	serviceDomain := "tts"
	serviceType := "call_service"
	serviceService := "cloud_say"
	t := &TtsCloudSay{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Language *string `json:"language,omitempty"`
			Message  *string `json:"message,omitempty"`
		}{
			Language: language,
			Message:  message,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return t
}

type TtsCloudSay struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Language *string `json:"language,omitempty"`
		Message  *string `json:"message,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (t *TtsCloudSay) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TtsCloudSay) SetID(id *int) {
	t.Id = id
}

// NewTtsGoogleTranslateSay creates the object that can be sent to Home Assistant for domain tts, service google_translate_say
// "Say something using text-to-speech on a media player with google_translate."
func NewTtsGoogleTranslateSay(entities []string, language *string, message *string) *TtsGoogleTranslateSay {
	serviceDomain := "tts"
	serviceType := "call_service"
	serviceService := "google_translate_say"
	t := &TtsGoogleTranslateSay{
		Domain:  &serviceDomain,
		Id:      nil,
		Service: &serviceService,
		ServiceData: struct {
			Language *string `json:"language,omitempty"`
			Message  *string `json:"message,omitempty"`
		}{
			Language: language,
			Message:  message,
		},
		Target: struct {
			EntityId []string `json:"entity_id,omitempty"`
		}{EntityId: entities},
		Type: &serviceType,
	}
	return t
}

type TtsGoogleTranslateSay struct {
	Id          *int    `json:"id"`
	Type        *string `json:"type"`
	Domain      *string `json:"domain"`
	Service     *string `json:"service"`
	ServiceData struct {
		Language *string `json:"language,omitempty"`
		Message  *string `json:"message,omitempty"`
	} `json:"service_data,omitempty"`
	Target struct {
		EntityId []string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func (t *TtsGoogleTranslateSay) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TtsGoogleTranslateSay) SetID(id *int) {
	t.Id = id
}
