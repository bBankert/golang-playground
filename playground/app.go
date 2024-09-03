package main

import (
	"fmt"
	"math"

	"example.com/playground/functions"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	//squaredNumbers := functions.TransformValues(&numbers, square)

	cubedNumbers := functions.TransformValues(&numbers, func(value int) int {
		return int(math.Pow(float64(value), 3))
	})

	fmt.Println(cubedNumbers)
}

func square(value int) int {
	return int(math.Pow(float64(value), 2))
}
