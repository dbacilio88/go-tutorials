package rabbitmq

import (
	"fmt"
	"github.com/dbacilio88/go/pkg/adapters/mq"
	"github.com/dbacilio88/go/pkg/config"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

/**
*
* rabbit
* <p>
* rabbit file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author bxcode
* @author dbacilio88@outlook.es
* @since 7/11/2024
*
 */

type Executor interface {
	RabbitMqConnection() (amqp091.Connection, error)
}
type ManagerConnection struct {
	console         *zap.Logger
	rabbitMqAdapter mq.RabbitMqAdapter
}

// NewManagerConnection create new instance ManageConnection
func NewManagerConnection(console *zap.Logger) *ManagerConnection {
	return &ManagerConnection{
		console: console,
	}
}

func (mc *ManagerConnection) RabbitMqConnection() (amqp091.Connection, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%s",
		config.Config.Rabbitmq.Protocol,
		config.Config.Rabbitmq.User,
		config.Config.Rabbitmq.Password,
		config.Config.Rabbitmq.Host,
		config.Config.Rabbitmq.Port)

	return mc.rabbitMqAdapter.CreateConnection(url, "")
}
