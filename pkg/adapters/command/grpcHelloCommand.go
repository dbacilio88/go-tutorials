package command

import (
	"context"
	"github.com/dbacilio88/go/pkg/clients/hservice"
	proto "github.com/dbacilio88/go/proto/hello"
	"go.uber.org/zap"
	"time"
)

/**
*
* grpcHelloCommand
* <p>
* grpcHelloCommand file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author bxcode
* @author dbacilio88@outlook.es
* @since 8/11/2024
*
 */

type Executor interface {
	ExecuteHelloCommand(context context.Context, req *proto.HelloRequest) (interface{}, string, error)
}

type GrpcHelloCommand struct {
	console      *zap.Logger
	helloService hservice.Executor
}

func NewGrpcHelloCommand(console *zap.Logger, helloService hservice.Executor) *GrpcHelloCommand {
	return &GrpcHelloCommand{
		console:      console,
		helloService: helloService,
	}
}

func (a *GrpcHelloCommand) ExecuteHelloCommand(context context.Context, req *proto.HelloRequest) (interface{}, string, error) {
	start := time.Now()

	var result interface{}

	res, err := a.helloService.Hello(context, req)

	if err != nil {
		a.failOnError(err, "error executing hello command")
		return nil, "999", err
	}

	result = res.CustomHello

	a.console.Info("successfully executed hello command", zap.Duration("elapsed", time.Since(start)))

	return result, "000", nil

}

func (a *GrpcHelloCommand) failOnError(err error, msg string) {
	if err != nil {
		a.console.Error(msg, zap.Error(err))
	}
}
