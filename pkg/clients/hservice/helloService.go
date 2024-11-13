package hservice

import (
	"context"
	proto "github.com/dbacilio88/go/proto/hello"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Executor interface {
	HelloService(context context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error)
	QueryDataService(context context.Context, req *proto.TransactionQueryRequest) (grpc.ServerStreamingClient[proto.TransactionQueryResponse], error)
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

func (a *HelloService) HelloService(context context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {

	cli := a.helloServiceCreator.HelloServiceClient(a.grpcConnection)

	response, err := cli.Hello(context, req)

	if err != nil {
		a.failOnError(err, "error client grpcs service service")
		return nil, err
	}

	return response, nil
}

func (a *HelloService) QueryDataService(context context.Context, req *proto.TransactionQueryRequest) (grpc.ServerStreamingClient[proto.TransactionQueryResponse], error) {

	cli := a.helloServiceCreator.QueryDataServiceClient(a.grpcConnection)

	stream, err := cli.ExecuteTransactionQuery(context, req)

	if err != nil {
		a.failOnError(err, "error client grpcs service service")
		return nil, err
	}

	return stream, nil
}

func (a *HelloService) failOnError(err error, msg string) {
	if err != nil {
		a.console.Error(msg, zap.Error(err))
	}
}
