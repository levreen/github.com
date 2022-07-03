package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	err := os.Stat
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NO error")
}

// Same as time.Time.Now()
