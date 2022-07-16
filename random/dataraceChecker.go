// dataraceChecker
// using go -race
package main

import (
	. "fmt"
)

func main() {
	i := 0
	go func() {
		i++ // write
	}()
	Println(i)
}
