package model

/**
*
* men
* <p>
* men file
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

type Men struct {
	Edad   int
	Run    bool
	Eat    bool
	Gender string
}

type Woman struct {
	Men
}

func (m *Men) Running() {
	m.Run = true
}

func (m *Men) Eating() {
	m.Run = false
}
func (m *Men) Genders() string {
	return "Male"
}

func (m *Woman) Running() {
	m.Run = true
}

func (m *Woman) Eating() {
	m.Run = false
}
func (m *Woman) Genders() string {
	return "Female"
}
