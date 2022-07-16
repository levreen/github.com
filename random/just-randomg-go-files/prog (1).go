package main

import (
	"fmt"
)

func snack() {
	defer fmt.Println("Closing refrigerator")
	fmt.Println("Opening refrigerator")
	panic("refrigerator is empty")
}
func main() {
	snack()
}

/*
OUTPUT

Opening refrigerator
Closing refrigerator
[T+0000ms]
panic: refrigerator is empty

goroutine 1 [running]:
main.snack()
	/tmp/sandbox3342129147/prog.go:10 +0xb8
main.main()
	/tmp/sandbox3342129147/prog.go:13 +0x17
[T+0001ms]
Program exited.
*/