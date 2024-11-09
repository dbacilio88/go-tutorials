package functions

import (
	"fmt"
	"strconv"
)

/**
*
* utils
* <p>
* utils file
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
	status, result := convertToText(0)
	fmt.Println("status = ", status, "result = ", result)

	num, text := convertToNumber("result")
	fmt.Println("num = ", num, "text = ", text)

}

// convertToText convert of int to string
func convertToText(num int) (bool, string) {
	result := strconv.Itoa(num)
	return true, result
}

// convertToNumber convert of string to int
func convertToNumber(text string) (int, string) {

	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, "hubo un error, messages:spring " + err.Error()
	}

	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}
}
