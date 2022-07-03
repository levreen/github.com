package main

import "fmt"

func status(grade float64) string { // It returns "failing" which is a string.
	if grade < 60.0 {
		return "failing"
	}
	return "passing" // You don't have to declare or make a statement if the grade is greater than 60.
}
func main() {
	fmt.Println(status(60.2))
}
