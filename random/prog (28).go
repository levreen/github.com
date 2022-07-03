package main

import (
	"fmt"
)

func main() {
	var ok bool
	ranks := make(map[string]int)
	var rank int
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank[] = %d, ok: = %v\n", rank, ok)
	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank[] = %d, ok: = %v\n", rank, ok)
}
