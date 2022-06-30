package calendar

import "errors"

// Date struct
type Date struct {
	year  int
	month int
	day   int
}

// Getter methods

// Getter method of Year
func (d *Date) Year() int {
	return d.year
}

// Getter method of Month
func (d *Date) Month() int {
	return d.month
}

// Getter method of Day
func (d *Date) Day() int {
	return d.day
}

// SetYear method
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

// SetMonth method
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

// SetDay method
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
