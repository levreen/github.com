package main

import (
	"fmt"
)

type Number int

func (n *Number) Double() { // Change the receiver parameter to a pointer type.
	*n *= 2 // Update the value at the pointer. 
}

func main() {
	number := Number(8)
	fmt.Println("Original value of number: ", number)
	number.Double()
	fmt.Println("Original value of number: ", number)
}
