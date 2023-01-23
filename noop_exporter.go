package otel

import (
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func build_NOOP_TraceExporter() *tracetest.NoopExporter {

	exporter := tracetest.NewNoopExporter()

	return exporter
}
