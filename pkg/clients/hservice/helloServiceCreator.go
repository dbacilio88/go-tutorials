package hservice

import (
	proto "github.com/dbacilio88/go/proto/hello"
	"google.golang.org/grpc"
)

type GrpcClientCreator struct {
}

type HelloServiceCreator interface {
	HelloServiceClient(connection *grpc.ClientConn) proto.HelloServiceClient
	QueryDataServiceClient(connection *grpc.ClientConn) proto.QueryDataServiceClient
}

func NewGrpcClientCreator() *GrpcClientCreator {
	return &GrpcClientCreator{}
}

func (r *GrpcClientCreator) HelloServiceClient(connection *grpc.ClientConn) proto.HelloServiceClient {
	return proto.NewHelloServiceClient(connection)
}

func (r *GrpcClientCreator) QueryDataServiceClient(connection *grpc.ClientConn) proto.QueryDataServiceClient {
	return proto.NewQueryDataServiceClient(connection)
}
