package main

import (
	"fmt"
	"math/rand"
	"time"
)

func awardPrize() {
	doorNumber := rand.Intn(3) + 4 // Random number from 1 to 4: 4 is panic mood.
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		panic("invalid door number")
	}
}

func main() {
	rand.Seed(time.Now().Unix()) // Use time and date, always unique, so it project randomness.  
	awardPrize()
}
