package slices

import "fmt"

/**
*
* slice
* <p>
* slice file
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

var table []int = []int{1, 2, 3}

func Slice() {
	fmt.Println(table)
	t := table[:]
	fmt.Println(t)
	t = table[1:]
	fmt.Println(t)
	t = table[2:]
	fmt.Println(t)
	t = table[:2]
	fmt.Println(t)
	t = table[:3]
	fmt.Println(t)
}

// Capacity sí se puede modificar la capacidad de lo slice
func Capacity() {
	elements := make([]int, 5, 20)
	fmt.Println("size:", len(elements), "capacity:", cap(elements))

	empty := make([]int, 0)
	fmt.Println("size:", len(empty), "capacity:", cap(empty))

	for i := 0; i < 100; i++ {
		empty = append(empty, i)
	}
	fmt.Println("size:", len(empty), "capacity:", cap(empty))
}
