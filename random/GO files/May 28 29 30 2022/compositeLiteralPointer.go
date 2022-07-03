// if type of element or key is a pointer(*T) then inside composite literal it's valid
// valid to omit &T.
package main

import (
	"fmt"
)

type T struct {
	f string
}

func main() {
	v1 := map[string]*T{"foo": {"bar"}}
	fmt.Printf("%#v\n", v1)
	v2 := [...]*T{{"foo"}, {"bar"}}
	fmt.Printf("%#v\n", v2)
	v3 := []*T{{"foo"}, {"bar"}}
	fmt.Printf("%#v\n", v3)
}
