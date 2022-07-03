package main

import (
	"fmt"
)

var subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	subscriber.name = "Aman Singh"
	subscriber.rate = 4.99
	subscriber.active = true
	fmt.Println("Name: ", subscriber.name)
	fmt.Println("Monthly rate: ", subscriber.rate)
	fmt.Println("Active? ", subscriber.active)
}
