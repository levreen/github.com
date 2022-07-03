// Variadic Function with two parameters in a function declaration.
package main

import (
	"fmt"
)

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}
func main() {
	mix(1, true, "a", "b")
	mix(2, false)
}
