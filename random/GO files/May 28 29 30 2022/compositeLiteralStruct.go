package main

import (
	"fmt"
)

type T struct {
	name string
	next *T
}

func main() {
	_ = T{"foo", &T{"bar", nil}}
	_ = T{"foo", {"bar", nil}}
}
/*
OUTPUT:
./prog.go:10:15: missing type in composite literal
*/