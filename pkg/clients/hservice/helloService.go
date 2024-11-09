package hservice

import (
	"context"
	proto "github.com/dbacilio88/go/proto/hello"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Executor interface {
	Hello(context context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error)
}

type HelloService struct {
	console             *zap.Logger
	grpcConnection      *grpc.ClientConn
	helloServiceCreator HelloServiceCreator
}

func NewHelloService(
	console *zap.Logger,
	grpcConnection *grpc.ClientConn,
	helloServiceCreator HelloServiceCreator,
) HelloService {
	return HelloService{
		console:             console,
		grpcConnection:      grpcConnection,
		helloServiceCreator: helloServiceCreator,
	}
}

func (a *HelloService) Hello(context context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {

	cli := a.helloServiceCreator.NewClient(a.grpcConnection)

	response, err := cli.Hello(context, req)

	if err != nil {
		a.failOnError(err, "error client grpcs service service")
		return nil, err
	}

	return response, nil
}

func (a *HelloService) failOnError(err error, msg string) {
	if err != nil {
		a.console.Error(msg, zap.Error(err))
	}
}
