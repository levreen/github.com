package main

import (
	"fmt"
)

type Gallons float64
type Liters float64
type Milliliter float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliter) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}
func main() {
	valueOfLiter := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", valueOfLiter, valueOfLiter.ToGallons())
	valueofMilliter := Milliliter(5)
	fmt.Printf("%0.3f Milliliters equals %0.3f gallons\n", valueofMilliter, valueofMilliter.ToGallons())
}
