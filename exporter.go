package otel

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

type exporters struct {
	http   *otlptrace.Exporter
	grpc   *otlptrace.Exporter
	stdout *stdouttrace.Exporter
	noop   *tracetest.NoopExporter
}

func newExporters() *exporters {
	return &exporters{}
}

func buildTraceExporter(ctx context.Context, config *OpenTelemetryConfiguration) *exporters {
	exporters := newExporters()

	if config.exporterProtocol == HTTP {
		httpExporter := build_HTTP_TraceExporter(ctx, config)
		exporters.http = httpExporter
	}

	if config.exporterProtocol == GRPC {
		grpcExporter := build_GRPC_TraceExporter(ctx, config)
		exporters.grpc = grpcExporter
	}

	if config.exporterProtocol == STDOUT {
		stdoutExporter := build_STDOUT_TraceExporter()
		exporters.stdout = stdoutExporter
	}

	if config.exporterProtocol == NOOP {
		noopExporter := build_NOOP_TraceExporter()
		exporters.noop = noopExporter
	}

	return exporters
}
