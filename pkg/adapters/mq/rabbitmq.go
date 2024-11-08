package mq

import (
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

/**
*
* rabbitmq
* <p>
* rabbitmq file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author bxcode
* @author dbacilio88@outlook.es
* @since 17/08/2024
*
 */

type Executor interface {
	CreateConnection(url string, host string) (*amqp091.Connection, error)
	Consumer(connection *amqp091.Connection, queueName string) (<-chan amqp091.Delivery, error)
}
type RabbitMqAdapter struct {
	console *zap.Logger
}

func NewRabbitMqAdapter(console *zap.Logger) *RabbitMqAdapter {
	return &RabbitMqAdapter{
		console: console,
	}
}
func (a *RabbitMqAdapter) CreateConnection(url string, host string) (*amqp091.Connection, error) {
	conn, err := amqp091.DialConfig(url, amqp091.Config{
		Vhost: host,
	})
	if err != nil {
		a.console.Error("failed to connect to RabbitMq", zap.Error(err))
		return nil, err
	}
	return conn, nil
}
func (adapter *RabbitMqAdapter) Consumer(connection *amqp091.Connection, queueName string) (<-chan amqp091.Delivery, error) {
	return nil, nil
}
