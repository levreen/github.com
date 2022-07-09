package main

import (
	"fmt"
	"prose"
)

func main() {
	phrases := []string{"my parents"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}

// Hello
