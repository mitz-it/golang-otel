package otel

import (
	"errors"
	"time"

	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type ExporterProtocol int

const (
	GRPC ExporterProtocol = iota
	HTTP
)

type SpanProcessorType int

const (
	BATCH SpanProcessorType = iota
	SIMPLE
)

type OpenTelemetryConfiguration struct {
	serviceName            string
	serviceNamespace       string
	serviceVersion         string
	serviceInstanceID      string
	collectorURL           string
	autoGenerateInstanceID bool
	insecure               bool
	contextTimeout         time.Duration
	grpcCredentials        credentials.TransportCredentials
	exporterProtocol       ExporterProtocol
	spanProcessorType      SpanProcessorType
	propagator             *propagation.TextMapPropagator
}

type ConfigureTracing func(*OpenTelemetryConfiguration)

func (config *OpenTelemetryConfiguration) WithServiceName(name string) {
	config.serviceName = name
}

func (config *OpenTelemetryConfiguration) WithServiceNamespace(namespace string) {
	config.serviceNamespace = namespace
}

func (config *OpenTelemetryConfiguration) WithServiceVersion(version string) {
	config.serviceVersion = version
}

func (config *OpenTelemetryConfiguration) WithServiceInstanceID(instanceID string) {
	config.serviceInstanceID = instanceID
}

func (config *OpenTelemetryConfiguration) AutoGenerateInstanceID(autoGenerate bool) {
	config.autoGenerateInstanceID = autoGenerate
}

func (config *OpenTelemetryConfiguration) UseInsecureCredentialsForHTTPExporter(insecure bool) {
	config.insecure = insecure
}

func (config *OpenTelemetryConfiguration) WithContextTimeout(timeout time.Duration) {
	config.contextTimeout = timeout
}

func (config *OpenTelemetryConfiguration) WithGrpcExporterCredentials(credentials credentials.TransportCredentials) {
	config.grpcCredentials = credentials
}

func (config *OpenTelemetryConfiguration) WithSpanProcessorType(processorType SpanProcessorType) {
	config.spanProcessorType = processorType
}

func (config *OpenTelemetryConfiguration) WithCustomTextMapPropagator(propagator *propagation.TextMapPropagator) {
	config.propagator = propagator
}

func (config *OpenTelemetryConfiguration) ExportTracesTo(url string) {
	config.collectorURL = url
}

func (config *OpenTelemetryConfiguration) ExportUsing(protocol ExporterProtocol) {
	config.exporterProtocol = protocol
}

func (config *OpenTelemetryConfiguration) guardForEmptyServiceName() {
	if config.serviceName == "" {
		err := errors.New("service name cannot be empty")
		panic(err)
	}
}

func NewOpenTelemetryConfiguraion() *OpenTelemetryConfiguration {
	config := new(OpenTelemetryConfiguration)
	config.exporterProtocol = GRPC
	config.serviceVersion = default_version
	config.autoGenerateInstanceID = true
	config.insecure = true
	config.contextTimeout = 100
	config.grpcCredentials = insecure.NewCredentials()
	config.spanProcessorType = BATCH
	config.propagator = nil
	return config
}
