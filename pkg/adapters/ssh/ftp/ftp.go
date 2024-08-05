package ftp

import (
	"github.com/google/uuid"
	"github.com/pkg/sftp"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"log"
	"path"
)

/**
*
* ftp
* <p>
* ftp file
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
	Connection(con *ssh.Client) (*sftp.Client, error)
	GetRootPath(conn *sftp.Client) error
}

type Ftp struct {
}

func NewFtp() *Ftp {
	return &Ftp{}
}

func (s *Ftp) Connection(con *ssh.Client) (*sftp.Client, error) {
	log.Println("Connecting to sftp server...")
	session, err := sftp.NewClient(con)
	if err != nil {
		log.Fatalf("Failed to sftp connection: %s", err)
		return nil, err
	}

	if err := s.GetRootPath(session); err != nil {
		log.Fatalf("Failed to get root path: %s", err)
	}

	if err = s.GetFiles(session); err != nil {
		log.Fatalf("Failed to get files: %s", err)
	}

	if err := s.CreateDir(session, uuid.New().String()); err != nil {
		log.Fatalf("Failed to create data directory: %s", err)
	}
	if err = s.GetFiles(session); err != nil {
		log.Fatalf("Failed to get files: %s", err)
	}

	defer func(session *sftp.Client) {
		_ = session.Close()
		log.Println("Session sftp closed")
	}(session)

	log.Println("Connected to sftp server")
	return session, nil
}

func (s *Ftp) GetRootPath(conn *sftp.Client) error {
	log.Println("Getting root path...")
	wd, err := conn.Getwd()
	if err != nil {
		return err
	}
	viper.SetDefault("sftp_path", wd)
	return nil
}

func (s *Ftp) GetFiles(conn *sftp.Client) error {

	log.Println("Getting files from sftp server...")

	dir, err := conn.ReadDir(viper.GetString("sftp_path"))
	if err != nil {
		log.Fatalf("Failed to get files from sftp server: %s", err)
		return err
	}
	for _, file := range dir {
		if file.IsDir() {
			log.Println("dir:", file.Name())
		} else {
			log.Println("file:", file.Name())
		}
	}
	return nil
}
func (s *Ftp) CreateFile(conn *sftp.Client) error {

	log.Println("Create files from sftp server...")

	dir, err := conn.ReadDir(viper.GetString("sftp_path"))
	if err != nil {
		log.Fatalf("Failed to get files from sftp server: %s", err)
		return err
	}
	for _, file := range dir {
		if file.IsDir() {
			log.Println("dir:", file.Name())
		} else {
			log.Println("file:", file.Name())
		}
	}
	return nil
}
func (s *Ftp) CreateDir(conn *sftp.Client, root string) error {

	log.Println("Create dir from sftp server...")
	p := path.Join(viper.GetString("sftp_path"), root)
	err := conn.Mkdir(p)
	if err != nil {
		log.Fatalf("Failed to create dir: %s", err)
		return err
	}
	return nil
}
