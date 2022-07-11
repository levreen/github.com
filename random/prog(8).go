package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%016b\n", os.O_RDONLY)
	fmt.Printf("%016b\n", os.O_WRONLY)
	fmt.Printf("%016b\n", os.O_RDWR)
	fmt.Printf("%016b\n", os.O_CREATE)
	fmt.Printf("%016b\n", os.O_APPEND)

	fmt.Println("\n\n\n")

	fmt.Printf("%016b\n", os.O_WRONLY|os.O_CREATE) // Or bitwise combines!
}
