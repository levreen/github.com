package main

import (
	"fmt"
)

func calmDown() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error()) // Logic is so good. Love it! <3
	}
}
func main() {
	defer calmDown()
	err := fmt.Errorf("there's an error")
	panic(err) // Recover returns a value from Panic from Errorf. Appreciate! 
}
