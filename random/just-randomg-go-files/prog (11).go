package main

import (
	. "fmt"
)

func main() {
	// never blocks because of this --->
	select {
		case x := <-ch:
		Println("Received", x)
		default:
		Println("Nothing available") // <----
	}
}
