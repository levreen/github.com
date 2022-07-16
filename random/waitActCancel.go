// waitActCancel.go
package main

import (
	"fmt"
)

func main() {
	func Foo() {
		timer = time.AfterFunc(time.Minute, func() {
			log.Println("Foo run for more than a minute")
		})
		defer timer.Stop()
	}
}
