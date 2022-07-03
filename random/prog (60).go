package main

import (
	"fmt"
)

type Appliance interface {
	TurnOn()
}

type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}

type CoffePot string

func (c CoffePot) TurnOn() {
	fmt.Println("Powering up")
}
func (c CoffePot) Brew() {
	fmt.Println("Heating up")
}

func main() {
	var device Appliance
	device = Fan("Windco Breeze")
	device.TurnOn()
	device = CoffePot("LuxBrew")
	device.TurnOn()
}
