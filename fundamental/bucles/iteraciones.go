package bucles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
*
* iteraciones
* <p>
* iteraciones file
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

func Execute() {

	fmt.Println(Iteraciones())
}

var num int
var err error

func Iteraciones() string {
	var text string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("ingresar un número: ")

	if scanner.Scan() {
		num, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("error: %s", err)
		}
	}

	for i := 1; i <= num; i++ {
		text += fmt.Sprintf("%d x %d = %d \n", num, i, i*num)
	}

	return text
}
