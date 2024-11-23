package rabbitmq

import (
	"github.com/dbacilio88/go/pkg/adapters/queue"
)

/**
*
* rabbit
* <p>
* rabbit file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 7/11/2024
*
 */

type Executor interface {
	RabbitMqConnection() error
}
type ManagerConnection struct {
	rabbitmqAdapter queue.Executor
}

// NewManagerConnection create new instance ManageConnection
func NewManagerConnection(rabbitmqAdapter queue.Executor) *ManagerConnection {
	return &ManagerConnection{
		rabbitmqAdapter: rabbitmqAdapter,
	}
}

func (r *ManagerConnection) RabbitMqConnection() error {
	return r.rabbitmqAdapter.Connection()
}
