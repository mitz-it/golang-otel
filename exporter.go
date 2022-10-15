package golang_otel

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createExporter(ctx context.Context, config *OpenTelemetryConfiguration) *otlptrace.Exporter {
	if config.exporterProtocol == HTTP {
		url := getHTTPexporterURL(config)
		client := otlptracehttp.NewClient(otlptracehttp.WithEndpoint(url))
		exporter, _ := otlptrace.New(ctx, client)
		return exporter
	}
	url := getGRPCexporterURL(config)
	connection, _ := grpc.DialContext(ctx, url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	exporter, _ := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(connection))
	return exporter
}

func getHTTPexporterURL(config *OpenTelemetryConfiguration) string {
	if config.collectorURL == "" {
		return default_HTTP_URL()
	}
	return config.collectorURL
}

func getGRPCexporterURL(config *OpenTelemetryConfiguration) string {
	if config.collectorURL == "" {
		return default_gRPC_URL()
	}
	return config.collectorURL
}
