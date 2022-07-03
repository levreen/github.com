package main

import (
	"fmt"
)

func main() {
	// declare variables that are set to their zero value.
	// using var pag wala tayong value na ilalagay - zero value
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c string \t %T [%v]\n", c, c)
	fmt.Printf("var d string \t %T [%v]\n", d, d)

	// declare variables and initialize.
	// using the short variables declaration operator = productivity operator.
	// using the short assignment pag mag-initialize / maglagay na ng values sa mga variables natin
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := 10 \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 10 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := 10 \t %T [%v]\n", dd, dd)

	// specify type and perform a conversion.
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}
