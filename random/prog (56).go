package main

import (
	"fmt"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) { // Needs to be a pointer receiver, so original value can be updated.
	d.Year = year // Updates the copy, not the original WITHOUT pointer to the Date struct.
}

func main() {
	date := Date{}
	date.SetYear(2023) // Automatically convert to a pointer.
	fmt.Println(date.Year)
}
