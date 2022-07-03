package main

import (
	"fmt"
)

func count(start, end int) {
	fmt.Println(start)
	if start < end { // <--- if we haven't reached the ending number
		count(start+1, end) // <--- the count function calls itself, with a starting number hihger than before.
	}
}
func main() {
	count(1, 3)
}
