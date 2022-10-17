package otel

import (
	"context"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func buildResource(ctx context.Context, config *OpenTelemetryConfiguration) *resource.Resource {
	attributes := buildResourceAttributes(config)

	resource, err := resource.New(
		ctx,
		resource.WithAttributes(attributes...),
	)

	if err != nil {
		panic(err)
	}

	return resource
}

func buildResourceAttributes(config *OpenTelemetryConfiguration) []attribute.KeyValue {
	attributes := make([]attribute.KeyValue, 0)

	attributes = append(attributes, semconv.ServiceNameKey.String(config.serviceName))

	version := getServiceVersion(config.serviceVersion)

	attributes = append(attributes, semconv.ServiceVersionKey.String(version))

	if config.serviceNamespace != "" {
		attributes = append(attributes, semconv.ServiceNamespaceKey.String(config.serviceNamespace))
	}

	if config.autoGenerateInstanceID {
		intanceID := generateInstanceID()
		attributes = append(attributes, semconv.ServiceInstanceIDKey.String(intanceID))
	} else {
		if config.serviceInstanceID != "" {
			attributes = append(attributes, semconv.ServiceInstanceIDKey.String(config.serviceInstanceID))
		}
	}

	return attributes
}

func getServiceVersion(version string) string {
	if version == "" {
		return default_version
	}
	return version
}

func generateInstanceID() string {
	return uuid.NewString()
}
