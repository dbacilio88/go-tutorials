package conditional

import (
	"fmt"
	"runtime"
)

/**
*
* operative
* <p>
* operative file
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

	if os := runtime.GOOS; os == "linux" {
		fmt.Println("Esto es linux")
	} else {
		fmt.Println("Esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux - switch")
	case "windows":
		fmt.Println("Esto es windows - switch")
	default:
		fmt.Printf("Esto es %s switch", os)
	}
}
