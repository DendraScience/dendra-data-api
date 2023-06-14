// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: v3/transformer.proto

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

// TransformerServiceClient is the client API for TransformerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransformerServiceClient interface {
	TransformMany(ctx context.Context, in *TransformManyRequest, opts ...grpc.CallOption) (*TransformManyResponse, error)
}

type transformerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransformerServiceClient(cc grpc.ClientConnInterface) TransformerServiceClient {
	return &transformerServiceClient{cc}
}

func (c *transformerServiceClient) TransformMany(ctx context.Context, in *TransformManyRequest, opts ...grpc.CallOption) (*TransformManyResponse, error) {
	out := new(TransformManyResponse)
	err := c.cc.Invoke(ctx, "/v3.TransformerService/TransformMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransformerServiceServer is the server API for TransformerService service.
// All implementations must embed UnimplementedTransformerServiceServer
// for forward compatibility
type TransformerServiceServer interface {
	TransformMany(context.Context, *TransformManyRequest) (*TransformManyResponse, error)
	mustEmbedUnimplementedTransformerServiceServer()
}

// UnimplementedTransformerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransformerServiceServer struct {
}

func (UnimplementedTransformerServiceServer) TransformMany(context.Context, *TransformManyRequest) (*TransformManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransformMany not implemented")
}
func (UnimplementedTransformerServiceServer) mustEmbedUnimplementedTransformerServiceServer() {}

// UnsafeTransformerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransformerServiceServer will
// result in compilation errors.
type UnsafeTransformerServiceServer interface {
	mustEmbedUnimplementedTransformerServiceServer()
}

func RegisterTransformerServiceServer(s grpc.ServiceRegistrar, srv TransformerServiceServer) {
	s.RegisterService(&TransformerService_ServiceDesc, srv)
}

func _TransformerService_TransformMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransformManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransformerServiceServer).TransformMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.TransformerService/TransformMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransformerServiceServer).TransformMany(ctx, req.(*TransformManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransformerService_ServiceDesc is the grpc.ServiceDesc for TransformerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransformerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v3.TransformerService",
	HandlerType: (*TransformerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransformMany",
			Handler:    _TransformerService_TransformMany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v3/transformer.proto",
}