package main

import (
	. "fmt"
	"time"
)

func main() {
	go Println("Hello from another goroutine")
	Println("Hello from main goroutine")

	time.Sleep(time.Second)
}
