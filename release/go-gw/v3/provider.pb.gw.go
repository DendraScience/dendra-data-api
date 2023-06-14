// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: v3/provider.proto

/*
Package v3 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v3

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_ProviderService_StreamDatapoints_0(ctx context.Context, marshaler runtime.Marshaler, client ProviderServiceClient, req *http.Request, pathParams map[string]string) (ProviderService_StreamDatapointsClient, runtime.ServerMetadata, error) {
	var protoReq ProviderStreamDatapointsRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.StreamDatapoints(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterProviderServiceHandlerServer registers the http handlers for service ProviderService to "mux".
// UnaryRPC     :call ProviderServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterProviderServiceHandlerFromEndpoint instead.
func RegisterProviderServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ProviderServiceServer) error {

	mux.Handle("POST", pattern_ProviderService_StreamDatapoints_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterProviderServiceHandlerFromEndpoint is same as RegisterProviderServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterProviderServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterProviderServiceHandler(ctx, mux, conn)
}

// RegisterProviderServiceHandler registers the http handlers for service ProviderService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterProviderServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterProviderServiceHandlerClient(ctx, mux, NewProviderServiceClient(conn))
}

// RegisterProviderServiceHandlerClient registers the http handlers for service ProviderService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ProviderServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ProviderServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ProviderServiceClient" to call the correct interceptors.
func RegisterProviderServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ProviderServiceClient) error {

	mux.Handle("POST", pattern_ProviderService_StreamDatapoints_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/v3.ProviderService/StreamDatapoints", runtime.WithHTTPPathPattern("/v3.ProviderService/StreamDatapoints"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProviderService_StreamDatapoints_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProviderService_StreamDatapoints_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ProviderService_StreamDatapoints_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v3.ProviderService", "StreamDatapoints"}, ""))
)

var (
	forward_ProviderService_StreamDatapoints_0 = runtime.ForwardResponseStream
)
