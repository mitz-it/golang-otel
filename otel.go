package otel

import (
	"context"

	"go.opentelemetry.io/otel"
)

func ConfigureOpenTelemetryTracing(ctx context.Context, configure ConfigureTracing) (func(context.Context) error, error) {
	config := NewOpenTelemetryConfiguraion()

	configure(config)

	config.guardForEmptyServiceName()

	resource := buildResource(ctx, config)

	ctx, cancel := buildTimeoutContext(ctx, config)
	defer cancel()

	exporters := buildTraceExporter(ctx, config)

	spanProcessor := buildSpanProcessor(config, exporters)

	tracerProvider := buildTracerProvider(resource, spanProcessor)

	otel.SetTracerProvider(tracerProvider)

	setTextMapPropagator(config)

	return tracerProvider.Shutdown, nil
}
