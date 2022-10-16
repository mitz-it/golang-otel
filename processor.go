package golang_otel

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func buildSpanProcessor(config *OpenTelemetryConfiguration, exporter *otlptrace.Exporter) sdktrace.SpanProcessor {
	if config.spanProcessorType == BATCH {
		spanProcessor := buildBatchSpanProcessor(exporter)
		return spanProcessor
	}

	spanProcessor := buildSimpleSpanProcessor(exporter)
	return spanProcessor
}

func buildBatchSpanProcessor(exporter *otlptrace.Exporter) sdktrace.SpanProcessor {
	batchSpanProcessor := sdktrace.NewBatchSpanProcessor(exporter)
	return batchSpanProcessor
}

func buildSimpleSpanProcessor(exporter *otlptrace.Exporter) sdktrace.SpanProcessor {
	simpleSpanProcessor := sdktrace.NewSimpleSpanProcessor(exporter)
	return simpleSpanProcessor
}
