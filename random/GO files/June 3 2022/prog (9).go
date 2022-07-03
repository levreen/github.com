package main

import (
	"fmt"
)

// users defines a set of users.
type users []user

func main() {
	
	// declare and make a map that uses a slice as the key.
	u := make(map[users]int)

	// ./example3.go:22:invalid map key type users

	// iterate over the map.
	for key, value := range u {
		fmt.Println(key,value)
	}
}
