package variables

import "fmt"

/**
*
* numeric
* <p>
* numeric file
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

func callMethods() {
	Boolean()
	Integer()
	Float()
}

func Boolean() {
	var boolean1 bool
	boolean1 = true
	fmt.Println("boolean1 = ", boolean1)
}

func Float() {
	var float1 float32
	var float2 float64
	float1 = 12.2
	float2 = 1234.2
	fmt.Println("float1 = ", float1)
	fmt.Println("float2 = ", float2)
}

func Integer() {
	var num1 int
	var num2 int
	var num3 int8
	var num4 int16
	var num5 int64
	var num6 uint
	var num7 uint8
	var num8 uint16
	var num9 uint32
	var num10 uint64

	num1 = 10
	num2 = 32
	num3 = 127
	num4 = 12349
	num5 = 12349555
	num6 = 1234
	num7 = 255
	num8 = 65535
	num9 = 4294967295
	num10 = 18446744073709551615
	fmt.Println("num1  = ", num1)
	fmt.Println("num2  = ", num2)
	fmt.Println("num3  = ", num3)
	fmt.Println("num4  = ", num4)
	fmt.Println("num5  = ", num5)
	fmt.Println("num6  = ", num6)
	fmt.Println("num7  = ", num7)
	fmt.Println("num8  = ", num8)
	fmt.Println("num9  = ", num9)
	fmt.Println("num10 = ", num10)
}
