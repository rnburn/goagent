package grpc

import (
	"github.com/hypertrace/goagent/instrumentation/opencensus"
	traceablegrpc "github.com/hypertrace/goagent/sdk/google.golang.org/grpc"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc/stats"
)

// WrapClientHandler wraps an OpenCensus ClientHandler and returns a new one that records
// the request/response body and metadata.
func WrapClientHandler(delegate *ocgrpc.ClientHandler) stats.Handler {
	return traceablegrpc.WrapStatsHandler(delegate, opencensus.SpanFromContext)
}
