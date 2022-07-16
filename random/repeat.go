// repeat.go
package main

import (
	"fmt"
	"time"
)

func statusUpdate() {}

func main() {
	go func() {
		for now := range time.Tick(time.Minute) {
			fmt.Println(now, statusUpdate)
		}
	}()

}
