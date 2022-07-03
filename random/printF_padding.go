package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")

	fmt.Println("--------------------")

	fmt.Printf("%11s | %5d\n", "Stamps", 50)
	fmt.Printf("%11s | %5d\n", "Paper Clips", 5)
	fmt.Printf("%11s | %5d\n", "Tape", 5000)

	fmt.Println("--------------------")
}
