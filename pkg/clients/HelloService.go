package clients

import (
	"github.com/dbacilio88/go/pkg/clients/utils"
	proto "github.com/dbacilio88/go/proto/hello"
	"google.golang.org/grpc"
	"log"
)

type Executor interface {
	Hello() error
}

type HelloServiceClient struct {
	grpcConnection      *grpc.ClientConn
	helloServiceCreator HelloServiceCreator
}

func NewHelloServiceClient(
	grpcConnection *grpc.ClientConn,
	helloServiceCreator HelloServiceCreator) *HelloServiceClient {
	return &HelloServiceClient{
		grpcConnection:      grpcConnection,
		helloServiceCreator: helloServiceCreator,
	}
}

func (c *HelloServiceClient) Hello() error {

	cli := c.helloServiceCreator.NewClient(c.grpcConnection)

	ctx := utils.AddParamToContext("")

	tqr := proto.HelloRequest{
		Hello: &proto.Hello{
			Prefix:    "Sr",
			FirstName: "Christian",
		},
	}

	response, err := cli.Hello(ctx, &tqr)

	if err != nil {
		log.Fatal("error client grpcs hello service ", err.Error())

		return err
	}

	log.Println("response ", response)

	return nil
}
