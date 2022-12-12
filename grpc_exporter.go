package otel

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"google.golang.org/grpc"
)

func build_GRPC_TraceExporter(ctx context.Context, config *OpenTelemetryConfiguration) *otlptrace.Exporter {
	connection := buildGRPCconnection(ctx, config)
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(connection))

	if err != nil {
		panic(err)
	}

	return exporter
}

func buildGRPCconnection(ctx context.Context, config *OpenTelemetryConfiguration) *grpc.ClientConn {
	url := buildGRPCexporterURL(config.collectorURL)
	credentials := config.grpcCredentials
	connection, err := grpc.DialContext(ctx, url, grpc.WithTransportCredentials(credentials), grpc.WithBlock())

	if err != nil {
		panic(err)
	}

	return connection
}

func buildGRPCexporterURL(collectorURL string) string {
	if collectorURL == "" {
		return default_gRPC_URL()
	}
	return collectorURL
}

func default_gRPC_URL() string {
	return fmt.Sprintf("%s:%d", localhost, gRPC_port)
}
