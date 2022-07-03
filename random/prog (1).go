package main

import (
	"fmt"
)

// data is a struct to bind methods to.
type data struct {
	name string
	age  int
}

// displayName provides a pretty print view of the name.
func (d data) displayName() { // takes value of d
	fmt.Println("My Name Is", d.name)
}

// setAge sets the age and displays the value.
func (d *data) setAge(age int) { // takes pointer of d
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

func main() {

	// declare a variable of type data.
	d := data{
		name: "Bill",
	}

	fmt.Println("Proper Calls to Methods:")

	// How we actually call methods in Go.
	d.displayName()

	fmt.Println("\nWhat the Compiler is Doing:")

	// this is what GO is doing underneath
	data.displayName(d)
	(*data).setAge(&d, 45)

	fmt.Print("\n\n\n")
	// the function variable will get its own copy of d because
	// is using a value receiver.
	f1 := d.displayName

	// call the method via the variable
	f1()

	fmt.Println()

	// change the value of d.
	d.name = "Joan"

	// call the method via the variable. We don't see the change
	f1()

	// declare a func var for the method
	// the function variable will get the address
	// is sing a pointer receiver.
	f2 := d.setAge
	f2(50)

	// call the method via the variable
	f2(45)

	// change the value of d.
	d.name = "Sammy"

	// call the method via the variable.
	f2(45)

	d.name = "Levi"

	fmt.Println()

	f2(50)
}
