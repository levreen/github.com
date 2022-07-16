// test111
// deadlock deadblocks
// a deadlock happens when a group of goroutines are
// waiting for each other and none of them is able to
// proceed. I know :(. Hello, Crush!
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
	// no call to goroutines
}
