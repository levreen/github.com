package main

import (
	"fmt"
)

type part struct {
	desc  string
	count int
}

func showInfo(p part) {
	fmt.Println("Description:", p.desc)
	fmt.Println("Count:", p.count)
}
func main() {
	var bolts part
	bolts.desc = "Hex Bolts"
	bolts.count = 24
	showInfo(bolts)
}
