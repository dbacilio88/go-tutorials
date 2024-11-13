package main

import (
	"github.com/dbacilio88/go/pkg/adapters/command"
	"github.com/dbacilio88/go/pkg/adapters/grpcs"
	"github.com/dbacilio88/go/pkg/adapters/messages"
	"github.com/dbacilio88/go/pkg/adapters/queue"
	"github.com/dbacilio88/go/pkg/adapters/ssh"
	"github.com/dbacilio88/go/pkg/clients/hservice"
	"github.com/dbacilio88/go/pkg/config"
	"github.com/dbacilio88/go/pkg/config/logger"
	"github.com/dbacilio88/go/pkg/config/rabbitmq"
	"github.com/dbacilio88/go/pkg/server"
	"github.com/dbacilio88/go/pkg/task"
	"github.com/dbacilio88/go/services/validation"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
)

/**
*
* main
* <p>
* main file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author bxcode
* @author dbacilio88@outlook.es
* @since 3/08/2024
*
 */

func main() {

	config.Load("./")

	console, err := logger.LogConfiguration(config.Config.Server.Logging)

	if err != nil {
		console.Error("error while initializing logger", zap.Error(err))
		return
	}

	undo := zap.RedirectStdLog(console)

	defer undo()

	// create instance grpc client adapter:
	grpcAdapter := grpcs.NewManagementGrpcService(console)
	// create connection grpc client adapter:
	grpcConnectionClient, err := grpcAdapter.GRPCConnectionClientManager()

	if err != nil {
		console.Error("error while initializing grpc connection client", zap.Error(err))
		return
	}

	defer func(grpcConnectionClient *grpc.ClientConn) {
		err := grpcConnectionClient.Close()
		if err != nil {
			console.Error("error while closing grpc connection client", zap.Error(err))
		}
	}(grpcConnectionClient)

	// create instance validate request adapter:
	validationInstance := validation.NewValidatorService(console)

	// create instance http server adapter:
	serverInstance := server.NewServer(console)

	// Configura el manejo de señales.
	stop := setupSignalHandler(console)

	// go routine http server arg port, channel.
	go serverInstance.ListenAndServe(config.Config.Server.Port, stop)

	// create instance client grpc creator:
	grpcClient := hservice.NewGrpcClientCreator()

	// create instance client grpc service:
	helloService := hservice.NewHelloService(console, grpcConnectionClient, grpcClient)

	// create instance grpc command adapter:
	commandAdapter := command.NewGrpcHelloCommand(console, &helloService)

	// create instance rabbit adapter:
	rabbitInstance := queue.NewMqAdapter(console)

	// create instance manager connection rabbit:
	rabbitConnection := rabbitmq.NewManagerConnection(rabbitInstance)

	// create connection rabbit:
	err = rabbitConnection.RabbitMqConnection()

	if err != nil {
		console.Error("error while initializing rabbitmq connection", zap.Error(err))
		return
	}

	// create instance messages received:
	messageInstance := messages.NewMessageAdapter(console, rabbitInstance, commandAdapter, validationInstance)

	go messageInstance.ReceiveMessages(config.Config.Queue.Consumer)

	// create instance ssh adapter:
	sshInstance := ssh.NewShhAdapter(console, rabbitInstance)

	// create instance scheduler adapter:
	taskInstance := task.NewScheduler(sshInstance, console, rabbitInstance)

	// Crear la tarea que se ejecutará cada 30 segundos
	exec := taskInstance.Create()

	console.Info("task is enable ", zap.Bool("enable", config.Config.Scheduler.Enable))
	if config.Config.Scheduler.Enable {
		taskInstance.Run(exec)
	}

	select {}
}

// SetupSignalHandler configura el manejo de señales para una parada controlada.
func setupSignalHandler(console *zap.Logger) (quitOs <-chan struct{}) {
	quit := make(chan struct{})
	// Canal para recibir señales del sistema
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Espera la primera señal y cierra el canal `stop`.
		next := <-s
		console.Info("caught signal next", zap.Any("signal", next))
		close(quit)
		// Espera una segunda señal para terminar inmediatamente.
		next = <-s
		console.Info("caught signal next", zap.Any("signal", next))
		os.Exit(1)
	}()
	return quit
}
