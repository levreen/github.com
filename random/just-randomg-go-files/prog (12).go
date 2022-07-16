package main

import (
	. "fmt"
)

func main() {
	// never blocks because of this --->
	select {
		case x := <-ch:
		Println("Received", x)
		default:
		Println("Nothing available") // <----
	}
}

/*
// An infinite random binary sequence
rand := make(chan int)
for {
	select {
		case rand <- 0: // no statement
		case rand <- 1:
	}
}
*/


/*
// A blocking operation with a timeout
select {
	case news := <AFK:
	Println(news)
	case <- time.After(time.Minute):
	Println("Time out: No nws in one minute")
}
*/


// A statement that blocks forever
// select{}

// A typical use would be at the end of the main function in some
// multithreaded programs. When main returns, the program exits
// and it does not wait for ther other goroutines to complete.
