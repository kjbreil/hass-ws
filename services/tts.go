package services

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

// NewTtsClearCache creates the object that can be sent to Home Assistant for domain tts, service clear_cache
// "Remove all text-to-speech cache files and RAM cache."
func NewTtsClearCache(target Target) *TtsClearCache {
	serviceDomain := "tts"
	serviceType := "call_service"
	serviceService := "clear_cache"
	t := &TtsClearCache{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: nil,
	}
	return t
}

type TtsClearCache struct {
	ServiceBase
	ServiceData interface{} `json:"service_data,omitempty"`
}

func (t *TtsClearCache) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TtsClearCache) Name() string {
	return fmt.Sprintf("%s.%s", *t.Domain, *t.Service)
}
func (t *TtsClearCache) SetID(id *int) {
	t.Id = id
}

// NewTtsCloudSay creates the object that can be sent to Home Assistant for domain tts, service cloud_say
// "Say something using text-to-speech on a media player with cloud."
func NewTtsCloudSay(target Target, ttsCloudSayParams *TtsCloudSayParams) *TtsCloudSay {
	serviceDomain := "tts"
	serviceType := "call_service"
	serviceService := "cloud_say"
	t := &TtsCloudSay{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *ttsCloudSayParams,
	}
	return t
}

type TtsCloudSay struct {
	ServiceBase
	ServiceData TtsCloudSayParams `json:"service_data,omitempty"`
}
type TtsCloudSayParams struct {
	Language *string `json:"language,omitempty"`
	Message  *string `json:"message,omitempty"`
}

func (t *TtsCloudSay) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TtsCloudSay) Name() string {
	return fmt.Sprintf("%s.%s", *t.Domain, *t.Service)
}
func (t *TtsCloudSay) SetID(id *int) {
	t.Id = id
}

// NewTtsGoogleTranslateSay creates the object that can be sent to Home Assistant for domain tts, service google_translate_say
// "Say something using text-to-speech on a media player with google_translate."
func NewTtsGoogleTranslateSay(target Target, ttsGoogleTranslateSayParams *TtsGoogleTranslateSayParams) *TtsGoogleTranslateSay {
	serviceDomain := "tts"
	serviceType := "call_service"
	serviceService := "google_translate_say"
	t := &TtsGoogleTranslateSay{
		ServiceBase: ServiceBase{
			Domain:  &serviceDomain,
			Id:      nil,
			Service: &serviceService,
			Target:  target,
			Type:    &serviceType,
		},
		ServiceData: *ttsGoogleTranslateSayParams,
	}
	return t
}

type TtsGoogleTranslateSay struct {
	ServiceBase
	ServiceData TtsGoogleTranslateSayParams `json:"service_data,omitempty"`
}
type TtsGoogleTranslateSayParams struct {
	Language *string `json:"language,omitempty"`
	Message  *string `json:"message,omitempty"`
}

func (t *TtsGoogleTranslateSay) JSON() string {
	data, _ := json.Marshal(t)
	return string(data)
}
func (t *TtsGoogleTranslateSay) Name() string {
	return fmt.Sprintf("%s.%s", *t.Domain, *t.Service)
}
func (t *TtsGoogleTranslateSay) SetID(id *int) {
	t.Id = id
}
