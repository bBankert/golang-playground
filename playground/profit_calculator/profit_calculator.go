package profit_calculator

import (
	"fmt"

	lib "example.com/playground/lib"
)

func CalculateProfit() {
	var revenue, expenses, taxRate float64

	lib.GetUserInput("Input revenue:", &revenue)
	lib.GetUserInput("Input expenses: ", &expenses)
	lib.GetUserInput("Input tax rate: ", &taxRate)

	earningsBeforeTax := revenue - expenses

	profit := earningsBeforeTax - (1 - taxRate/100)

	ratio := earningsBeforeTax / profit

	fmt.Println(earningsBeforeTax)
	fmt.Println(profit)
	fmt.Println(ratio)
}
