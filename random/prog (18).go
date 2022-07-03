package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	// fmt.Println(CToF(50))
	// fmt.Println(FToC(50))

	fmt.Printf("%g\n", BoilingC-FreezingC) // "100"°C
	boilingF := CToF(BoilingC)
	fmt.Printf("boilingF: %T & CToF(FreezingC): %T: %g\n", boilingF-CToF(FreezingC), boilingF, CToF(FreezingC)) // "180" °F
	//fmt.Printf("%g\n", boilingF-FreezingC) // compilr error: type mismatch
	fmt.Printf("%T\n", boilingF)
	fmt.Printf("%T\n", FreezingC)
	fmt.Println("\n\n\n")

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	//fmt.Println(c==f)// compile error; type mismatch
	fmt.Println(c == Celsius(f)) // "true"

	fmt.Println("\n\n\n")
	cel := FToC(212.0)
	fmt.Println(cel.String()) // 100°C
	fmt.Printf("%v\n", cel)   // 100°C; no need to call String explicitly
	fmt.Printf("%s\n", cel)
	fmt.Println(cel)
	fmt.Printf("%g\n", cel)
	fmt.Println(float64(cel))

	fah := CToF(100)
	fmt.Println(fah)

}

func (c Celsius) String() string { // formatting - galing nga eh 
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
