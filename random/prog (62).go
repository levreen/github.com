package main

import (
	"fmt"
)

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}

func (t Truck) Brake() {
	fmt.Println("Stopping")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
	// No LoadCargo method
	// To use it we use "type assertion"
}

func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()
	truck, ok := vehicle.(Truck) // <--- This is the type assertion declaration. From interface to concrete type.
	if ok {
		truck.LoadCargo("test cargo")
	}
}

func main() {
	TryVehicle(Truck("Fnord F180"))
}
