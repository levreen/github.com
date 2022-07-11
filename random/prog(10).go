package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 29; i++ {
		fmt.Printf("%3d: %04o\n", i, i)
	}
}
