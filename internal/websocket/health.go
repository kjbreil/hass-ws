package websocket

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/kjbreil/hass-ws/internal/callbacks"
	"github.com/kjbreil/hass-ws/model"
)

// HealthMonitor monitors WebSocket connection health via ping/pong
type HealthMonitor struct {
	ctx          context.Context
	sender       func(*model.Message, chan *model.Message) (int, error)
	callbacks    *callbacks.Callbacks
	reconnect    func() error
	closeConn    func() error
	logger       *slog.Logger
	pingInterval time.Duration
	pongTimeout  time.Duration
}

// NewHealthMonitor creates a new health monitor
func NewHealthMonitor(
	ctx context.Context,
	sender func(*model.Message, chan *model.Message) (int, error),
	cb *callbacks.Callbacks,
	reconnect func() error,
	closeConn func() error,
	logger *slog.Logger,
) *HealthMonitor {
	return &HealthMonitor{
		ctx:          ctx,
		sender:       sender,
		callbacks:    cb,
		reconnect:    reconnect,
		closeConn:    closeConn,
		logger:       logger,
		pingInterval: time.Second * 10,
		pongTimeout:  time.Second * 5,
	}
}

// Start begins the health monitoring loop
func (h *HealthMonitor) Start() {
	go func() {
		ticker := time.NewTicker(h.pingInterval)

		for {
			select {
			case <-h.ctx.Done():
				return
			case <-ticker.C:
				callback := make(chan *model.Message)
				id, err := h.sender(&model.Message{
					Type: model.MessageTypePing,
				}, callback)
				if err != nil {
					h.logger.Error(fmt.Sprintf("ping send error: %v", err), "ping could not be sent")
					h.handleReconnect()
					return
				}
				go h.waitForPong(id, callback)
			}
		}
	}()
}

// waitForPong waits for a pong response or triggers reconnect
func (h *HealthMonitor) waitForPong(id int, callback chan *model.Message) {
	defer h.callbacks.Delete(id)
	restartTicker := time.NewTicker(h.pongTimeout)
	select {
	case <-callback:
		// Pong received, all good
	case <-restartTicker.C:
		h.logger.Error("pong not received attempting reconnect")
		h.handleReconnect()
	}
}

// handleReconnect attempts to reconnect and panics if it fails
func (h *HealthMonitor) handleReconnect() {
	err := h.reconnect()
	if err != nil {
		h.logger.Error(fmt.Sprintf("reconnect failed: %v", err))
		closeErr := h.closeConn()
		if closeErr != nil {
			h.logger.Error(fmt.Sprintf("close error: %v", closeErr))
		}
		panic(errors.New("connection lost and cannot be reconnected"))
	}
	h.logger.Info("reconnect succeeded")
}
