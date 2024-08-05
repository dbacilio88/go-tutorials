package pointers

import "fmt"

/**
*
* pointer
* <p>
* pointer file
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

func Pointers() {

	// Declaración de punteros:
	var ptrA *int // en esta línea declaramos un puntero de tipo int
	fmt.Println("ptrA:", ptrA)

	// Derreferenciación de punteros:
	var x int = 10
	var ptrB = &x // ptrB ahora apunta a la dirección de memoria de x
	fmt.Println("x:", x)
	fmt.Println("ptrB:", ptrB) // Imprime el valor de 'x' mediante la derreferenciación de ptrB
	fmt.Println("*ptrB:", *ptrB)

}
