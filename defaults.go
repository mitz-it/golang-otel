package golang_otel

import "fmt"

const default_version = "1.0.0"

const localhost = "127.0.0.1"

const gRPC_port = 4317

const http_port = 4318

func default_gRPC_URL() string {
	return fmt.Sprintf("%s:%d", localhost, gRPC_port)
}

func default_HTTP_URL() string {
	return fmt.Sprintf("%s:%d", localhost, http_port)
}
