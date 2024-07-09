package server

import (
	"net/http"
)

/**
*
* server
* <p>
* server file
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

func WebServer() {
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./server/view/index.html")
}
