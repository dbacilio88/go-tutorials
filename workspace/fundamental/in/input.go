package in

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
*
* input
* <p>
* input file
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

var num1 int
var num2 int
var legend string
var err error

func Execute() {
	input()
}

func input() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese dato 1 :")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Error to convert datos to int")
		}
	}

	fmt.Println("Ingrese dato 2 :")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Error to convert datos to int")
		}
	}

	fmt.Println("Ingrese legenda: ")
	if scanner.Scan() {
		legend = scanner.Text()
	}
	fmt.Println(legend, num1*num2)

}
