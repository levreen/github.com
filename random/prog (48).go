package main

import (
	"fmt"
)

type MyType stringo

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func main() {
	value := MyType("a MyType value")
	value.sayHi()
}
