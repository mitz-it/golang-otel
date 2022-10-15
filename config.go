package golang_otel

import (
	"errors"

	"github.com/google/uuid"
)

type ExporterProtocol int

const (
	GRPC ExporterProtocol = iota
	HTTP
)

type OpenTelemetryConfiguration struct {
	serviceName            string
	serviceNamespace       string
	serviceVersion         string
	serviceInstanceID      string
	autoGenerateInstanceID bool
	exporterProtocol       ExporterProtocol
	collectorURL           string
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

func generateInstanceID() string {
	return uuid.NewString()
}

func DefaultOpenTelemetryConfiguration() *OpenTelemetryConfiguration {
	config := new(OpenTelemetryConfiguration)
	config.exporterProtocol = GRPC
	config.collectorURL = default_gRPC_URL()
	config.serviceVersion = default_version
	config.autoGenerateInstanceID = true
	return config
}

func NewOtelConfiguration() *OpenTelemetryConfiguration {
	return new(OpenTelemetryConfiguration)
}
