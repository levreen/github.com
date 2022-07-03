// FtoC prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("Freezing:	%g°F = %.2f°C\n", freezingF, fToC(freezingF))
	fmt.Printf("Boiling:	%g°F = %.2f°C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
