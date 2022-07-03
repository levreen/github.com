package main

import (
	"fmt"
)

func find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Hello World")
}
