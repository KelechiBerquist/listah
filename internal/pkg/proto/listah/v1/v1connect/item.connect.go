// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: listah/v1/item.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	v1 "cornucopia/listah/internal/pkg/proto/listah/v1"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ItemServiceName is the fully-qualified name of the ItemService service.
	ItemServiceName = "listah.v1.ItemService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ItemServiceReadProcedure is the fully-qualified name of the ItemService's Read RPC.
	ItemServiceReadProcedure = "/listah.v1.ItemService/Read"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	itemServiceServiceDescriptor    = v1.File_listah_v1_item_proto.Services().ByName("ItemService")
	itemServiceReadMethodDescriptor = itemServiceServiceDescriptor.Methods().ByName("Read")
)

// ItemServiceClient is a client for the listah.v1.ItemService service.
type ItemServiceClient interface {
	Read(context.Context, *connect.Request[v1.ItemServiceReadRequest]) (*connect.Response[v1.ItemServiceReadResponse], error)
}

// NewItemServiceClient constructs a client for the listah.v1.ItemService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewItemServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ItemServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &itemServiceClient{
		read: connect.NewClient[v1.ItemServiceReadRequest, v1.ItemServiceReadResponse](
			httpClient,
			baseURL+ItemServiceReadProcedure,
			connect.WithSchema(itemServiceReadMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// itemServiceClient implements ItemServiceClient.
type itemServiceClient struct {
	read *connect.Client[v1.ItemServiceReadRequest, v1.ItemServiceReadResponse]
}

// Read calls listah.v1.ItemService.Read.
func (c *itemServiceClient) Read(ctx context.Context, req *connect.Request[v1.ItemServiceReadRequest]) (*connect.Response[v1.ItemServiceReadResponse], error) {
	return c.read.CallUnary(ctx, req)
}

// ItemServiceHandler is an implementation of the listah.v1.ItemService service.
type ItemServiceHandler interface {
	Read(context.Context, *connect.Request[v1.ItemServiceReadRequest]) (*connect.Response[v1.ItemServiceReadResponse], error)
}

// NewItemServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewItemServiceHandler(svc ItemServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	itemServiceReadHandler := connect.NewUnaryHandler(
		ItemServiceReadProcedure,
		svc.Read,
		connect.WithSchema(itemServiceReadMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/listah.v1.ItemService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ItemServiceReadProcedure:
			itemServiceReadHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedItemServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedItemServiceHandler struct{}

func (UnimplementedItemServiceHandler) Read(context.Context, *connect.Request[v1.ItemServiceReadRequest]) (*connect.Response[v1.ItemServiceReadResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("listah.v1.ItemService.Read is not implemented"))
}