package grpcs

import (
	"fmt"
	"github.com/dbacilio88/go/pkg/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcAdapter struct {
	console *zap.Logger
}

func NewManagementGrpcService(console *zap.Logger) *GrpcAdapter {
	return &GrpcAdapter{
		console: console,
	}
}

func (a *GrpcAdapter) GRPCConnectionClientManager() (*grpc.ClientConn, error) {
	fmt.Println("Host ", config.Config.Grpc.Server)
	a.console.Info("connecting to server grpcs", zap.String("url", config.Config.Grpc.Server))
	return grpc.NewClient(config.Config.Grpc.Server, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func (a *GrpcAdapter) EnsureConnection(connection *grpc.ClientConn) error {
	a.console.Info("start ensure connection")
	a.console.Info("connection grpcs down, retry get connection")

	conn, err := a.GRPCConnectionClientManager()
	if err != nil {
		a.failOnError(err, "error retrying get connection grpcs")
		return err
	}

	a.console.Info("connection grpcs restored")

	connection = conn

	defer func(connection *grpc.ClientConn) {
		_ = connection.Close()
	}(connection)

	return nil

}
func (a *GrpcAdapter) failOnError(err error, msg string) {
	if err != nil {
		a.console.Error(msg, zap.Error(err))
	}
}
