package main

import (
	"fmt"
)

func main() {
	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Print("-")
	for i, _ := range notes {
		fmt.Print(notes[i], "-")
	}
}
