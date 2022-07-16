// goroutinesKill
// kill goroutines thru stop channel
package main

import (
	"fmt"
)

//// incomplete
//func main() {
//	quit := make(chan bool)
//	go func() {
//		for {
//			select {
//			case <-quit:
//				return
//			default:
//				// ...
//			}
//		}
//	}()
//	quit <- true
//}

// Generator returns a channel that produces the number 1, 2, 3...
// to stop the underlying goroutine, send a number on this channel.
func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func main() {
	number := Generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	number <- 0
	fmt.Println(<-number)
}
