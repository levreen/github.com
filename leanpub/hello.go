// hello.go
package main

import . "fmt"

const Hello = "hello"

var world = "pinas"

func main() {
	//world += "!"
	//world := world + "!"
	var w string
	w = world + "!!!"
	Println(Hello, w)
}

// Shadowing accessing global package level in local package level
// having the same variable name.
