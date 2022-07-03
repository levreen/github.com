package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

// duration is a named type with a base type of int.
type duration int

// notify implements the notifier interface.
func (d *duration) notify() {
	fmt.Println("Sending Notification in", *d)
}

func main() {
	
	duration(42).notify()

	// cannot call pointer method on duration(42)
	// cannot take the address of duration(42)
}
