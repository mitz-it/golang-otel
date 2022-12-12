package otel

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
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
	default:
		exporter := exporters.grpc
		return buildSpanProcessorForOTLP(config, exporter)
	}
}

func buildSpanProcessorForSTDOUT(config *OpenTelemetryConfiguration, exporter *stdouttrace.Exporter) sdktrace.SpanProcessor {
	if config.spanProcessorType == BATCH {
		spanProcessor := buildBatchSpanProcessorForSTDOUT(exporter)
		return spanProcessor
	}
	spanProcessor := buildSimpleSpanProcessorForSTDOUT(exporter)
	return spanProcessor
}

func buildSpanProcessorForOTLP(config *OpenTelemetryConfiguration, exporter *otlptrace.Exporter) sdktrace.SpanProcessor {
	if config.spanProcessorType == BATCH {
		spanProcessor := buildBatchSpanProcessorForOTLP(exporter)
		return spanProcessor
	}
	spanProcessor := buildSimpleSpanProcessorForOTLP(exporter)
	return spanProcessor
}

func buildBatchSpanProcessorForOTLP(exporter *otlptrace.Exporter) sdktrace.SpanProcessor {
	batchSpanProcessor := sdktrace.NewBatchSpanProcessor(exporter)
	return batchSpanProcessor
}

func buildBatchSpanProcessorForSTDOUT(exporter *stdouttrace.Exporter) sdktrace.SpanProcessor {
	batchSpanProcessor := sdktrace.NewBatchSpanProcessor(exporter)
	return batchSpanProcessor
}

func buildSimpleSpanProcessorForOTLP(exporter *otlptrace.Exporter) sdktrace.SpanProcessor {
	simpleSpanProcessor := sdktrace.NewSimpleSpanProcessor(exporter)
	return simpleSpanProcessor
}

func buildSimpleSpanProcessorForSTDOUT(exporter *stdouttrace.Exporter) sdktrace.SpanProcessor {
	simpleSpanProcessor := sdktrace.NewSimpleSpanProcessor(exporter)
	return simpleSpanProcessor
}
