package otel

import (
	"context"
	"os"
	"os/signal"
	"time"
)

func CreateSignalContext() (context.Context, context.CancelFunc) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	return ctx, cancel
}

func buildTimeoutContext(ctx context.Context, config *OpenTelemetryConfiguration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, config.contextTimeout*time.Second)
	return ctx, cancel
}
