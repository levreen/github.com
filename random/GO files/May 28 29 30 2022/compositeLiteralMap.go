package main

import (
	"fmt"
)

type T struct {
	f string
}

func main() {
	v1 := map[string]T{"foo:": {"bar"}}
	fmt.Printf("%#v\n", v1)
	v2 := map[T]string{{"foo"}: "bar"}
	fmt.Printf("%#v\n", v2)
	v3 := map[T]T{{"foo"}: {"bar"}}
	fmt.Printf("%#v\n", v3)
}

// composite literal map