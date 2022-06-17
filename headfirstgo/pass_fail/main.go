package main

import (
	"fmt"
	"log"
	"math/rand" // package's import path
	"time"

	"github.com/headfirstgo/keyboard"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second)

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1

	fmt.Println("I've chosen a random number between 1 to 100.")
	fmt.Println("Can you guess it?")
	//+!

	success := false
	for i := 10; i > 0; i-- {
		fmt.Printf("You only have %v left to guess\n", i)
		fmt.Print("Make a guess: ")
		guess, err := keyboard.GetFloat()
		if err != nil {
			log.Fatal(err)
		}

		if int(guess) < target { // Type conversion here from float64 to int. Then it works. Because I'm getting an error based on the float, int not compatible.
			fmt.Println("Oops. Your guess was LOW.")
		} else if int(guess) > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		}
	}

	if !success { // not false
		fmt.Println("Sorry, you didn't guess my number. It was: ", target)
	}

	//-!

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %.5s seconds\n", elapsed)
}
