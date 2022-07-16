package main

import . "fmt"

func main() {
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// blocks until there's data avaiable on ch1 or ch2

	ch1 = nil // disables this channel
	select {
	case <-ch1:
		Println("Received from ch1")
	case <-ch2:
		Println("Received from ch2")
	}
}
