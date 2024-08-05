package server

import (
	"context"
	"errors"
	"fmt"
	"log"
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
type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		router: http.NewServeMux(),
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, World!")
	if err != nil {
		return
	}
}
func (s *Server) ListenAndServe(addr string, quit <-chan struct{}) {
	http.HandleFunc("/", handler)
	// Configura el servidor HTTP.
	serv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
	}
	go func() {
		log.Println("start http server on", serv.Addr)
		if err := serv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server listen error: %v", err)
		}
	}()
	// Esperar a recibir una señal
	<-quit
	log.Println("shutting down server...")

	// Establece un tiempo límite para la parada del servidor.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// Intenta cerrar el servidor de manera ordenada.
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatalf("http server shutdown error: %v", err)
	}
	log.Println("server shutdown successfully")
}
