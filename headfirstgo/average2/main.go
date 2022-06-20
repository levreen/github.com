// average2 calculate the average of several numbers.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(num ...float64) { // This is float64 using Variadic Function *NOT* slice.
	var sums float64
	for _, val := range num {
		sums += val
	}
	length := float64(len(num))
	fmt.Printf("Average: %0.2f", sums/length)
}

func main() { // This uses slice of float64. But using ellipsis ... when passing sum as an argument works to the function average.
	arguments := os.Args[1:]
	var sum []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum = append(sum, number)

	}
	average(sum...)
}
