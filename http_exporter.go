package otel

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

func buildHTTPTraceExporter(ctx context.Context, config *OpenTelemetryConfiguration) *otlptrace.Exporter {
	collectorURL := config.collectorURL
	insecure := config.insecure
	client := buildHTTPClient(collectorURL, insecure)
	exporter, err := otlptrace.New(ctx, client)

	if err != nil {
		panic(err)
	}

	return exporter
}

func buildHTTPClient(collectorURL string, insecure bool) otlptrace.Client {

	options := buildHTTPClientOptions(collectorURL, insecure)

	client := otlptracehttp.NewClient(options...)

	return client
}

func buildHTTPClientOptions(collectorURL string, insecure bool) []otlptracehttp.Option {
	options := make([]otlptracehttp.Option, 0)
	url := buildHTTPexporterURL(collectorURL)
	options = append(options, otlptracehttp.WithEndpoint(url))

	if insecure {
		options = append(options, otlptracehttp.WithInsecure())
	}

	return options
}

func buildHTTPexporterURL(collectorURL string) string {
	if collectorURL == "" {
		return default_HTTP_URL()
	}

	return collectorURL
}

func default_HTTP_URL() string {
	return fmt.Sprintf("%s:%d", localhost, http_port)
}
