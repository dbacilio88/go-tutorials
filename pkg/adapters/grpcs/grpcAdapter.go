package grpcs

import (
	"fmt"
	"github.com/dbacilio88/go/pkg/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"
)

type server struct {
	//proto.UnimplementedHelloServiceServer
}

type GrpcAdapter struct {
	console *zap.Logger
}

func NewManagementGrpcService(console *zap.Logger) *GrpcAdapter {
	return &GrpcAdapter{
		console: console,
	}
}

func (a GrpcAdapter) GRPCConnectionClientManager(wg *sync.WaitGroup) (*grpc.ClientConn, error) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	addr := fmt.Sprintf("%s:%s", config.Config.GrpcServer.Host, config.Config.GrpcServer.Port)
	a.console.Info("Connecting to client GRPC", zap.String("url", addr))
	client, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		a.console.Error("error connecting to grpc server", zap.Error(err))
		return nil, err
	}

	a.console.Info("Connected to client GRPC", zap.String("state", client.GetState().String()))
	return client, nil
}
