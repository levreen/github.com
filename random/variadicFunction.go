// Variadic function that takes any number of arguments when function calls
package main

import (
	"fmt"
)

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}
func main() {
	severalInts(1)
	severalInts(1, 2, 3)
}
