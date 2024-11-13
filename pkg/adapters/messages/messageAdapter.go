package messages

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dbacilio88/go/pkg/adapters/command"
	"github.com/dbacilio88/go/pkg/adapters/queue"
	"github.com/dbacilio88/go/pkg/clients/utils"
	"github.com/dbacilio88/go/proto/hello"
	"github.com/dbacilio88/go/services/validation"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

/**
*
* messagingAdapter
* <p>
* messagingAdapter file
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
	ReceiveMessages(queue string)
	ProcessMessages()
}

type MessageAdapter struct {
	console          *zap.Logger
	rabbitAdapter    queue.Executor
	grpcHelloCommand command.Executor
	validatorService validation.Executor
}

func NewMessageAdapter(console *zap.Logger,
	rabbitAdapter queue.Executor,
	grpcHelloCommand command.Executor,
	validatorService validation.Executor,
) *MessageAdapter {
	return &MessageAdapter{
		console:          console,
		rabbitAdapter:    rabbitAdapter,
		grpcHelloCommand: grpcHelloCommand,
		validatorService: validatorService,
	}
}

func (a *MessageAdapter) ReceiveMessages(queue string) {
	a.console.Info("received messages", zap.String("queue", queue))
	messages, err := a.rabbitAdapter.ReceiveMessage(queue)

	if err != nil {
		a.console.Error(err.Error())
		return
	}

	for {
		go func() {
			for msg := range messages {
				var code string
				var err error
				var response interface{}

				var data map[string]interface{}
				switch msg.ContentType {
				case "application/json", "text/plain":
					data, err = a.parseBodyMessage(queue, msg.Body)
					a.failOnError(err, "Failed to parse messages")
					break
				default:
					a.console.Error("Message contentType no supported", zap.String("contentType", msg.ContentType))
				}

				if data != nil {
					a.console.Info("messages", zap.Any("data", data))
					request, err := a.validatorService.ValidateRequest(data)
					if err != nil {
						a.failOnError(err, "Failed to validate request")
						return
					}

					a.console.Info("validation data", zap.Any("request", request))

					ctx := utils.AddParamToContext(uuid.New().String())
					reqH := hello.HelloRequest{
						Hello: &hello.Hello{
							Prefix:    request.Prefix,
							FirstName: request.Name,
						},
					}
					response, code, err = a.grpcHelloCommand.ExecuteHelloServiceCommand(ctx, &reqH)

					if err != nil {
						a.failOnError(err, "Failed to execute hello command")
						return
					}

					a.console.Info("data from server grpc", zap.Any("response", response), zap.String("code", code))

					tqr := hello.TransactionQueryRequest{
						NameView: "DATA_LOG",
					}

					parameters := []Parameter{
						{"DL_ID", "ID", 1},
						{"DL_UUID", "UUID", 2},
						{"DL_NAME", "NAME", 3},
					}

					var tqd []*hello.TransactionQueryDetail

					for i, filter := range parameters {
						fmt.Println(i)
						tqd = append(tqd, &hello.TransactionQueryDetail{
							MappingSqlModel: filter.ViewColumn,
							Name:            filter.NameMapping,
							Order:           (int32)(i),
						})
					}

					tqr.TransactionQueryDetail = tqd

					responses, code, err := a.grpcHelloCommand.ExecuteQueryDataServiceCommand(ctx, &tqr)
					if err != nil {
						a.failOnError(err, "Failed to execute query data service command")
						return
					}
					a.console.Info("data from server grpc", zap.Any("responses", responses), zap.String("code", code))

				}

				err = msg.Ack(false)
				if err != nil {
					a.failOnError(err, "Failed to ack messages")
					return
				}
			}
		}()

		if channel := <-a.rabbitAdapter.GetConnectionError(); channel != nil {
			time.Sleep(30 * time.Second)
			err := a.rabbitAdapter.Reconnect()
			if err != nil {
				a.failOnError(err, "Failed to reconnect to RabbitMQ")
				continue
			}
			messages, err = a.rabbitAdapter.ReceiveMessage(queue)
			if err != nil {
				a.failOnError(err, "Failed to receive messages")
				return
			}
		}
	}
}
func (a *MessageAdapter) parseBodyMessage(queueName string, body []byte) (map[string]interface{}, error) {
	a.console.Info("Message receive", zap.String("Message", string(body)), zap.String("QueueName", queueName))
	var result map[string]interface{}

	if len(body) == 0 {
		a.console.Warn("Message body is empty")
		return nil, errors.New("messages body is empty")
	}

	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.New("error parsing body messages")
	}

	return result, nil
}

func (a *MessageAdapter) ProcessMessage() {

}

func (a *MessageAdapter) failOnError(err error, msg string) {
	if err != nil {
		fullMessage := fmt.Sprintf("%s: %s", msg, err.Error())
		a.console.Error(fullMessage, zap.Error(err))
	}
}

type Parameter struct {
	ViewColumn  string
	NameMapping string
	Order       int
}
