package main

import (
	"fmt"
)

type subs struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s subs) {
	fmt.Println("Name: ", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active? ", s.active)
}
func defaultSubs(name string) subs {
	var s subs
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}
func main() {
	subs1 := defaultSubs("Levi Ocampo")
	subs1.rate = 4.99
	printInfo(subs1)

	subs2 := defaultSubs("Noreen Neric")
	subs2.active = false
	printInfo(subs2)
}
