package main

import (
	"fmt"
)

type T struct {
	f int
}

func main() {
	v1 := [2][2]T{{{1}, {2}}, {{3}, {4}}}
	fmt.Printf("%#v\n", v1)
	v2 := [][]int{{1, 2}, {3}}
	fmt.Printf("%#v\n", v2)
}

// multi-dimensional elements