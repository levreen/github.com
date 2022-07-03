package main

import (
	"fmt"
)

type OverheadError float64

func (o OverheadError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func main() {
	var o OverheadError
	o = 59.5
	fmt.Println(o.Error())
}
