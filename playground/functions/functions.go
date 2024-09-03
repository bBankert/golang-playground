package functions

func TransformValues(values *[]int, transform func(int) int) []int {
	result := []int{}

	for _, value := range *values {
		result = append(result, transform(value))
	}

	return result
}
