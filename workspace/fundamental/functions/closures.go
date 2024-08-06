package functions

import "fmt"

/**
*
* closures
* <p>
* closures file
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

func Table(value int) func() int {
	num := value
	sequence := 0
	return func() int {
		sequence++
		return num * sequence
	}
}

func CallClosure() {
	tableDel := 2
	myTable := Table(tableDel)
	for i := 0; i < 10; i++ {
		fmt.Println(myTable())
	}
}
