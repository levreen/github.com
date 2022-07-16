package main

import (
	. "fmt"
	"time"
)

// publish prints text to stdout after the given time has expired.
// it does NOT block but returns right away.

func Publish(text string, delay time.Duration) {
	go func() { // <--- this is the anoymous function
		time.Sleep(delay)
		Println("BREAKING NEWS:", text)
	}() // We must call the anonymous function
}
func main() {
	Publish("A goroutine starts a new thread.", 5*time.Second)
	Println("Let's hope the news will published before I leave.")

	// Wait for the news to be published.
	time.Sleep(10 * time.Second)

	Println("Ten seconds later: I'm leaving now.")
}
