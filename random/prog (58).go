package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) error { // Needs to be a pointer receiver, so original value can be updated.
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year // Updates the copy, not the original WITHOUT pointer to the Date struct.
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid days")
	}
	d.Day = day
	return nil
}

func main() {
	date := Date{}
	err := date.SetMonth(56)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Month)
}
