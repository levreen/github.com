// test222
// waiting for goroutines
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		// do work.
		wg.Done()
	}()

	go func() {
		// do work.
		wg.Done()
	}()
	wg.Wait()
}
