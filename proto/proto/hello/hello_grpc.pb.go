// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: hello.proto

package hello

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	HelloService_Hello_FullMethodName = "/proto.HelloService/Hello"
)

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	// unary
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, HelloService_Hello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility.
type HelloServiceServer interface {
	// unary
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHelloServiceServer struct{}

func (UnimplementedHelloServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}
func (UnimplementedHelloServiceServer) testEmbeddedByValue()                      {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	// If the following call pancis, it indicates UnimplementedHelloServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloService_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

const (
	QueryDataService_ExecuteTransactionQuery_FullMethodName = "/proto.QueryDataService/ExecuteTransactionQuery"
)

// QueryDataServiceClient is the client API for QueryDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryDataServiceClient interface {
	ExecuteTransactionQuery(ctx context.Context, in *TransactionQueryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TransactionQueryResponse], error)
}

type queryDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryDataServiceClient(cc grpc.ClientConnInterface) QueryDataServiceClient {
	return &queryDataServiceClient{cc}
}

func (c *queryDataServiceClient) ExecuteTransactionQuery(ctx context.Context, in *TransactionQueryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TransactionQueryResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &QueryDataService_ServiceDesc.Streams[0], QueryDataService_ExecuteTransactionQuery_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[TransactionQueryRequest, TransactionQueryResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type QueryDataService_ExecuteTransactionQueryClient = grpc.ServerStreamingClient[TransactionQueryResponse]

// QueryDataServiceServer is the server API for QueryDataService service.
// All implementations must embed UnimplementedQueryDataServiceServer
// for forward compatibility.
type QueryDataServiceServer interface {
	ExecuteTransactionQuery(*TransactionQueryRequest, grpc.ServerStreamingServer[TransactionQueryResponse]) error
	mustEmbedUnimplementedQueryDataServiceServer()
}

// UnimplementedQueryDataServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueryDataServiceServer struct{}

func (UnimplementedQueryDataServiceServer) ExecuteTransactionQuery(*TransactionQueryRequest, grpc.ServerStreamingServer[TransactionQueryResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ExecuteTransactionQuery not implemented")
}
func (UnimplementedQueryDataServiceServer) mustEmbedUnimplementedQueryDataServiceServer() {}
func (UnimplementedQueryDataServiceServer) testEmbeddedByValue()                          {}

// UnsafeQueryDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryDataServiceServer will
// result in compilation errors.
type UnsafeQueryDataServiceServer interface {
	mustEmbedUnimplementedQueryDataServiceServer()
}

func RegisterQueryDataServiceServer(s grpc.ServiceRegistrar, srv QueryDataServiceServer) {
	// If the following call pancis, it indicates UnimplementedQueryDataServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&QueryDataService_ServiceDesc, srv)
}

func _QueryDataService_ExecuteTransactionQuery_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransactionQueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryDataServiceServer).ExecuteTransactionQuery(m, &grpc.GenericServerStream[TransactionQueryRequest, TransactionQueryResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type QueryDataService_ExecuteTransactionQueryServer = grpc.ServerStreamingServer[TransactionQueryResponse]

// QueryDataService_ServiceDesc is the grpc.ServiceDesc for QueryDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.QueryDataService",
	HandlerType: (*QueryDataServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteTransactionQuery",
			Handler:       _QueryDataService_ExecuteTransactionQuery_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hello.proto",
}