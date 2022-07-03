package main

import (
	"fmt"
)

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) {
	fmt.Println("Name: ", s.name)
	fmt.Println("Monthly rate: ", s.rate)
	fmt.Println("Active?", s.active)
}
func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Levi Ocampo")
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	fmt.Println("\n\n\n")

	subscriber2 := defaultSubscriber("Noreen Neric")
	printInfo(subscriber2)
}
