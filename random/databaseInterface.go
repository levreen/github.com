package main

import "fmt"

type Database interface {
	Fetch(key string) (int, error)
}

var DB Database

func isOver9000() bool {
	i, err := DB.Fetch("powerlevel")
	if err != nil {
		// If this were a real program, this should return an error
		// but this is an example, so it's ok
		return false
	}

	if i > 9000 {
		return true
	}

	return false
}

func main() {
	if isOver9000() {
		fmt.Println("It's over 9000!!!")
	}
}
