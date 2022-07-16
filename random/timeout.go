// timeout
package main

import (
	"fmt"
	"time"
)

func main() {
	AFP := make(chan string)
	select {
	case news := <-AFP:
		fmt.Println(news)
	case <-time.After(time.Hour):
		fmt.Println("No news in an hour")
	}

	for alive := true; alive; {
		timer := time.NewTimer(time.Hour)
		select {
		case news := <-AFP:
			timer.Stop()
			fmt.Println(news)
		case <-timer.C:
			alive = false
			fmt.Println("No news in an hour. Service aborting.")
		}
	}
}
