package main

import (
	"fmt"
)

// printer displays information
type printer interface {
	print()
}

// user defines a user in the program.
type user struct {
	name string // method 
}

// print displays the user's name.
func (u user) print() {
	fmt.Println("User Name: %s\n", u.name) // calling method using dot notation
}
func main() {
	
	// create values of type user and admin.
	u := user{"Bill"}

	// add the values and pointers to the slice
	// of printer interface values.
	entities := []printer{
		// store a copy of the user value in the interface value.
		u,
		// store a copy of the address of the user value in the interfance
		&u,
	}

	// change the name field on the user value.
	u.name = "Bill_Change"

	// iterate over the slice of entities and call
	// print against the copied interface value. 

	for _, e := range entities {
		e.print()
	}

	// when we store a value, the interface value has its own
	// copy of the value. changes to the original value will 
	// not be seen. 

	// when we store a pointer, the interface value has its own
	// copy of the address. changes to the original value will
	// not be seen.
	
}

 
