package main

import (
	"fmt"
)

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything(Whistle("Tokyo Canary"))
}
