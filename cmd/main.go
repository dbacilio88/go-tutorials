package main

import (
	"github.com/dbacilio88/go/pkg/adapters/ssh"
	"github.com/dbacilio88/go/pkg/adapters/ssh/config"
	"github.com/dbacilio88/go/pkg/adapters/ssh/server"
	"github.com/dbacilio88/go/pkg/adapters/ssh/task"
	"log"
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
	config.Load("workspace/")
	// create instance ssh adapter:
	_ssh := ssh.NewShhAdapter()
	_server := server.NewServer()
	_task := task.NewScheduler(_ssh)

	// Configura el manejo de señales.
	stop := setupSignalHandler()
	go _server.ListenAndServe(config.Config.Server.Port, stop)

	// Crear la tarea que se ejecutará cada 30 segundos
	exec := _task.Create()

	log.Println("task is enable: ", config.Config.Scheduler.Enable)
	if config.Config.Scheduler.Enable {
		_task.Run(exec)
	}

	select {}
}

// SetupSignalHandler configura el manejo de señales para una parada controlada.
func setupSignalHandler() (quitOs <-chan struct{}) {
	quit := make(chan struct{})
	// Canal para recibir señales del sistema
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Espera la primera señal y cierra el canal `stop`.
		next := <-s
		log.Println("caught signal", next)
		close(quit)
		// Espera una segunda señal para terminar inmediatamente.
		next = <-s
		log.Println("caught signal", next)
		os.Exit(1)
	}()
	return quit
}
