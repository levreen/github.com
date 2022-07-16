package main

import (
	"fmt"
)

func calmDown() {
	recover() // If I comment out recover() the program will panic. 
				// Although, we deferred calmDown(), it won't run.
	fmt.Println("deferred function")
}

func freakOut() {
	defer calmDown()
	panic("oh no!")
}

func main() {
	freakOut()
	fmt.Println("Exiting Normally")
}
