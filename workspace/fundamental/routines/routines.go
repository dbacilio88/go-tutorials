package routines

import (
	"fmt"
	"strings"
	"time"
)

/**
*
* routines
* <p>
* routines file
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

func SlowNames(name string) {
	letters := strings.Split(name, "")
	fmt.Println(letters)

	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("letter ", letter)
	}
}
