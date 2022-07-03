package main

import (
	"fmt"
)

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToMilliters() Milliliters {
	return Milliliters(l * 1000)
}

func (m Milliliters) ToLiters() Liters {
	return Liters(m / 100)
}
func main() {
	liters := Liters(3)
	fmt.Printf("%0.1f liters is %0.1f milliliters\n", liters, liters.ToMilliters())
	milliliters := Milliliters(500)
	fmt.Printf("%0.1f milliliters is %0.1f liters\n", milliliters, milliliters.ToLiters())
}
