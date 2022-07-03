package main

import (
	"fmt"
)

type MyType string

func (m MyType) WithReturn() int {
	return len(m)
}

func main() {
	value := MyType("MyType value")
	fmt.Println(value.WithReturn())
}
