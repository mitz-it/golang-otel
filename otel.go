package golang_otel

import (
	"context"
	"os"
	"os/signal"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func ConfigureOpenTelemetryTracing(ctx context.Context, configure ConfigureTracing) (func(context.Context) error, error) {
	config := NewOtelConfiguration()

	configure(config)

	config.guardForEmptyServiceName()

	resource := buildResource(ctx, config)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	exporter := createExporter(ctx, config)

	batchSpanProcessor := sdktrace.NewBatchSpanProcessor(exporter)

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(resource),
		sdktrace.WithSpanProcessor(batchSpanProcessor),
	)

	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return traceProvider.Shutdown, nil
}

func CreateSignalContext() (context.Context, context.CancelFunc) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	return ctx, cancel
}
