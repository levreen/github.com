package main

import (
	"fmt"
)

var a = "a"

func main() {
	a = "a"
	b := "b"
	if true {
		c := "c"
		if true {
			d := "d"
			fmt.Println(a) // This
			fmt.Println(b) // This
			fmt.Println(c) // It was wrong. It can call on this var within this block. 
			fmt.Println(d) // This
		}
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		//fmt.Println(d) // Not. Correct. It is not declared. The idea of building on top of the root manifests here. 

	}
	fmt.Println(a)
	fmt.Println(b)
	//fmt.Println(c) // Surely enough. This won't work.
	//fmt.Println(d) // Surely enough. This won't work. 
}
