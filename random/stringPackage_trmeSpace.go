package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "\t formerly surrounded by space \n"
	fmt.Printf("%#v", strings.TrimSpace(s))
}
