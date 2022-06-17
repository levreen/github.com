package main

import (
	"fmt"
)

func main() {
	var total float64
	weeks := [3]float64{71.8, 56.2, 89.5}
	for _, val := range weeks {
		total += val
	}
	fmt.Println("average: ", total/float64(len(weeks)))
}
