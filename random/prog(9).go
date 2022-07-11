package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.FileMode(0700))
	fmt.Println(os.FileMode(0070))
	fmt.Println(os.FileMode(0007))
}
