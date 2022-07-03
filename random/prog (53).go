package main

import (
	"fmt"
)

type Number int

func (n *Number) Display() {
	fmt.Println(*n) // Print the address of n.
}

func (n *Number) Double() {
	*n *= 2
}

func main() {
	number := Number(16)
	number.Double()
	number.Display()

}
