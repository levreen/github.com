package main

import (
	"fmt"
)

type subscriber struct {
	rate float64
}

func applyDiscount(s *subscriber) { // We take the value of subscriber type by means of modifying its value. 
	s.rate = 4.99
}

func main() {
	var s subscriber
	applyDiscount(&s) // We access the address of s which is a subscriber type. It has zero value for rate but since
	fmt.Println(s.rate) // it get its value from the subscriber struct. It will return what it has for the rate. 
 }
