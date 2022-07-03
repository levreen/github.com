package main

import (
	"fmt"
)

func negate(myBool *bool) {
	*myBool = !*myBool
}
func main() {
	myVal := true
	negate(&myVal)
	fmt.Println(myVal)
}
