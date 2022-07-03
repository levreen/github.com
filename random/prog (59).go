package main

import (
	"fmt"
)

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

type NoiseMaker interface {
	MakeSound()
}

// func main() {
// 	var toy NoiseMaker
// 	toy = Whistle("Toyco Canary")
// 	toy.MakeSound()
// 	toy = Horn("Toyco Blaster")
// 	toy.MakeSound()
// }

func play(n NoiseMaker) {
	n.MakeSound()
	// n.Walk() <---------- as long as you don't use the other methods not included on the interface.
}

func main() {
	// play(Whistle("Toyco Canary"))
	// play(Horn("Toyco Blaster"))
	play(Robot("Botco Ambler"))
}
