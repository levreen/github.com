package main

import (
	"fmt"
)

func main() {
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2
	for i, v := range numbers {
		fmt.Println(i, v)
	}
	var letters = []string{"a", "b", "c"}
	for i, v := range letters {
		fmt.Println(i, v)
	}
}
