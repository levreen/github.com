package main

import (
	. "fmt"
)

// new returns a pseudorandom number generator Rand with a given seed.
// every time you call Rand, you get a new "random" number.

func New(seed int) (Rand func() int) {
	current := seed
	return func() int {
		next := (17 * current) % 97
		current = next
		return next
	}
}
func main() {
	rand1 := New(1)
	Println(rand1(), rand1(), rand1())

	rand1 = New(2)
	Println(rand1(), rand1(), rand1())
}
