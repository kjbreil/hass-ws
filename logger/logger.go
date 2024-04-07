package logger

import (
	"context"
	"fmt"
	"log/slog"
)

type WrapHandler struct {
	Handler slog.Handler
}

func (w *WrapHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return w.Handler.Enabled(ctx, level)
}

func (w *WrapHandler) Handle(ctx context.Context, record slog.Record) error {
	record.Message = fmt.Sprintf("(hass-ws) %s", record.Message)
	return w.Handler.Handle(ctx, record)
}

func (w *WrapHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return w.Handler.WithAttrs(attrs)
}

func (w *WrapHandler) WithGroup(name string) slog.Handler {
	return w.Handler.WithGroup(name)
}

func NewWrapHandler(handler slog.Handler) *WrapHandler {
	return &WrapHandler{
		Handler: handler,
	}
}
