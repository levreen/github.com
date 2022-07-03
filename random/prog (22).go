// go identity
package main

import (
	"fmt"
)

// BinaryOp type int returns in
type BinaryOp func(int, int) int // bago sakin to

// Do func has fun that takes BinaryOp, a, and b; type(value semantic) return int
func Do(fun BinaryOp, a, b int) int {
	return fun(a, b)
}

func main() {
	result := Do(func(a, b int) int { return a + b },
		1, 2)
	fmt.Println(result)
}

/*
if BinaryOp (a named type) and func(int, int) int (an unnamed type)
weren't identical, writing anonymous functions would be quite a bit ugly.
*/