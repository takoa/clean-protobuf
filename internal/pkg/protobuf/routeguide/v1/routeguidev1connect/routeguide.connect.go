// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: routeguide/v1/routeguide.proto

package routeguidev1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/takoa/clean-protobuf/internal/pkg/protobuf/routeguide/v1"
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
	// RouteGuideServiceName is the fully-qualified name of the RouteGuideService service.
	RouteGuideServiceName = "routeguide.v1.RouteGuideService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RouteGuideServiceGetFeatureProcedure is the fully-qualified name of the RouteGuideService's
	// GetFeature RPC.
	RouteGuideServiceGetFeatureProcedure = "/routeguide.v1.RouteGuideService/GetFeature"
	// RouteGuideServiceListFeaturesProcedure is the fully-qualified name of the RouteGuideService's
	// ListFeatures RPC.
	RouteGuideServiceListFeaturesProcedure = "/routeguide.v1.RouteGuideService/ListFeatures"
	// RouteGuideServiceRecordRouteProcedure is the fully-qualified name of the RouteGuideService's
	// RecordRoute RPC.
	RouteGuideServiceRecordRouteProcedure = "/routeguide.v1.RouteGuideService/RecordRoute"
	// RouteGuideServiceRouteChatProcedure is the fully-qualified name of the RouteGuideService's
	// RouteChat RPC.
	RouteGuideServiceRouteChatProcedure = "/routeguide.v1.RouteGuideService/RouteChat"
)

// RouteGuideServiceClient is a client for the routeguide.v1.RouteGuideService service.
type RouteGuideServiceClient interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(context.Context, *connect_go.Request[v1.GetFeatureRequest]) (*connect_go.Response[v1.GetFeatureResponse], error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(context.Context, *connect_go.Request[v1.ListFeaturesRequest]) (*connect_go.ServerStreamForClient[v1.ListFeaturesResponse], error)
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(context.Context) *connect_go.ClientStreamForClient[v1.RecordRouteRequest, v1.RecordRouteResponse]
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(context.Context) *connect_go.BidiStreamForClient[v1.RouteChatRequest, v1.RouteChatResponse]
}

// NewRouteGuideServiceClient constructs a client for the routeguide.v1.RouteGuideService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRouteGuideServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) RouteGuideServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &routeGuideServiceClient{
		getFeature: connect_go.NewClient[v1.GetFeatureRequest, v1.GetFeatureResponse](
			httpClient,
			baseURL+RouteGuideServiceGetFeatureProcedure,
			opts...,
		),
		listFeatures: connect_go.NewClient[v1.ListFeaturesRequest, v1.ListFeaturesResponse](
			httpClient,
			baseURL+RouteGuideServiceListFeaturesProcedure,
			opts...,
		),
		recordRoute: connect_go.NewClient[v1.RecordRouteRequest, v1.RecordRouteResponse](
			httpClient,
			baseURL+RouteGuideServiceRecordRouteProcedure,
			opts...,
		),
		routeChat: connect_go.NewClient[v1.RouteChatRequest, v1.RouteChatResponse](
			httpClient,
			baseURL+RouteGuideServiceRouteChatProcedure,
			opts...,
		),
	}
}

// routeGuideServiceClient implements RouteGuideServiceClient.
type routeGuideServiceClient struct {
	getFeature   *connect_go.Client[v1.GetFeatureRequest, v1.GetFeatureResponse]
	listFeatures *connect_go.Client[v1.ListFeaturesRequest, v1.ListFeaturesResponse]
	recordRoute  *connect_go.Client[v1.RecordRouteRequest, v1.RecordRouteResponse]
	routeChat    *connect_go.Client[v1.RouteChatRequest, v1.RouteChatResponse]
}

// GetFeature calls routeguide.v1.RouteGuideService.GetFeature.
func (c *routeGuideServiceClient) GetFeature(ctx context.Context, req *connect_go.Request[v1.GetFeatureRequest]) (*connect_go.Response[v1.GetFeatureResponse], error) {
	return c.getFeature.CallUnary(ctx, req)
}

