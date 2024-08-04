package app

import (
	"fmt"
	"github.com/dbacilio88/go/interfaces"
	"github.com/dbacilio88/go/structs/model"
)

/**
*
* crud
* <p>
* crud file
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

func AddUser() {
	u := new(model.User)
	u.Add(1, "David", true)
	fmt.Println("Add User Success ", u)
}

func ExecuteInterfaces() {
	David := new(model.Men)
	MenIsEating(David)

	Maria := new(model.Woman)
	MenIsEating(Maria)

}

func MenIsEating(m interfaces.Human) {
	m.Eating()
	fmt.Println("Eating Success ", m)
}
