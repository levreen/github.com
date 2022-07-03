package main

import (
	"fmt"
)

func main() {

	// fmt.Println can be called with values of any type.
	fmt.Println("Hello, World")
	fmt.Println(12345)
	fmt.Println(3.14159)
	fmt.Println(true)

	// how can we do the same?
	myPrintln("Hello, World")
	myPrintln(12345)
	myPrintln(3.14159)
	myPrintln(true)

	// an interface is satisfied by any piece of data when the data exhibits
	// the full method set of behavior defined by the interface.
	// the empty interface defines no method set of behavior and therefore
	// requires no method by the data being stored.
}

func myPrintln(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Printf("Is string	: type(%T)	:	value(%s)\n", v, v)
	case int:
		fmt.Printf("Is string	: type(%T)	:	value(%s)\n", v, v)
	case float64:
		fmt.Printf("Is string	: type(%T)	:	value(%s)\n", v, v)
	default:
		fmt.Printf("Is string	: type(%T)	:	value(%s)\n", v, v)
	}
}
