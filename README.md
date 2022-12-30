# OpenTelemetry - Mitz IT

Abstraction over OpenTelemetry packages. Applying the _convention over configuration_ concept to generate observability to Mitz IT microservices.

## Installation

```bash
go get -u github.com/mitz-it/golang-otel
```

## Usage

```go
package main

import (
  otel "github.com/mitz-it/golang-otel"
)

func main() {
  ctx, cancel := otel.CreateSignalContext()
  defer cancel()

  shutdown, err := otel.ConfigureOpenTelemetryTracing(ctx, configureOpenTelemetryTracing)
  defer shutdown(ctx)

  if err != nil {
    panic(err)
  }

  //...
}

func configureOpenTelemetryTracing(config *otel.OpenTelemetryConfiguration) {
  // This is the only required property
  config.WithServiceName("otel-sample")

  // Setup the service namespace
  config.WithServiceNamespace("otel-sample-namespace")

  // The following configs are the defaults for Mitz IT OpenTelemetry package.
  // You don't need to call these methods, just showing what you can do with Mitz-IT OpenTelemetry package.
  // We are using "convention over configuration" strategy.
  config.AutoGenerateInstanceID(true)
  config.WithContextTimeout(100)
  // OpenTelemetry sends traces and metrics to the collector through gRPC or HTTP protocols.
  // The default collector port for gRPC is 4317 and for HTTP is 4318.
  config.ExportTracesTo("localhost:4317") 
  config.ExportUsing(otel.GRPC)
  // This is the default text map propagator
  propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}) 
  config.WithCustomTextMapPropagator(&propagator)
  config.WithGrpcExporterCredentials(insecure.NewCredentials())
  config.WithSpanProcessorType(otel.BATCH)
}
```

## Modules Integration

After setting up this OpenTelemetry at the application level, you can start to add instrumentations.

The following code shows how to configure OpenTelemetry to work with [golang-modules](https://github.com/mitz-it/golang-modules):

```go
package main

import (
  modules "github.com/mitz-it/golang-modules"
  otel "github.com/mitz-it/golang-otel"
  samplemodule "github.com/payly/otel-sample/src/sample-module"
  "github.com/payly/otel-sample/src/docs"
  "go.opentelemetry.io/otel/propagation"
)

func main() {
  ctx, cancel := otel.CreateSignalContext()
  defer cancel()

  shutdown, err := otel.ConfigureOpenTelemetryTracing(ctx, configureOpenTelemetryTracing)
  defer shutdown(ctx)

  if err != nil {
    panic(err)
  }

  builder := modules.NewHostBuilder()
  builder.AddModule(samplemodule.SampleModule)
  builder.ConfigureAPI(func(api *modules.API) {
    api.UseSwagger(docs.SwaggerInfo)
    api.UseOpenTelemetryMiddleware("otel-sample")
  })

  host := builder.Build()
  host.Run()
}

func configureOpenTelemetryTracing(config *otel.OpenTelemetryConfiguration) {
  config.WithServiceName("otel-sample")
}
```
