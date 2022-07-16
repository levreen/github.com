package main

import (
	. "fmt"
	"time"
)

func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		Println(text)
		close(ch)
	}()
	return ch
}

func main() {
	wait := Publish("important news", 2*time.Minute)
	<-wait

}
