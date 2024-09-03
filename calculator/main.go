package main

import (
	"fmt"

	"example.com/file_operations"
	"example.com/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := file_operations.New("prices.txt", fmt.Sprintf("result_%.2f.json", taxRate*100))
		taxedPrice := prices.New(
			[]float64{10, 20, 30},
			taxRate,
			*fm)

		taxedPrice.Process()
	}

}
