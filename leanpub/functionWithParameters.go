// functionWithParameters.go
package main

import . "fmt"

func main() {
	Println(message("pinas!"))
}

func message(name string) (message string) {
	message = Sprintf("hello %v", name) // We can assign the Sprintf to a var.
	return                              // Since there is a explicit declaration
	// of string return message, no need to reference it in the return statement.
}
