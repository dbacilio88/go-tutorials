package clients

import (
	proto "github.com/dbacilio88/go/proto/hello"
	"google.golang.org/grpc"
)

type HelloGrpcCreator struct {
}

type HelloServiceCreator interface {
	NewClient(connection *grpc.ClientConn) proto.HelloServiceClient
}

func NewHelloCreator() *HelloGrpcCreator {
	return &HelloGrpcCreator{}
}

func (r *HelloGrpcCreator) NewClient(connection *grpc.ClientConn) proto.HelloServiceClient {
	return proto.NewHelloServiceClient(connection)
}
