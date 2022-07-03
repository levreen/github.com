package main

import (
	"fmt"
)

func main() {
	text := [3]string{
		"This is a series of long strings",
		"which would be awkward to place",
		"together on a single line.", // <------ this comma at the end is required.
	}

	for _, val := range text {
		fmt.Println(val)
	}
}
