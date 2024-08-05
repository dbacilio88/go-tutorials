package task

import (
	"github.com/dbacilio88/go/pkg/adapters/ssh"
	"github.com/dbacilio88/go/pkg/adapters/ssh/ftp"
	"github.com/madflojo/tasks"
	client "golang.org/x/crypto/ssh"
	"log"
	"time"
)

/**
*
* task
* <p>
* task file
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
	Create() *tasks.Scheduler
	Run(exec *tasks.Scheduler)
}

type Scheduler struct {
	ssh ssh.Executor
}

func NewScheduler(ssh ssh.Executor) *Scheduler {
	return &Scheduler{
		ssh: ssh,
	}
}
func (s *Scheduler) Create() *tasks.Scheduler {
	log.Println("create task")
	return tasks.New()
}
func (s *Scheduler) Run(exec *tasks.Scheduler) {
	log.Println("run task")
	task := &tasks.Task{
		Interval:          10 * time.Second,
		RunOnce:           false,
		RunSingleInstance: false,
		TaskFunc: func() error {
			log.Println("run task for 1 minute")
			con, err := s.ssh.Connection()
			if err != nil {
				log.Fatalf("ssh connection error: %s", err)
				return err
			}

			instance := ftp.NewFtp()
			_, err = instance.Connection(con)

			if err != nil {
				log.Fatalf("sftp connection error: %s", err)
				return err
			}
			defer func(con *client.Client) {
				_ = con.Close()
				log.Println("ssh connection closed")
			}(con)

			return nil
		},
		ErrFunc: func(err error) {
			log.Fatalf("task err: %v", err)
		},
	}

	add, err := exec.Add(task)
	if err != nil {
		log.Fatalf("add task fail: %s", err)
		return
	} else {
		log.Println("add task success", add)
	}
}
