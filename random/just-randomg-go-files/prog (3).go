package main

import (
	"fmt"
)

func freakOut() {
	panic("oh no!")
	recover()
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}

// Panic and Recover do not sit together.