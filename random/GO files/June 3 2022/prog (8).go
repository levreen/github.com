package main

import (
	"fmt"
)

func main() {

	// create a map to track scores for players in a game.
	scores := make(map[string]int)

	// read the element at the key "anna" It is absent so we get
	// the zero-value for this map's value type.
	score := scores["anna"] // valsem new alloc score

	fmt.Println("Score: ", score) // output: nil

	// if we need to check for the present of a key we use
	// a 2 variable assignemnt. the 2nd variable is a bool.
	score, ok := scores["anna"]

	fmt.Println("Score:", score, "Present:", ok)

	// we can leverage the zero-value behavior to write
	// convenient code like this:
	scores["anna"]++ // adds 1 to anna value [anna:1]

	// without this behavior we would have to code in a
	// defensive way like this.
	if n, ok := scores["anna"]; ok {
		scores["anna"] = n + 1
	} else {
		scores["anna"] = 1
	}

	score, ok = scores["anna"]
	fmt.Println("Score:", score, "Present:", ok)
}
