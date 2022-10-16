package golang_otel

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func buildDefaultPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
}

func setTextMapPropagator(config *OpenTelemetryConfiguration) {
	if config.propagator == nil {
		propagator := buildDefaultPropagator()
		otel.SetTextMapPropagator(propagator)
		return
	}

	propagator := config.propagator
	otel.SetTextMapPropagator(*propagator)
}
