// Print out the maximum number among the arguments in the slice of numbers
// Variadic function.
package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max { // This is where logic starts
			max = number // When the float is bigger than -1. It will become the max. So on.
		}
	}
	fmt.Println(max)
}
func main() {
	maximum(78.5, 66.7, 99.55, 105.2)
}
