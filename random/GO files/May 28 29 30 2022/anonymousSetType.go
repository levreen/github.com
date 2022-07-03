package main

import (
	"fmt"
)

type bill struct {
	flag    bool
	counter int16
	pi      float32
}

type nancy struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// declare a variable of an anonymous type set
	// ot its zero value.
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// display the value.
	fmt.Printf("%#v\n", e1)

	// declare a variable of an anonymous type and init
	// using a struct literal
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	var b bill
	var n nancy
	b = e2
	fmt.Println(b, n)

	// display the values.
	fmt.Printf("%#v\n", e2)
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}

// name-type
// literal-type
// functions are first-class values - non-name-type
