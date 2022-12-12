package otel

import (
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
)

func buildStdoutTraceExporter() *stdouttrace.Exporter {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		panic(err)
	}
	return exporter
}
