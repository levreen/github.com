// funcWithNOReturnValue.go
package main

import . "fmt"

func main() {
	greet("Levi")
}

func greet(name string) {
	Println("hello, ", name)
}
