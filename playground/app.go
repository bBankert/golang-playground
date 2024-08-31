package main

import (
	"example.com/playground/custom_types"
)

func main() {
	var name custom_types.Str = "Testing"

	name.Log()
}
