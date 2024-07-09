package middleware

import "fmt"

/**
*
* middleware
* <p>
* middleware file
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

// es un interceptor [mismo data de entrada que salida]

func Middleware(a, b int) {
	fmt.Println("Start")
	defer fmt.Println("End")
	result := operation(plus)(2, 3)
	fmt.Println(result)
	result = operation(sub)(2, 3)
	fmt.Println(result)
}
func sub(a, b int) int {
	return a - b
}
func plus(a, b int) int {
	return a + b
}

func operation(f func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		fmt.Println("x = ", x, "y = ", y)
		return f(x, y)
	}
}
