package main

import (
	"fmt"
)

func calmDown() {
	fmt.Println(recover())
}

func main() {
	defer calmDown()
	panic("oh no")
}

// Recover returns a value from panic when we print out in a deferred function.