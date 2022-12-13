package otel

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
)

type exporters struct {
	http   *otlptrace.Exporter
	grpc   *otlptrace.Exporter
	stdout *stdouttrace.Exporter
}

func newExporters() *exporters {
	return &exporters{}
}

func buildTraceExporter(ctx context.Context, config *OpenTelemetryConfiguration) *exporters {
	exporters := newExporters()

	httpExporter := build_HTTP_TraceExporter(ctx, config)
	grpcExporter := build_GRPC_TraceExporter(ctx, config)
	stdoutExporter := build_STDOUT_TraceExporter()

	exporters.http = httpExporter
	exporters.grpc = grpcExporter
	exporters.stdout = stdoutExporter

	return exporters
}
