package main

import (
	"fmt"
)

type T struct {
	f string
}

func main() {
	v1 := [...]T{{"foo"}, {"bar"}}
	fmt.Printf("%#v\n", v1)
	v2 := [2]T{{"foo"}, {"bar"}}
	fmt.Printf("%#v\n", v2)
	v3 := []T{{"foo"}, {"bar"}}
	fmt.Printf("%#v\n", v3)
}

/* OUTPUT:

[2]main.T{main.T{f:"foo"}, main.T{f:"bar"}}
[2]main.T{main.T{f:"foo"}, main.T{f:"bar"}}
[]main.T{main.T{f:"foo"}, main.T{f:"bar"}}
*/
