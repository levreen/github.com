package main

import "math"

// Vectro2D struct
type Vectro2D struct {
	x float64
	y float64
}

// Add method that returns Vectro2D
func (v1 Vectro2D) Add(v2 Vectro2D) Vectro2D {
	return Vectro2D{x: v1.x + v2.x, y: v1.y + v2.y}
}

// Substract method that returns Vectro2D
func (v1 Vectro2D) Substract(v2 Vectro2D) Vectro2D {
	return Vectro2D{x: v1.x - v2.x, y: v1.y - v2.y}
}

// Multiply method that returns Vectro2D
func (v1 Vectro2D) Multiply(v2 Vectro2D) Vectro2D {
	return Vectro2D{x: v1.x * v2.x, y: v1.y * v2.y}
}

// AddV method that returns Vectro2D
func (v1 Vectro2D) AddV(d float64) Vectro2D {
	return Vectro2D{x: v1.x + d, y: v1.y + d}
}

// Multiplyd method that returns Vectro2D
func (v1 Vectro2D) Multiplyd(d float64) Vectro2D {
	return Vectro2D{x: v1.x * d, y: v1.y * d}
}

// Divisiond method that returns Vectro2D
func (v1 Vectro2D) Divisiond(d float64) Vectro2D {
	return Vectro2D{x: v1.x / d, y: v1.y / d}
}

// unexported limit method
func (v1 Vectro2D) limit(lower, upper float64) Vectro2D {
	return Vectro2D{
		x: math.Min(math.Max(v1.x, lower), upper),
		y: math.Min(math.Max(v1.y, lower), upper),
	}
}

// Distance method
func (v1 Vectro2D) Distance(v2 Vectro2D) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
