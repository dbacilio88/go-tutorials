package channels

import (
	"fmt"
	"strings"
	"time"
)

/**
*
* channels
* <p>
* channels file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author cbaciliod
* @author dbacilio88@outlook.es
* @since 9/07/2024
*
 */

func SlowData(text string, channel chan bool) {
	letters := strings.Split(text, "")
	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}

	channel <- true

}
