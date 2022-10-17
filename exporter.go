package otel

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
)

func buildTraceExporter(ctx context.Context, config *OpenTelemetryConfiguration) *otlptrace.Exporter {
	if config.exporterProtocol == HTTP {
		exporter := buildHTTPTraceExporter(ctx, config)
		return exporter
	}

	exporter := buildGRPCTraceExporter(ctx, config)
	return exporter
}
