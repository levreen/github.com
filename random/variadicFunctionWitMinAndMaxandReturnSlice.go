// Variadic Function that takes min and max and compare it on the arguments in numbers parameter.
package main

import (
	"fmt"
)

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number <= max && number >= min { // This is logic starts
			result = append(result, number)
		}
	}
	return result
}
func main() {
	fmt.Println(inRange(1, 5, -1, 2, 4, 6))
	fmt.Println(inRange(1, 5))
}
