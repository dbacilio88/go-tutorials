package queue

import (
	"errors"
	"github.com/dbacilio88/go/pkg/clients/utils"
	"github.com/dbacilio88/go/pkg/config"
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
* @author christian
* @author dbacilio88@outlook.es
* @since 17/08/2024
*
 */

type Executor interface {
	Connection() error
	Reconnect() error
	ReceiveMessage(queueName string) (<-chan amqp091.Delivery, error)
	SendMessage(queueName string, data []byte) error
	GetConnectionError() <-chan error
}

type ChannelInterface interface {
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp091.Table) (amqp091.Queue, error)
}

type MqAdapter struct {
	console    *zap.Logger
	connection *amqp091.Connection
	err        chan error
}

func NewMqAdapter(console *zap.Logger) *MqAdapter {
	return &MqAdapter{
		console: console,
		err:     make(chan error),
	}
}

func (s *MqAdapter) Connection() error {
	var url = config.GetDomainRabbitConnection()
	s.console.Info("Connecting to rabbitmq", zap.String("url", url))
	var err error
	s.connection, err = amqp091.DialConfig(url, amqp091.Config{
		Vhost: config.Config.Rabbitmq.Vhost,
	})

	go func() {
		if err != nil {
			s.console.Error("failed to connect to RabbitMq", zap.Error(err))
			return
		}
		if s.connection != nil {
			<-s.connection.NotifyClose(make(chan *amqp091.Error, 1))
			s.err <- errors.New("RabbitMq adapter already closed")
		}
	}()

	s.console.Info("Connected to RabbitMQ")
	return err
}

func (s *MqAdapter) Reconnect() error {
	err := s.Connection()
	if err != nil {
		return err
	}
	return nil
}

func (s *MqAdapter) ReceiveMessage(queueName string) (<-chan amqp091.Delivery, error) {
	s.console.Info("Receiving messages from RabbitMq", zap.String("queueName", queueName))
	channel, err := s.connection.Channel()
	s.failOnError(err, "Failed to open a channel")
	err = channel.Qos(0, 0, false)
	s.failOnError(err, "Failed to set QoS")
	err = createQueueIfNotExist(channel, queueName)
	s.failOnError(err, "Failed to create queue")
	message, err := channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	s.failOnError(err, "Failed to register a consumer")
	return message, nil
}

func (s *MqAdapter) SendMessage(queueName string, data []byte) error {
	s.console.Info("Sending messages to RabbitMq", zap.String("queueName", queueName))
	channel, err := s.connection.Channel()
	s.failOnError(err, "Failed to open a channel")
	err = createQueueIfNotExist(channel, queueName)
	s.failOnError(err, "Failed to create queue")
	err = channel.PublishWithContext(
		utils.AddParamToContext(""),
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
	s.failOnError(err, "Failed to publish a messages")
	return nil
}

func (s *MqAdapter) GetConnectionError() <-chan error {
	return s.err
}

func (s *MqAdapter) failOnError(err error, msg string) {
	if err != nil {
		s.console.Error(msg, zap.Error(err))
	}
}

func createQueueIfNotExist(ch *amqp091.Channel, queueName string) error {
	_, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}
