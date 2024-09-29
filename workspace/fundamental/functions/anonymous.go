package functions

import "fmt"

/**
*
* anonymous
* <p>
* anonymous file
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

func Calculate() {
	suma := func(num1 int, num2 int) int {
		return num1 + num2
	}
	fmt.Println("sum ", suma(1, 2))
}
