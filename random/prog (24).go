package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileSize, err := os.Stat("This is Levi!") // It only takes a file.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileSize.Size()) // Why Size() didn't take argument here? But only attach to the fileSize
}

// Same as time.Time.Now()
