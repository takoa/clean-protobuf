// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: routeguide/v1/routeguide.proto

package routeguidev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RouteGuideService_GetFeature_FullMethodName   = "/routeguide.v1.RouteGuideService/GetFeature"
	RouteGuideService_ListFeatures_FullMethodName = "/routeguide.v1.RouteGuideService/ListFeatures"
	RouteGuideService_RecordRoute_FullMethodName  = "/routeguide.v1.RouteGuideService/RecordRoute"
	RouteGuideService_RouteChat_FullMethodName    = "/routeguide.v1.RouteGuideService/RouteChat"
)

// RouteGuideServiceClient is the client API for RouteGuideService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteGuideServiceClient interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(ctx context.Context, in *GetFeatureRequest, opts ...grpc.CallOption) (*GetFeatureResponse, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(ctx context.Context, in *ListFeaturesRequest, opts ...grpc.CallOption) (RouteGuideService_ListFeaturesClient, error)
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuideService_RecordRouteClient, error)
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuideService_RouteChatClient, error)
}

type routeGuideServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteGuideServiceClient(cc grpc.ClientConnInterface) RouteGuideServiceClient {
	return &routeGuideServiceClient{cc}
}

func (c *routeGuideServiceClient) GetFeature(ctx context.Context, in *GetFeatureRequest, opts ...grpc.CallOption) (*GetFeatureResponse, error) {
	out := new(GetFeatureResponse)
	err := c.cc.Invoke(ctx, RouteGuideService_GetFeature_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideServiceClient) ListFeatures(ctx context.Context, in *ListFeaturesRequest, opts ...grpc.CallOption) (RouteGuideService_ListFeaturesClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuideService_ServiceDesc.Streams[0], RouteGuideService_ListFeatures_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideServiceListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuideService_ListFeaturesClient interface {
	Recv() (*ListFeaturesResponse, error)
	grpc.ClientStream
}

type routeGuideServiceListFeaturesClient struct {
	grpc.ClientStream
}

func (x *routeGuideServiceListFeaturesClient) Recv() (*ListFeaturesResponse, error) {
	m := new(ListFeaturesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideServiceClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuideService_RecordRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuideService_ServiceDesc.Streams[1], RouteGuideService_RecordRoute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideServiceRecordRouteClient{stream}
	return x, nil
}

type RouteGuideService_RecordRouteClient interface {
	Send(*RecordRouteRequest) error
	CloseAndRecv() (*RecordRouteResponse, error)
	grpc.ClientStream
}

type routeGuideServiceRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeGuideServiceRecordRouteClient) Send(m *RecordRouteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideServiceRecordRouteClient) CloseAndRecv() (*RecordRouteResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RecordRouteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideServiceClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuideService_RouteChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuideService_ServiceDesc.Streams[2], RouteGuideService_RouteChat_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideServiceRouteChatClient{stream}
	return x, nil
}

type RouteGuideService_RouteChatClient interface {
	Send(*RouteChatRequest) error
	Recv() (*RouteChatResponse, error)
	grpc.ClientStream
}

type routeGuideServiceRouteChatClient struct {
	grpc.ClientStream
}

func (x *routeGuideServiceRouteChatClient) Send(m *RouteChatRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideServiceRouteChatClient) Recv() (*RouteChatResponse, error) {
	m := new(RouteChatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuideServiceServer is the server API for RouteGuideService service.
// All implementations must embed UnimplementedRouteGuideServiceServer
// for forward compatibility
type RouteGuideServiceServer interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(context.Context, *GetFeatureRequest) (*GetFeatureResponse, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(*ListFeaturesRequest, RouteGuideService_ListFeaturesServer) error
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(RouteGuideService_RecordRouteServer) error
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(RouteGuideService_RouteChatServer) error
	mustEmbedUnimplementedRouteGuideServiceServer()
}

// UnimplementedRouteGuideServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRouteGuideServiceServer struct {
}

func (UnimplementedRouteGuideServiceServer) GetFeature(context.Context, *GetFeatureRequest) (*GetFeatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
}
func (UnimplementedRouteGuideServiceServer) ListFeatures(*ListFeaturesRequest, RouteGuideService_ListFeaturesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFeatures not implemented")
}
func (UnimplementedRouteGuideServiceServer) RecordRoute(RouteGuideService_RecordRouteServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordRoute not implemented")
}
func (UnimplementedRouteGuideServiceServer) RouteChat(RouteGuideService_RouteChatServer) error {
	return status.Errorf(codes.Unimplemented, "method RouteChat not implemented")
}
func (UnimplementedRouteGuideServiceServer) mustEmbedUnimplementedRouteGuideServiceServer() {}

// UnsafeRouteGuideServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteGuideServiceServer will
// result in compilation errors.
type UnsafeRouteGuideServiceServer interface {
	mustEmbedUnimplementedRouteGuideServiceServer()
}

func RegisterRouteGuideServiceServer(s grpc.ServiceRegistrar, srv RouteGuideServiceServer) {
	s.RegisterService(&RouteGuideService_ServiceDesc, srv)
}

func _RouteGuideService_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServiceServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteGuideService_GetFeature_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServiceServer).GetFeature(ctx, req.(*GetFeatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuideService_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListFeaturesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServiceServer).ListFeatures(m, &routeGuideServiceListFeaturesServer{stream})
}

type RouteGuideService_ListFeaturesServer interface {
	Send(*ListFeaturesResponse) error
	grpc.ServerStream
}

type routeGuideServiceListFeaturesServer struct {
	grpc.ServerStream
}

func (x *routeGuideServiceListFeaturesServer) Send(m *ListFeaturesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuideService_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServiceServer).RecordRoute(&routeGuideServiceRecordRouteServer{stream})
}

type RouteGuideService_RecordRouteServer interface {
	SendAndClose(*RecordRouteResponse) error
	Recv() (*RecordRouteRequest, error)
	grpc.ServerStream
}

type routeGuideServiceRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeGuideServiceRecordRouteServer) SendAndClose(m *RecordRouteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideServiceRecordRouteServer) Recv() (*RecordRouteRequest, error) {
	m := new(RecordRouteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuideService_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServiceServer).RouteChat(&routeGuideServiceRouteChatServer{stream})
}

type RouteGuideService_RouteChatServer interface {
	Send(*RouteChatResponse) error
	Recv() (*RouteChatRequest, error)
	grpc.ServerStream
}

type routeGuideServiceRouteChatServer struct {
	grpc.ServerStream
}

func (x *routeGuideServiceRouteChatServer) Send(m *RouteChatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideServiceRouteChatServer) Recv() (*RouteChatRequest, error) {
	m := new(RouteChatRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuideService_ServiceDesc is the grpc.ServiceDesc for RouteGuideService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteGuideService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.v1.RouteGuideService",
	HandlerType: (*RouteGuideServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteGuideService_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _RouteGuideService_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _RouteGuideService_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _RouteGuideService_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "routeguide/v1/routeguide.proto",
}