// ListFeatures calls routeguide.v1.RouteGuideService.ListFeatures.
func (c *routeGuideServiceClient) ListFeatures(ctx context.Context, req *connect_go.Request[v1.ListFeaturesRequest]) (*connect_go.ServerStreamForClient[v1.ListFeaturesResponse], error) {
	return c.listFeatures.CallServerStream(ctx, req)
}

// RecordRoute calls routeguide.v1.RouteGuideService.RecordRoute.
func (c *routeGuideServiceClient) RecordRoute(ctx context.Context) *connect_go.ClientStreamForClient[v1.RecordRouteRequest, v1.RecordRouteResponse] {
	return c.recordRoute.CallClientStream(ctx)
}

// RouteChat calls routeguide.v1.RouteGuideService.RouteChat.
func (c *routeGuideServiceClient) RouteChat(ctx context.Context) *connect_go.BidiStreamForClient[v1.RouteChatRequest, v1.RouteChatResponse] {
	return c.routeChat.CallBidiStream(ctx)
}

// RouteGuideServiceHandler is an implementation of the routeguide.v1.RouteGuideService service.
type RouteGuideServiceHandler interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(context.Context, *connect_go.Request[v1.GetFeatureRequest]) (*connect_go.Response[v1.GetFeatureResponse], error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(context.Context, *connect_go.Request[v1.ListFeaturesRequest], *connect_go.ServerStream[v1.ListFeaturesResponse]) error
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(context.Context, *connect_go.ClientStream[v1.RecordRouteRequest]) (*connect_go.Response[v1.RecordRouteResponse], error)
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(context.Context, *connect_go.BidiStream[v1.RouteChatRequest, v1.RouteChatResponse]) error
}

// NewRouteGuideServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRouteGuideServiceHandler(svc RouteGuideServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	routeGuideServiceGetFeatureHandler := connect_go.NewUnaryHandler(
		RouteGuideServiceGetFeatureProcedure,
		svc.GetFeature,
		opts...,
	)
	routeGuideServiceListFeaturesHandler := connect_go.NewServerStreamHandler(
		RouteGuideServiceListFeaturesProcedure,
		svc.ListFeatures,
		opts...,
	)
	routeGuideServiceRecordRouteHandler := connect_go.NewClientStreamHandler(
		RouteGuideServiceRecordRouteProcedure,
		svc.RecordRoute,
		opts...,
	)
	routeGuideServiceRouteChatHandler := connect_go.NewBidiStreamHandler(
		RouteGuideServiceRouteChatProcedure,
		svc.RouteChat,
		opts...,
	)
	return "/routeguide.v1.RouteGuideService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RouteGuideServiceGetFeatureProcedure:
			routeGuideServiceGetFeatureHandler.ServeHTTP(w, r)
		case RouteGuideServiceListFeaturesProcedure:
			routeGuideServiceListFeaturesHandler.ServeHTTP(w, r)
		case RouteGuideServiceRecordRouteProcedure:
			routeGuideServiceRecordRouteHandler.ServeHTTP(w, r)
		case RouteGuideServiceRouteChatProcedure:
			routeGuideServiceRouteChatHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRouteGuideServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRouteGuideServiceHandler struct{}

func (UnimplementedRouteGuideServiceHandler) GetFeature(context.Context, *connect_go.Request[v1.GetFeatureRequest]) (*connect_go.Response[v1.GetFeatureResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("routeguide.v1.RouteGuideService.GetFeature is not implemented"))
}

func (UnimplementedRouteGuideServiceHandler) ListFeatures(context.Context, *connect_go.Request[v1.ListFeaturesRequest], *connect_go.ServerStream[v1.ListFeaturesResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("routeguide.v1.RouteGuideService.ListFeatures is not implemented"))
}

func (UnimplementedRouteGuideServiceHandler) RecordRoute(context.Context, *connect_go.ClientStream[v1.RecordRouteRequest]) (*connect_go.Response[v1.RecordRouteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("routeguide.v1.RouteGuideService.RecordRoute is not implemented"))
}

func (UnimplementedRouteGuideServiceHandler) RouteChat(context.Context, *connect_go.BidiStream[v1.RouteChatRequest, v1.RouteChatResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("routeguide.v1.RouteGuideService.RouteChat is not implemented"))
}
