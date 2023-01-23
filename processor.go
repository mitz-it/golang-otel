package otel

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func buildSpanProcessor(config *OpenTelemetryConfiguration, exporters *exporters) sdktrace.SpanProcessor {
	switch config.exporterProtocol {
	case GRPC:
		exporter := exporters.grpc
		return buildSpanProcessorForOTLP(config, exporter)
	case HTTP:
		exporter := exporters.http
		return buildSpanProcessorForOTLP(config, exporter)
	case STDOUT:
		exporter := exporters.stdout
		return buildSpanProcessorForSTDOUT(config, exporter)
	case NOOP:
		exporter := exporters.noop
		return buildSpanProcessorForNOOP(config, exporter)
	default:
		exporter := exporters.grpc
		return buildSpanProcessorForOTLP(config, exporter)
	}
}

func buildSpanProcessorForSTDOUT(config *OpenTelemetryConfiguration, exporter *stdouttrace.Exporter) sdktrace.SpanProcessor {
	if config.spanProcessorType == BATCH {
		spanProcessor := sdktrace.NewBatchSpanProcessor(exporter)
		return spanProcessor
	}
	spanProcessor := sdktrace.NewSimpleSpanProcessor(exporter)
	return spanProcessor
}

func buildSpanProcessorForOTLP(config *OpenTelemetryConfiguration, exporter *otlptrace.Exporter) sdktrace.SpanProcessor {
	if config.spanProcessorType == BATCH {
		spanProcessor := sdktrace.NewBatchSpanProcessor(exporter)
		return spanProcessor
	}
	spanProcessor := sdktrace.NewSimpleSpanProcessor(exporter)
	return spanProcessor
}

func buildSpanProcessorForNOOP(config *OpenTelemetryConfiguration, exporter *tracetest.NoopExporter) sdktrace.SpanProcessor {
	if config.spanProcessorType == BATCH {
		spanProcessor := sdktrace.NewBatchSpanProcessor(exporter)
		return spanProcessor
	}
	spanProcessor := sdktrace.NewSimpleSpanProcessor(exporter)
	return spanProcessor
}
