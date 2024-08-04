package arrays

import (
	"fmt"
	"math/rand"
)

/**
*
* arrays
* <p>
* arrays file
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
// slice is []int (not size vector)
var table [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func Arrays() {

	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
}

func GenerareArray(num int) []int {
	var array []int
	for i := 0; i < num; i++ {
		value := rand.Intn(10)
		array = append(array, value)
	}
	return array
}

func Mod2(num int) []int {
	var array []int
	if num%2 == 0 {
		array = append(array, num)
	}
	return array
}
