// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/counter.proto

package protoconnect

import (
	context "context"
	proto "crimtech_deliverable/proto/proto"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// CounterServiceName is the fully-qualified name of the CounterService service.
	CounterServiceName = "counter.v1.CounterService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CounterServiceIncrementProcedure is the fully-qualified name of the CounterService's Increment
	// RPC.
	CounterServiceIncrementProcedure = "/counter.v1.CounterService/Increment"
)

// CounterServiceClient is a client for the counter.v1.CounterService service.
type CounterServiceClient interface {
	Increment(context.Context, *connect_go.Request[proto.IncrementRequest]) (*connect_go.Response[proto.IncrementResponse], error)
}

// NewCounterServiceClient constructs a client for the counter.v1.CounterService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCounterServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CounterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &counterServiceClient{
		increment: connect_go.NewClient[proto.IncrementRequest, proto.IncrementResponse](
			httpClient,
			baseURL+CounterServiceIncrementProcedure,
			opts...,
		),
	}
}

// counterServiceClient implements CounterServiceClient.
type counterServiceClient struct {
	increment *connect_go.Client[proto.IncrementRequest, proto.IncrementResponse]
}

// Increment calls counter.v1.CounterService.Increment.
func (c *counterServiceClient) Increment(ctx context.Context, req *connect_go.Request[proto.IncrementRequest]) (*connect_go.Response[proto.IncrementResponse], error) {
	return c.increment.CallUnary(ctx, req)
}

// CounterServiceHandler is an implementation of the counter.v1.CounterService service.
type CounterServiceHandler interface {
	Increment(context.Context, *connect_go.Request[proto.IncrementRequest]) (*connect_go.Response[proto.IncrementResponse], error)
}

// NewCounterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCounterServiceHandler(svc CounterServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	counterServiceIncrementHandler := connect_go.NewUnaryHandler(
		CounterServiceIncrementProcedure,
		svc.Increment,
		opts...,
	)
	return "/counter.v1.CounterService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CounterServiceIncrementProcedure:
			counterServiceIncrementHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCounterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCounterServiceHandler struct{}

func (UnimplementedCounterServiceHandler) Increment(context.Context, *connect_go.Request[proto.IncrementRequest]) (*connect_go.Response[proto.IncrementResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("counter.v1.CounterService.Increment is not implemented"))
}
