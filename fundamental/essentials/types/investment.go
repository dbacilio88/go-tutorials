package main

import (
	"fmt"
	"math"
)

/*
*
@note: https://www.calculator.net/future-value-calculator.html
*/
func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64 = 10
	investmentAmount = 2000

	scan, err := fmt.Scan(&investmentAmount)
	fmt.Println(scan)
	if err != nil {
		return
	}

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
