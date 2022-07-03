package main

import (
	"fmt"
)

func newInt() *int {
	return new(int)
}
func main() {
	p := new(int)                  // p, of type *int, points to an unnamed int variable
	fmt.Println("addr of v\t:", p) // prints the add of p, fmt.Println(*p) - prints its value
	*p = 2                         // sets the unamed int to 2
	fmt.Println("val of v\t:", *p)
	q := new(int)
	fmt.Println(p == q)
	fmt.Println(q)
}


