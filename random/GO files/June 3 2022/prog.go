package main

import (
	"fmt"
)

func main() {

	// create a slice with a length of 5 elements
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// you can't access an index of a slice beyond its length.
	fruits[5] = "Runtime Error"

	// error: panic: runtime error: index out of range

	fmt.Println(fruits)
}
