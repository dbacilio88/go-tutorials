package command

import (
	"context"
	"errors"
	"github.com/dbacilio88/go/pkg/clients/hservice"
	proto "github.com/dbacilio88/go/proto/hello"
	"go.uber.org/zap"
	"io"
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
* @author christian
* @author dbacilio88@outlook.es
* @since 8/11/2024
*
 */

type Executor interface {
	ExecuteHelloServiceCommand(context context.Context, req *proto.HelloRequest) (interface{}, string, error)
	ExecuteQueryDataServiceCommand(context context.Context, req *proto.TransactionQueryRequest) (interface{}, string, error)
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

func (a *GrpcHelloCommand) ExecuteHelloServiceCommand(context context.Context, req *proto.HelloRequest) (interface{}, string, error) {
	start := time.Now()

	var result interface{}

	res, err := a.helloService.HelloService(context, req)

	if err != nil {
		a.failOnError(err, "error executing hello command")
		return nil, "999", err
	}

	result = res.CustomHello

	a.console.Info("successfully executed hello command", zap.Duration("elapsed", time.Since(start)))

	return result, "000", nil

}

func (a *GrpcHelloCommand) ExecuteQueryDataServiceCommand(context context.Context, req *proto.TransactionQueryRequest) (interface{}, string, error) {
	start := time.Now()

	var result interface{}

	stream, err := a.helloService.QueryDataService(context, req)

	if err != nil {
		a.failOnError(err, "error executing hello command")
		return nil, "999", err
	}

	done := make(chan bool)

	go func() {
		for {

			response, err := stream.Recv()

			if errors.Is(err, io.EOF) {
				done <- true
				return
			}

			if err != nil {
				a.failOnError(err, "error receiving data")
				done <- true
				return
			}

			a.console.Info("received data", zap.Any("data", response))
		}
	}()
	<-done
	{
		a.console.Info("successfully executed query data service command", zap.Duration("elapsed", time.Since(start)))

		return result, "000", nil
	}
}
func (a *GrpcHelloCommand) failOnError(err error, msg string) {
	if err != nil {
		a.console.Error(msg, zap.Error(err))
	}
}
