package main

import (
	"fmt"
)

type part struct {
	desc  string
	count int
}

func minOrder(desc string) part {
	var p part
	p.desc = desc
	p.count = 100
	return p
}
func main() {
	p := minOrder("Hex Bolts")
	fmt.Printf("Type of P %T : C %T\n", p.desc, p.count)
	fmt.Printf("Type: %T", minOrder("Levi"))
}
