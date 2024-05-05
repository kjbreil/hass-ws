package model

import (
	"fmt"
	"github.com/goccy/go-json"
	"strconv"
	"strings"
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

func (m *Message) DomainEntity() string {
	if m != nil && m.Event != nil && m.Event.Data != nil && m.Event.Data.EntityId != nil {
		return *m.Event.Data.EntityId
	}
	return ""
}
func (m *Message) Domain() string {
	if m.Event != nil && m.Event.Data != nil && m.Event.Data.EntityId != nil {
		return strings.Split(*m.Event.Data.EntityId, ".")[0]
	}
	return ""
}

func (m *Message) EntityID() string {
	if m.Event != nil && m.Event.Data != nil && m.Event.Data.EntityId != nil {
		return strings.Join(strings.Split(*m.Event.Data.EntityId, ".")[1:], "")
	}
	return ""
}
func (m *Message) State() string {
	if m.Event != nil &&
		m.Event.Data != nil &&
		m.Event.Data.NewState != nil &&
		m.Event.Data.NewState.State != nil {
		return *m.Event.Data.NewState.State
	}
	return ""
}

func (m *Message) OldState() *string {
	if m.Event != nil &&
		m.Event.Data != nil &&
		m.Event.Data.OldState != nil &&
		m.Event.Data.OldState.State != nil {
		return m.Event.Data.OldState.State
	}
	return nil
}

func (m *Message) StateFloat() (float64, error) {
	if m.Event != nil &&
		m.Event.Data != nil &&
		m.Event.Data.NewState != nil &&
		m.Event.Data.NewState.State != nil {
		return strconv.ParseFloat(*m.Event.Data.NewState.State, 64)
	}
	return 0, fmt.Errorf("no NewState data")
}

func (m *Message) Attributes() map[string]interface{} {
	if m.Event != nil &&
		m.Event.Data != nil &&
		m.Event.Data.NewState != nil &&
		m.Event.Data.NewState.Attributes != nil {
		return m.Event.Data.NewState.Attributes
	}
	return nil
}
func (m *Message) OldAttributes() map[string]interface{} {
	if m.Event != nil &&
		m.Event.Data != nil &&
		m.Event.Data.OldState != nil &&
		m.Event.Data.OldState.Attributes != nil {
		return m.Event.Data.OldState.Attributes
	}
	return nil
}

func (m *Message) FriendlyName() string {
	if m.Event != nil &&
		m.Event.Data != nil &&
		m.Event.Data.NewState != nil &&
		m.Event.Data.NewState.Attributes != nil {
		if fn, ok := m.Event.Data.NewState.Attributes["friendly_name"]; ok {
			return fn.(string)
		}
	}
	return ""
}

type MessageType string

const (
	MessageTypeAuthRequired      MessageType = "auth_required"
	MessageTypeAuth                          = "auth"
	MessageTypeAuthOK                        = "auth_ok"
	MessageTypeSubscribeEvents               = "subscribe_events"
	MessageTypeSubscribeTrigger              = "subscribe_trigger"
	MessageTypeEvent                         = "event"
	MessageTypePing                          = "ping"
	MessageTypeGetStates                     = "get_states"
	MessageTypeGetServices                   = "get_services"
	MessageTypeGetEntityRegistry             = "config/entity_registry/list"
	MessageTypeGetDeviceRegistry             = "config/device_registry/list"
)

func (mt MessageType) Valid() bool {
	switch mt {
	case MessageTypeAuthRequired,
		MessageTypeAuth,
		MessageTypeAuthOK,
		MessageTypeSubscribeEvents,
		MessageTypeSubscribeTrigger,
		MessageTypeEvent,
		MessageTypePing,
		MessageTypeGetStates,
		MessageTypeGetServices,
		MessageTypeGetEntityRegistry,
		MessageTypeGetDeviceRegistry:
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
