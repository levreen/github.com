package main

import (
	"fmt"
)

func f() *int {
	v := 1
	// local address of f()
	return &v // return loc ng "v" not val of "v"
}
func main() {
	var p = f()
	fmt.Println(p)
	fmt.Println(f() == f()) // false
}
