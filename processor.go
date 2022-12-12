package otel

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func buildSpanProcessor(config *OpenTelemetryConfiguration, exporters *exporters) sdktrace.SpanProcessor {
	if exporter := exporters.otlp; config.exporterProtocol == GRPC || config.exporterProtocol == HTTP {
		if config.spanProcessorType == BATCH {
			spanProcessor := buildBatchSpanProcessorForOTLP(exporter)
			return spanProcessor
		}
		spanProcessor := buildSimpleSpanProcessorForOTLP(exporter)
		return spanProcessor
	}

	exporter := exporters.stdout

	if config.spanProcessorType == BATCH {
		spanProcessor := buildBatchSpanProcessorForSTDOUT(exporter)
		return spanProcessor
	}
	spanProcessor := buildSimpleSpanProcessorForSTDOUT(exporter)
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
