package websocket

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"

	"github.com/kjbreil/hass-ws/internal/callbacks"
	"github.com/kjbreil/hass-ws/internal/handlers"
	"github.com/kjbreil/hass-ws/model"
	ws "nhooyr.io/websocket"
)

// Receiver handles receiving and processing WebSocket messages
type Receiver struct {
	ctx       context.Context
	reader    func() (*model.Message, error)
	callbacks *callbacks.Callbacks
	router    *handlers.Router
	logger    *slog.Logger
}

// NewReceiver creates a new message receiver
func NewReceiver(
	ctx context.Context,
	reader func() (*model.Message, error),
	cb *callbacks.Callbacks,
	router *handlers.Router,
	logger *slog.Logger,
) *Receiver {
	return &Receiver{
		ctx:       ctx,
		reader:    reader,
		callbacks: cb,
		router:    router,
		logger:    logger,
	}
}

// Start begins the message receiving loop
func (r *Receiver) Start() {
	go func() {
	mainLoop:
		for {
			if r.ctx.Err() != nil {
				return
			}
			message, err := r.reader()
			if err != nil {
				r.logger.Error(fmt.Sprintf("read error: %v", err), "message read error")
				switch {
				case errors.Is(err, context.Canceled),
					ws.CloseStatus(err) != -1,
					errors.Is(err, io.EOF),
					errors.Is(err, io.ErrUnexpectedEOF):
					return
				default:
					continue
				}
			}

			// Check if this is a callback response
			var callbackCh chan *model.Message
			callbackFound := false
			if message.ID != nil {
				if cb, ok := r.callbacks.Get(*message.ID); ok {
					callbackCh = cb
					callbackFound = true
				}
			}

			// Route the message
			if r.router.Route(message, callbackCh, callbackFound) && callbackFound {
				continue mainLoop
			}
		}
	}()
}
