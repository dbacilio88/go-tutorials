package utils

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func AddParamToContext(entityUUID string) context.Context {
	md := metadata.Pairs("entityUUID", entityUUID)
	return metadata.NewOutgoingContext(context.Background(), md)
}
