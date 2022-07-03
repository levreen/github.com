package main

import (
	"fmt"
)

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Meow!")
}

// func AcceptAnything(thing interface{}) {
// 	fmt.Println(thing)
// 	whistle, ok := thing.(Whistle) // From interface to concrete type Whistle with MakeSound() method.
// 	if ok {
// 		whistle.MakeSound()
// 	}
// }

func AcceptWhistle(whistle Whistle) {
	fmt.Println(whistle)
	whistle.MakeSound()
}

// func main() {
// 	AcceptAnything(3.1415)
// 	AcceptAnything(Whistle("Tokyo canary"))
// }

func main() {
	AcceptWhistle("Levi")
}
