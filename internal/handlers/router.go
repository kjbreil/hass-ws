package handlers

import (
	"github.com/kjbreil/hass-ws/model"
)

// Router handles routing of incoming messages to appropriate handlers
type Router struct {
	OnMessage   func(model.Message)
	OnType      *model.OnTypeHandlers
	OnEntity    *model.OnEntityHandlers
	OnUnhandled func(model.Message)
}

// Route routes a message to the appropriate handler
func (r *Router) Route(message *model.Message, callbackCh chan *model.Message, callbackFound bool) bool {
	// First, call OnMessage if set
	if r.OnMessage != nil {
		r.OnMessage(*message)
	}

	// If this is a callback response, send it to the callback channel
	if message.ID != nil && callbackFound {
		callbackCh <- message
		return true
	}

	// Try type-based handlers
	onTypeRan := r.OnType.Run(message)

	// Try entity-specific handlers
	onEntityRan := r.OnEntity.Run(message)

	// If no handlers ran and OnUnhandled is set, call it
	if !onTypeRan && !onEntityRan && r.OnUnhandled != nil {
		r.OnUnhandled(*message)
	}

	return onTypeRan || onEntityRan
}
