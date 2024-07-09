package maps

import "fmt"

/**
*
* maps
* <p>
* maps file
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

func Maps() {
	counties := make(map[string]string, 4)

	fmt.Println("counties = ", counties)
	fmt.Println("len ", len(counties))

	counties["Perú"] = "Lima"
	fmt.Println(counties)
	fmt.Println(counties["Perú"])

	person := map[string]int{
		"David":     20,
		"Christian": 30,
	}

	fmt.Println(person)

	for name, age := range person {
		fmt.Println("name:", name, "age:", age)
	}

	delete(person, "David")
	fmt.Println(person)

	age, ok := person["David"]
	fmt.Println(age, ok)
}
