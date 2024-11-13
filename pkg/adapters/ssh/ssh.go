package ssh

import (
	"fmt"
	"github.com/dbacilio88/go/pkg/adapters/queue"
	"github.com/dbacilio88/go/pkg/config"
	"go.uber.org/zap"
	client "golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"net"
	"os"
)

/**
*
* ssh
* <p>
* ssh file
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

// El paquete ssh implementa un cliente y un servidor SSH.
// SSH es un protocolo de seguridad de transporte,
// un protocolo de autenticación y una familia de protocolos de aplicación.
// El protocolo de nivel de aplicación más típico es un shell remoto y
// este se implementa específicamente. Sin embargo,
// la naturaleza multiplexada de SSH está expuesta a los usuarios que desean brindar soporte a otros.

type Executor interface {
	Connection() (*client.Client, error)
}

type ShhAdapter struct {
	console       *zap.Logger
	rabbitAdapter queue.Executor
}

func NewShhAdapter(console *zap.Logger, rabbitAdapter queue.Executor) *ShhAdapter {
	return &ShhAdapter{
		console:       console,
		rabbitAdapter: rabbitAdapter,
	}
}

func (s *ShhAdapter) Connection() (*client.Client, error) {
	s.console.Info("connecting to server ssh")

	conn, err := client.Dial(
		config.Config.Ssh.Protocol,
		net.JoinHostPort(config.Config.Ssh.Host, config.Config.Ssh.Port),
		s.clientConfig(&config.Config.Ssh),
	) //Tema familiar urgente

	if err != nil {
		s.failOnError(err, "error connecting to server ssh")
		return nil, err
	}
	s.console.Info("connected to server ssh")
	return conn, nil
}

func (s *ShhAdapter) clientConfig(config *config.Ssh) *client.ClientConfig {
	s.console.Info("load configuration properties ssh")
	s.console.Info("connection ssh for private key", zap.Bool("enable", config.Enable))
	if config.Enable {
		s.console.Info("connecting to server ssh")
		host, err := knownhosts.New(config.KnownHosts)

		if err != nil {
			s.failOnError(err, "could not load known hosts")
			return nil
		}

		key, err := os.ReadFile(config.PrivateKey)

		if err != nil {
			s.failOnError(err, "unable to read private key")
			return nil
		}
		// Create the Signer for this private key.
		private, err := client.ParsePrivateKey(key)
		if err != nil {
			s.failOnError(err, "unable to parse private key")
			return nil
		}
		return &client.ClientConfig{
			User: config.Username,
			Auth: []client.AuthMethod{
				client.PublicKeys(private),
			},
			HostKeyCallback: host,
			Timeout:         0,
		}
	} else {
		return &client.ClientConfig{
			User: config.Username,
			Auth: []client.AuthMethod{
				client.Password(config.Password),
			},
			HostKeyCallback: client.InsecureIgnoreHostKey(),
			Timeout:         0,
		}
	}
}

func (a *ShhAdapter) failOnError(err error, msg string) {
	if err != nil {
		fullMessage := fmt.Sprintf("%s: %s", msg, err.Error())
		a.console.Error(fullMessage, zap.Error(err))
	}
}
