// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: v3/provider.proto

package v3

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

// ProviderServiceClient is the client API for ProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderServiceClient interface {
	StreamDatapoints(ctx context.Context, in *ProviderStreamDatapointsRequest, opts ...grpc.CallOption) (ProviderService_StreamDatapointsClient, error)
}

type providerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderServiceClient(cc grpc.ClientConnInterface) ProviderServiceClient {
	return &providerServiceClient{cc}
}

func (c *providerServiceClient) StreamDatapoints(ctx context.Context, in *ProviderStreamDatapointsRequest, opts ...grpc.CallOption) (ProviderService_StreamDatapointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProviderService_ServiceDesc.Streams[0], "/v3.ProviderService/StreamDatapoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &providerServiceStreamDatapointsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProviderService_StreamDatapointsClient interface {
	Recv() (*ProviderStreamDatapointsResponse, error)
	grpc.ClientStream
}

type providerServiceStreamDatapointsClient struct {
	grpc.ClientStream
}

func (x *providerServiceStreamDatapointsClient) Recv() (*ProviderStreamDatapointsResponse, error) {
	m := new(ProviderStreamDatapointsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProviderServiceServer is the server API for ProviderService service.
// All implementations must embed UnimplementedProviderServiceServer
// for forward compatibility
type ProviderServiceServer interface {
	StreamDatapoints(*ProviderStreamDatapointsRequest, ProviderService_StreamDatapointsServer) error
	mustEmbedUnimplementedProviderServiceServer()
}

// UnimplementedProviderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServiceServer struct {
}

func (UnimplementedProviderServiceServer) StreamDatapoints(*ProviderStreamDatapointsRequest, ProviderService_StreamDatapointsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamDatapoints not implemented")
}
func (UnimplementedProviderServiceServer) mustEmbedUnimplementedProviderServiceServer() {}

// UnsafeProviderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServiceServer will
// result in compilation errors.
type UnsafeProviderServiceServer interface {
	mustEmbedUnimplementedProviderServiceServer()
}

func RegisterProviderServiceServer(s grpc.ServiceRegistrar, srv ProviderServiceServer) {
	s.RegisterService(&ProviderService_ServiceDesc, srv)
}

func _ProviderService_StreamDatapoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProviderStreamDatapointsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProviderServiceServer).StreamDatapoints(m, &providerServiceStreamDatapointsServer{stream})
}

type ProviderService_StreamDatapointsServer interface {
	Send(*ProviderStreamDatapointsResponse) error
	grpc.ServerStream
}

type providerServiceStreamDatapointsServer struct {
	grpc.ServerStream
}

func (x *providerServiceStreamDatapointsServer) Send(m *ProviderStreamDatapointsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ProviderService_ServiceDesc is the grpc.ServiceDesc for ProviderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProviderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v3.ProviderService",
	HandlerType: (*ProviderServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamDatapoints",
			Handler:       _ProviderService_StreamDatapoints_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v3/provider.proto",
}
