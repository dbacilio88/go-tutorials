package request

/**
*
* message
* <p>
* message file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author bxcode
* @author dbacilio88@outlook.es
* @since 9/11/2024
*
 */

type MessageRequest struct {
	Prefix string `json:"prefix" validate: "require"`
	Name   string `json:"name" validate: "require"`
}

type ValidationData struct {
	Key   string
	Value string
}
