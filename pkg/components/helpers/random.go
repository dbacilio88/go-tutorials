package helpers

import (
	"math/rand"
	"strings"
	"time"
)

/**
*
* random
* <p>
* random file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 5/08/2024
*
 */

const ALPHABET = "abcdefghijklmnopqrstuvwxyz"

func init() {
	// Generar un slice de bytes aleatorio
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(ALPHABET)
	for i := 0; i < n; i++ {
		c := ALPHABET[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomUser generates a random username
func RandomUser() string {
	return RandomString(6)
}

// RandomPassword generates a random password
func RandomPassword() string {
	return RandomString(10)
}

// RandomRole generates a random role
func RandomRole() string {
	roles := []string{"ADMIN", "USER"}
	n := len(roles)
	return roles[rand.Intn(n)]
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currency := []string{"USD", "EUR", "GBP", "JPY"}
	n := len(currency)
	return currency[rand.Intn(n)]
}
