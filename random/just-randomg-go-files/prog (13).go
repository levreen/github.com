// A data race happens when two goroutines access the same
// variable concurrently, and at least one of the accesses is a write.

package main

import (
	"fmt"
)

func race() {
	wait := make(chan struct{})
	n := 0
	go func() {
		n++ // read, increment, write
		close(wait)
	}()
	n++ // conflicting access
	<-wait
	fmt.Println(n) // Output: <unspecified>
}

func main() {
	race()
}
