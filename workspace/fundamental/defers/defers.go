package defers

import (
	"fmt"
	"log"
)

/**
*
* defers
* <p>
* defers file
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

func Defers() {

	fmt.Println("first messages")
	defer fmt.Println("final messages")
	fmt.Println("second messages")
	DemoPanic()
}

func DemoPanic() {

	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurred error %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("search value 1")
	}
}
