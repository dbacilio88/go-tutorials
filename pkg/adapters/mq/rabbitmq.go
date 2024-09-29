package mq

/**
*
* rabbitmq
* <p>
* rabbitmq file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author bxcode
* @author dbacilio88@outlook.es
* @since 17/08/2024
*
 */

type Executor interface{}
type RabbitMqAdapter struct {
}

func NewRabbitMqAdapter() *RabbitMqAdapter {
	return &RabbitMqAdapter{}
}
