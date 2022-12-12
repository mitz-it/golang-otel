package otel

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
)

type exporters struct {
	otlp   *otlptrace.Exporter
	stdout *stdouttrace.Exporter
}

func newExporters(otlp *otlptrace.Exporter, stdout *stdouttrace.Exporter) *exporters {
	return &exporters{
		otlp:   otlp,
		stdout: stdout,
	}
}

func buildTraceExporter(ctx context.Context, config *OpenTelemetryConfiguration) *exporters {
	if config.exporterProtocol == STDOUT {
		exporter := buildStdoutTraceExporter()
		return newExporters(nil, exporter)
	}

	if config.exporterProtocol == HTTP {
		exporter := buildHTTPTraceExporter(ctx, config)
		return newExporters(exporter, nil)
	}

	exporter := buildGRPCTraceExporter(ctx, config)
	return newExporters(exporter, nil)
}
