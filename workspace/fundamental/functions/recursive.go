package functions

import "fmt"

/**
*
* recursive
* <p>
* recursive file
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

func Exponent(value int) {
	if value > 10 {
		return
	}
	fmt.Println(value)
	Exponent(value * 2)
}
