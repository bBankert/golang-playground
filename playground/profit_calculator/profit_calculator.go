package profit_calculator

import (
	"fmt"
	"os"

	lib "example.com/playground/lib"
)

func CalculateProfit() {
	var revenue, expenses, taxRate float64

	lib.GetUserInput("Input revenue:", &revenue)
	lib.GetUserInput("Input expenses: ", &expenses)
	lib.GetUserInput("Input tax rate: ", &taxRate)

	valid, errorMessage := validateInputs(revenue, expenses, taxRate)

	if !valid {
		panic(fmt.Sprintf("*** Error: %s *****", errorMessage))
	}

	earningsBeforeTax, profit, ratio := calculateFinancials(
		revenue,
		expenses,
		taxRate)

	fmt.Printf("Earnings before tax: %.0f\n Profit: %.0f\n Ratio: %.0f",
		earningsBeforeTax,
		profit,
		ratio)

	outputData := fmt.Sprintf("Earnings before Tax: %.2f\n Profit: %.2f\n Ratio: %.2f",
		earningsBeforeTax,
		profit,
		ratio)

	os.WriteFile("profit.txt", []byte(outputData), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}

func validateInputs(revenue, expenses, taxRate float64) (valid bool, errorMessage string) {

	valid = true

	if revenue <= 0 {
		valid = false
		errorMessage = "Invalid revenue amount, must be greater than 0"
	}

	fmt.Printf("Testing %#v\n", valid)
	if expenses <= 0 && valid {
		valid = false
		errorMessage = "Invalid expenses amount, must be greater than 0"
	}

	if taxRate <= 0 && valid {
		valid = false
		errorMessage = "Invalid taxRate amount, must be greater than 0"
	}

	return valid, errorMessage
}
