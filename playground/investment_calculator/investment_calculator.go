package investment_calculator

import (
	"fmt"
	"math"

	lib "example.com/playground/lib"
)

func CalculateInvestment() {
	const inflationRate = 2.5

	var investmentAmount, expectedReturnRate float64
	var years int

	lib.GetUserInput("Input initial investment amount:", &investmentAmount)
	lib.GetUserInput("Input expected return rate: ", &expectedReturnRate)
	lib.GetUserInput("Input years: ", &years)

	futureValue, futureRealValue := calculateFutureValues(
		investmentAmount,
		expectedReturnRate,
		inflationRate,
		years)

	fmt.Printf("Future value: %.0f\n Future Value (adjusted for Inflaction): %.0f",
		futureValue,
		futureRealValue)

}

func calculateFutureValues(
	investmentAmount, expectedReturnRate, inflationRate float64,
	years int) (futureValue float64, futureRealValue float64) {

	futureValue = float64(investmentAmount) *
		math.Pow(1+expectedReturnRate/100, float64(years))

	futureRealValue = futureValue / math.Pow(1+inflationRate/100, float64(years))

	return futureValue, futureRealValue

}
