package ssh

import (
	"github.com/dbacilio88/go/pkg/config"
	"go.uber.org/zap"
	client "golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"log"
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
	console *zap.Logger
}

func NewShhAdapter(console *zap.Logger) *ShhAdapter {
	return &ShhAdapter{
		console: console,
	}
}

func (s *ShhAdapter) Connection() (*client.Client, error) {
	s.console.Info("connecting to server ssh")

	//log.Println(net.JoinHostPort(config.Config.Ssh.Host, config.Config.Ssh.Port))

	conn, err := client.Dial(
		config.Config.Ssh.Protocol,
		net.JoinHostPort(config.Config.Ssh.Host, config.Config.Ssh.Port),
		s.clientConfig(&config.Config.Ssh),
	) //Tema familiar urgente

	if err != nil {
		log.Fatalf("error connecting to server ssh: %v", err)
		return nil, err
	}
	log.Println("connected to server ssh")
	return conn, nil
}

func (s *ShhAdapter) clientConfig(config *config.Ssh) *client.ClientConfig {
	log.Println("load configuration properties ssh")
	log.Println("connection ssh for private key enable: ", config.Enable)
	if config.Enable {
		log.Println("connecting to server ssh ", config.KnownHosts)
		host, err := knownhosts.New(config.KnownHosts)

		if err != nil {
			log.Fatalf("could not load known hosts: %s", err)
			return nil
		}

		key, err := os.ReadFile(config.PrivateKey)

		if err != nil {
			log.Fatalf("unable to read private key: %v", err)
			return nil
		}
		// Create the Signer for this private key.
		private, err := client.ParsePrivateKey(key)
		if err != nil {
			log.Fatalf("unable to parse private key: %v", err)
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
