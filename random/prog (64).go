package main

import (
	"fmt"
)

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}


func main() {
	var err error // so there's a type error
	err = ComedyError("What's a progmmer's favorite better? Logger!")
	fmt.Println(err)
}
