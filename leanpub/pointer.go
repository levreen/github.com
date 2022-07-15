// pointer.go
package main

import . "fmt"

type Text string

func main() {
	var name Text = "Ellie"
	var pointer_to_name *Text
	Println(pointer_to_name)
	Println(name)
	Println(&name)
	pointer_to_name = &name
	Println(pointer_to_name)
	Printf("name = %v stored at %v\n", name, pointer_to_name)
	Printf("pointer_to_name references %v\n", *pointer_to_name)
}
