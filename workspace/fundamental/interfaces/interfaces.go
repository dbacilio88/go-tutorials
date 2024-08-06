package interfaces

/**
*
* interfaces
* <p>
* interfaces file
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

type Human interface {
	Running()
	Eating()
}

type Animal interface {
	Eat()
}

type Vegetal interface {
	Classification()
}
