package main

import (
	"fmt"
)

func average(num ...float64) {
	var sums float64
	for _, val := range num {
		sums += val
	}
	length := float64(len(num))
	fmt.Println(sums / length)
}

func main() {
	average(100, 50)
	average(90.7, 89.7, 98.5, 92.3)
}
