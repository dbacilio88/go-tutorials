package server

import (
	"context"
	"errors"
	"fmt"
	proto "github.com/dbacilio88/go/proto/hello"
	"go.uber.org/zap"
	"net/http"
	"time"
)

/**
*
* server
* <p>
* server file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 4/08/2024
*
 */

type Executor interface {
}

// Server estructura que gestiona los servidores HTTP y gRPC
type Server struct {
	router  *http.ServeMux
	console *zap.Logger
}

// gRPC grpcServer estructura
type grpcServer struct {
	proto.UnimplementedHelloServiceServer
	console *zap.Logger
}

// NewServer crea una nueva instancia del servidor HTTP y gRPC
func NewServer(console *zap.Logger) *Server {
	return &Server{
		router:  http.NewServeMux(),
		console: console,
	}
}

// Hello Implementación del método Hello del servicio gRPC
func (s *grpcServer) Hello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	s.console.Info("Received h-service request", zap.String("name", in.GetHello().GetFirstName()))

	tqr := in.GetHello()
	prefix := tqr.GetPrefix()
	firstName := tqr.GetFirstName()
	message := "Hello " + prefix + ", " + firstName + " welcome"
	response := &proto.HelloResponse{
		CustomHello: message,
	}
	return response, nil
}

// handler para el servidor HTTP
func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, World! :P")
	if err != nil {
		return
	}
}

// startHttpServer Función para iniciar el servidor HTTP
func (s *Server) startHttpServer(addr string) *http.Server {
	http.HandleFunc("/", handler)

	// Configurar y devolver el servidor HTTP
	serv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
	}
	return serv
}

// ListenAndServe Escuchar y servir tanto HTTP como gRPC
func (s *Server) ListenAndServe(addr string, quit <-chan struct{}) {
	// Iniciar servidor HTTP
	httpServer := s.startHttpServer(addr)

	go func() {
		s.console.Info("start http server on", zap.String("addr", addr))
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.console.Error("http server listen error", zap.Error(err))
			return
		}
	}()

	// Esperar a recibir una señal
	<-quit

	s.console.Info("shutting down server...")
	// Establece un tiempo límite para la parada del servidor.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// Intenta cerrar el servidor de manera ordenada.
	if err := httpServer.Shutdown(ctx); err != nil {
		s.console.Error("HTTP server shutdown failed", zap.Error(err))
	}
	//grpcServer.GracefulStop()
	s.console.Info("Servers shutdown successfully")
}
