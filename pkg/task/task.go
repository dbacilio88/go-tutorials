package task

import (
	"github.com/dbacilio88/go/pkg/adapters/queue"
	"github.com/dbacilio88/go/pkg/adapters/sftp"
	"github.com/dbacilio88/go/pkg/adapters/ssh"
	"github.com/madflojo/tasks"
	"go.uber.org/zap"
	client "golang.org/x/crypto/ssh"
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
	ssh           ssh.Executor
	console       *zap.Logger
	rabbitAdapter queue.Executor
}

func NewScheduler(ssh ssh.Executor, console *zap.Logger, rabbitAdapter queue.Executor) *Scheduler {
	return &Scheduler{
		ssh:           ssh,
		console:       console,
		rabbitAdapter: rabbitAdapter,
	}
}
func (s *Scheduler) Create() *tasks.Scheduler {
	s.console.Info("create task")
	return tasks.New()
}
func (s *Scheduler) Run(exec *tasks.Scheduler) {
	s.console.Info("run task")
	task := &tasks.Task{
		Interval:          10 * time.Minute,
		RunOnce:           false,
		RunSingleInstance: false,
		TaskFunc: func() error {
			s.console.Info("run task for 1 minutes")
			con, err := s.ssh.Connection()
			if err != nil {
				s.console.Fatal("ssh connection error", zap.Error(err))
				return err
			}

			instance := sftp.NewFtp(s.console, s.rabbitAdapter)
			_, err = instance.Connection(con)

			if err != nil {
				s.console.Fatal("sftp connection error", zap.Error(err))
				return err
			}
			defer func(con *client.Client) {
				_ = con.Close()
				s.console.Info("ssh connection closed")
			}(con)

			return nil
		},
		ErrFunc: func(err error) {
			s.console.Fatal("task err", zap.Error(err))
		},
	}

	add, err := exec.Add(task)
	if err != nil {
		s.console.Fatal("add task fail", zap.Error(err))
		return
	} else {
		s.console.Info("add task success", zap.String("add", add))
	}
}
