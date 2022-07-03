package main

import (
	"fmt"
)

func main() {

	x := 1
	p := &x                                         // p, of type *int, points to x; $x = address of x; now p has the address of x
	fmt.Println("address of x:\t", p)               // prints address of x
	fmt.Println("value of x thru p using *:\t", *p) // prints the value of x which it has address on
	*p = 2                                          // equivalent to x = 2
	fmt.Println("assigning new value of x thru *p:\t", x)
	*p = 1
	fmt.Println("assigning new value of x thru *p:\t", x)

}
