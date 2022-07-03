package main

import (
	"fmt"
)

// player represents someone playing our game.
type player struct {
	name string
	score int
}

func main() {
	
	// declare a map with initial values using a map literal.
	players := map[string]player{
		"anne": {"Anna", 42},
		"jacob": {"Jacob", 21},
	}

	// trying to take the address of a map element fails.
	// anna := &players["anna"]
	// anna.score++

	// ./example4.go:23:10 cannot take the address of player["anna"]

	// instead take the element, modify it, and put it back.
	player := players["anna"]
	player.score++
	players["anna"] = player

}
