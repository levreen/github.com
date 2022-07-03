package main

import (
	"fmt"
)

func count(start, end int) {
	fmt.Printf("count(%d, %d) called\n", start, end)
	if start < end { // <--- if we haven't reached the ending number
		count(start+1, end) // <--- the count function calls itself, with a starting number hihger than before.
	}
	fmt.Printf("Returning from count (%d, %d) call\n", start, end)
}
func main() {
	count(1, 3)
}
