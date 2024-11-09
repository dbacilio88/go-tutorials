package hservice

import (
	proto "github.com/dbacilio88/go/proto/hello"
	"google.golang.org/grpc"
)

type GrpcClientCreator struct {
}

type HelloServiceCreator interface {
	NewClient(connection *grpc.ClientConn) proto.HelloServiceClient
}

func NewGrpcClientCreator() *GrpcClientCreator {
	return &GrpcClientCreator{}
}

func (r *GrpcClientCreator) NewClient(connection *grpc.ClientConn) proto.HelloServiceClient {
	return proto.NewHelloServiceClient(connection)
}
