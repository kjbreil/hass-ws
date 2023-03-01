package model

import (
	"encoding/json"
)

type Message struct {
	ID          *int        `json:"id,omitempty"`
	Type        MessageType `json:"type"`
	AccessToken *string     `json:"access_token,omitempty"`
	Success     *bool       `json:"success,omitempty"`

	EventType *EventType `json:"event_type,omitempty"`

	Error         *Error                 `json:"error,omitempty"`
	Event         *Event                 `json:"event,omitempty"`
	Result        []Result               `json:"result,omitempty"`
	ServiceResult map[string]interface{} `json:"service_result,omitempty"`

	Raw []byte `json:"-"`
}

type MessageType string

const (
	MessageTypeAuthRequired     MessageType = "auth_required"
	MessageTypeAuth                         = "auth"
	MessageTypeAuthOK                       = "auth_ok"
	MessageTypeSubscribeEvents              = "subscribe_events"
	MessageTypeSubscribeTrigger             = "subscribe_trigger"
	MessageTypeEvent                        = "event"
	MessageTypeGetStates                    = "get_states"
	MessageTypeGetServices                  = "get_services"
)

func (mt MessageType) Valid() bool {
	switch mt {
	case MessageTypeAuthRequired,
		MessageTypeAuth,
		MessageTypeAuthOK,
		MessageTypeSubscribeEvents,
		MessageTypeSubscribeTrigger,
		MessageTypeEvent,
		MessageTypeGetStates,
		MessageTypeGetServices:
		return true
	}

	return false
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ParseMessage(msgJson []byte) (Message, error) {
	var message Message

	err := json.Unmarshal(msgJson, &message)
	if err != nil {
		return message, err
	}

	return message, nil
}
