package main

import (
	"fmt"
)

type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)-otherNumber)
}

func main() {
	value := Number(10)
	value.Add(4)
	value.Subtract(6)
}
