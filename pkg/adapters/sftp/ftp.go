package sftp

import (
	"encoding/json"
	"fmt"
	"github.com/dbacilio88/go/pkg/adapters/queue"
	"github.com/dbacilio88/go/pkg/config"
	"github.com/google/uuid"
	"github.com/pkg/sftp"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"path"
)

/**
*
* sftp
* <p>
* sftp file
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
	Connection(con *ssh.Client) (*sftp.Client, error)
	GetRootPath(conn *sftp.Client) error
}

type FtpAdapter struct {
	console       *zap.Logger
	rabbitAdapter queue.Executor
}

func NewFtp(console *zap.Logger, rabbitAdapter queue.Executor) *FtpAdapter {
	return &FtpAdapter{
		console:       console,
		rabbitAdapter: rabbitAdapter,
	}
}

func (s *FtpAdapter) Connection(con *ssh.Client) (*sftp.Client, error) {
	s.console.Info("connecting to sftp server...")
	session, err := sftp.NewClient(con)
	if err != nil {
		s.failOnError(err, "failed to sftp connection")
		return nil, err
	}

	if err := s.GetRootPath(session); err != nil {
		s.failOnError(err, "failed to get root path")
		return nil, err
	}

	if err = s.GetFiles(session); err != nil {
		s.failOnError(err, "failed to get files")
		return nil, err
	}

	nameDir := uuid.New().String()

	if err := s.CreateDir(session, nameDir); err != nil {
		s.failOnError(err, "failed to create data directory")
		return nil, err
	}

	data, err := json.Marshal(nameDir)
	if err != nil {
		s.failOnError(err, "failed to marshal data directory")
		return nil, err
	}

	err = s.rabbitAdapter.SendMessage(config.Config.Queue.Producer, data)
	if err != nil {
		s.failOnError(err, "failed to send messages")
		return nil, err
	}

	if err = s.GetFiles(session); err != nil {
		s.failOnError(err, "failed to get files")
	}

	defer func(session *sftp.Client) {
		_ = session.Close()
		s.console.Info("session sftp closed")
	}(session)

	s.console.Info("connected to sftp server")
	return session, nil
}

func (s *FtpAdapter) GetRootPath(conn *sftp.Client) error {
	wd, err := conn.Getwd()
	if err != nil {
		return err
	}
	viper.SetDefault("sftp_path", wd)
	return nil
}

func (s *FtpAdapter) GetFiles(conn *sftp.Client) error {
	dir, err := conn.ReadDir(viper.GetString("sftp_path"))
	if err != nil {
		s.failOnError(err, "Failed to get files from sftp server")
		return err
	}
	for _, file := range dir {
		if file.IsDir() {
			s.console.Info("name directory", zap.String("name", file.Name()))
		} else {
			s.console.Info("name file", zap.String("name", file.Name()))
		}
	}
	return nil
}
func (s *FtpAdapter) CreateFile(conn *sftp.Client) error {
	dir, err := conn.ReadDir(viper.GetString("sftp_path"))
	if err != nil {
		s.failOnError(err, "Failed to get files from sftp server")
		return err
	}
	for _, file := range dir {
		if file.IsDir() {
			s.console.Info("name directory", zap.String("name", file.Name()))
		} else {
			s.console.Info("name file", zap.String("name", file.Name()))
		}
	}
	return nil
}
func (s *FtpAdapter) CreateDir(conn *sftp.Client, root string) error {
	p := path.Join(viper.GetString("sftp_path"), root)
	err := conn.Mkdir(p)
	if err != nil {
		s.failOnError(err, "failed to create dir")
		return err
	}
	return nil
}

func (s *FtpAdapter) failOnError(err error, msg string) {
	if err != nil {
		fullMessage := fmt.Sprintf("%s: %s", msg, err.Error())
		s.console.Error(fullMessage, zap.Error(err))
	}
}
