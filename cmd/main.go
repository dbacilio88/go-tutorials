package main

import (
	"github.com/dbacilio88/go/pkg/adapters/ssh"
	"github.com/dbacilio88/go/pkg/config"
	"github.com/dbacilio88/go/pkg/config/logger"
	"github.com/dbacilio88/go/pkg/server"
	"github.com/dbacilio88/go/pkg/task"
	"go.uber.org/zap"
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

	// create instance http server adapter:
	_server := server.NewServer(console)
	// Configura el manejo de señales.
	stop := setupSignalHandler(console)

	// go routine http server arg port, channel.
	go _server.ListenAndServe(config.Config.Server.Port, stop)

	// Usamos un WaitGroup para esperar que todos los servicios terminen

	// create instance ssh adapter:
	_ssh := ssh.NewShhAdapter(console)
	// create instance scheduler adapter:
	_task := task.NewScheduler(_ssh, console)
	// create instance service grpcs:
	//grpcAdapterInstance := grpcs.NewManagementGrpcService(console)

	// create instance service client grpcs:
	//grpcClient, err := grpcAdapterInstance.GRPCConnectionClientManager(&wg)

	//helloCreatorInstance := clients.NewHelloCreator()

	//helloServiceClientInstance := clients.NewHelloServiceClient(grpcClient, helloCreatorInstance)

	//err = helloServiceClientInstance.Hello()

	/*
		defer func(connection *grpc.ClientConn) {
			_ = connection.Close()
		}(grpcClient)

	*/
	// Crear la tarea que se ejecutará cada 30 segundos
	exec := _task.Create()

	console.Info("task is enable ", zap.Bool("enable", config.Config.Scheduler.Enable))
	if config.Config.Scheduler.Enable {
		_task.Run(exec)
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
