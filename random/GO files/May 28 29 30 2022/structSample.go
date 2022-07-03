package main

import "fmt"

// example represents a type with different fields.

type example struct { // 8-byte value because of alignment
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// declare a variable of type example set ot its
	// zero value.
	var e1 example

	// display the value.
	fmt.Printf("%#v\n", e1)

	// declare a variable of type example and init using
	// a struct literal
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// dispaly the field values
	fmt.Println("Flag \t", e2.flag)
	fmt.Println("Counter\t", e2.counter)
	fmt.Println("Pi \t", e2.pi)

}
