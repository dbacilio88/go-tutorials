package model

import "time"

/**
*
* user
* <p>
* user file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author cbaciliod
* @author dbacilio88@outlook.es
* @since 8/07/2024
*
 */

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	status    bool
}

func (u *User) Add(id int, name string, status bool) {
	u.Id = id
	u.Name = name
	u.CreatedAt = time.Now()
	u.status = status
}
