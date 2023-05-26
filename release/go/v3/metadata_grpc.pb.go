// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: v3/metadata.proto

package v3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetadataServiceClient is the client API for MetadataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetadataServiceClient interface {
	GetDatastream(ctx context.Context, in *GetDatastreamRequest, opts ...grpc.CallOption) (*GetDatastreamResponse, error)
	ListUoms(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListUomsResponse, error)
}

type metadataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataServiceClient(cc grpc.ClientConnInterface) MetadataServiceClient {
	return &metadataServiceClient{cc}
}

func (c *metadataServiceClient) GetDatastream(ctx context.Context, in *GetDatastreamRequest, opts ...grpc.CallOption) (*GetDatastreamResponse, error) {
	out := new(GetDatastreamResponse)
	err := c.cc.Invoke(ctx, "/v3.MetadataService/GetDatastream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) ListUoms(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListUomsResponse, error) {
	out := new(ListUomsResponse)
	err := c.cc.Invoke(ctx, "/v3.MetadataService/ListUoms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServiceServer is the server API for MetadataService service.
// All implementations must embed UnimplementedMetadataServiceServer
// for forward compatibility
type MetadataServiceServer interface {
	GetDatastream(context.Context, *GetDatastreamRequest) (*GetDatastreamResponse, error)
	ListUoms(context.Context, *emptypb.Empty) (*ListUomsResponse, error)
	mustEmbedUnimplementedMetadataServiceServer()
}

// UnimplementedMetadataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetadataServiceServer struct {
}

func (UnimplementedMetadataServiceServer) GetDatastream(context.Context, *GetDatastreamRequest) (*GetDatastreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatastream not implemented")
}
func (UnimplementedMetadataServiceServer) ListUoms(context.Context, *emptypb.Empty) (*ListUomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUoms not implemented")
}
func (UnimplementedMetadataServiceServer) mustEmbedUnimplementedMetadataServiceServer() {}

// UnsafeMetadataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetadataServiceServer will
// result in compilation errors.
type UnsafeMetadataServiceServer interface {
	mustEmbedUnimplementedMetadataServiceServer()
}

func RegisterMetadataServiceServer(s grpc.ServiceRegistrar, srv MetadataServiceServer) {
	s.RegisterService(&MetadataService_ServiceDesc, srv)
}

func _MetadataService_GetDatastream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatastreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetDatastream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.MetadataService/GetDatastream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetDatastream(ctx, req.(*GetDatastreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_ListUoms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).ListUoms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.MetadataService/ListUoms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).ListUoms(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MetadataService_ServiceDesc is the grpc.ServiceDesc for MetadataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetadataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v3.MetadataService",
	HandlerType: (*MetadataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDatastream",
			Handler:    _MetadataService_GetDatastream_Handler,
		},
		{
			MethodName: "ListUoms",
			Handler:    _MetadataService_ListUoms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v3/metadata.proto",
}
