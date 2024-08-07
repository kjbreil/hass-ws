package services

import (
	"fmt"
	gojson "github.com/goccy/go-json"
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
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
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
	data, _ := gojson.Marshal(t)
	return string(data)
}
func (t *TtsClearCache) Targets() []string {
	return t.Target.EntityId
}
func (t *TtsClearCache) Name() string {
	return fmt.Sprintf("%s.%s", *t.Domain, *t.Service)
}

// NewTtsCloudSay creates the object that can be sent to Home Assistant for domain tts, service cloud_say
// "Say something using text-to-speech on a media player with cloud."
func NewTtsCloudSay(target Target) *TtsCloudSay {
	serviceDomain := "tts"
	serviceType := "call_service"
	serviceService := "cloud_say"
	t := &TtsCloudSay{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: TtsCloudSayParams{},
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

func (t *TtsCloudSay) Language(language string) *TtsCloudSay {
	t.ServiceData.Language = &language
	return t
}
func (t *TtsCloudSay) Message(message string) *TtsCloudSay {
	t.ServiceData.Message = &message
	return t
}
func (t *TtsCloudSay) JSON() string {
	data, _ := gojson.Marshal(t)
	return string(data)
}
func (t *TtsCloudSay) Targets() []string {
	return t.Target.EntityId
}
func (t *TtsCloudSay) Name() string {
	return fmt.Sprintf("%s.%s", *t.Domain, *t.Service)
}

// NewTtsGoogleTranslateSay creates the object that can be sent to Home Assistant for domain tts, service google_translate_say
// "Say something using text-to-speech on a media player with google_translate."
func NewTtsGoogleTranslateSay(target Target) *TtsGoogleTranslateSay {
	serviceDomain := "tts"
	serviceType := "call_service"
	serviceService := "google_translate_say"
	t := &TtsGoogleTranslateSay{
		ServiceBase: ServiceBase{
			Domain:         &serviceDomain,
			Id:             nil,
			ReturnResponse: false,
			Service:        &serviceService,
			Target:         target,
			Type:           &serviceType,
		},
		ServiceData: TtsGoogleTranslateSayParams{},
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

func (t *TtsGoogleTranslateSay) Language(language string) *TtsGoogleTranslateSay {
	t.ServiceData.Language = &language
	return t
}
func (t *TtsGoogleTranslateSay) Message(message string) *TtsGoogleTranslateSay {
	t.ServiceData.Message = &message
	return t
}
func (t *TtsGoogleTranslateSay) JSON() string {
	data, _ := gojson.Marshal(t)
	return string(data)
}
func (t *TtsGoogleTranslateSay) Targets() []string {
	return t.Target.EntityId
}
func (t *TtsGoogleTranslateSay) Name() string {
	return fmt.Sprintf("%s.%s", *t.Domain, *t.Service)
}
