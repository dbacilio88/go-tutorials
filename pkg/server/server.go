package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/dbacilio88/go/pkg/config"
	proto "github.com/dbacilio88/go/proto/hello"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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
* @author bxcode
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
	s.console.Info("Received hello request", zap.String("name", in.GetHello().GetFirstName()))

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

// startGrpcServer Función para iniciar el servidor gRPC con TLS
func (s *Server) startGrpcServer() (*grpc.Server, net.Listener, error) {
	address := fmt.Sprintf("%s:%s", config.Config.GrpcServer.Host, config.Config.GrpcServer.Port)
	s.console.Info("Starting gRPC server on", zap.String("address", address))

	listen, err := net.Listen(config.Config.GrpcServer.Protocol, address)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start gRPC listener: %v", err)
	}

	// Cargar certificados TLS
	cred, err := credentials.NewServerTLSFromFile(config.Config.GrpcServer.Cert, config.Config.GrpcServer.Key)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load TLS credentials: %v", err)
	}

	// Crear el servidor gRPC con credenciales TLS
	grpcSrv := grpc.NewServer(grpc.Creds(cred))

	// Registrar el servicio gRPC
	proto.RegisterHelloServiceServer(grpcSrv, &grpcServer{})

	// Habilitar reflexión
	reflection.Register(grpcSrv)

	return grpcSrv, listen, nil
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

	// Iniciar servidor gRPC
	grpcServer, grpcListener, err := s.startGrpcServer()
	if err != nil {
		s.console.Error("gRPC server startup failed", zap.Error(err))
		return
	}

	// Inicia el servidor gRPC en un puerto diferente
	go func() {
		if err := grpcServer.Serve(grpcListener); err != nil {
			s.console.Error("gRPC server error", zap.Error(err))
		}
	}()

	// Esperar a recibir una señal
	<-quit
	log.Println("shutting down server...")

	// Establece un tiempo límite para la parada del servidor.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// Intenta cerrar el servidor de manera ordenada.
	if err := httpServer.Shutdown(ctx); err != nil {
		s.console.Error("HTTP server shutdown failed", zap.Error(err))
	}
	grpcServer.GracefulStop()
	s.console.Info("Servers shutdown successfully")
}
